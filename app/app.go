package app

import (
	"time"

	"code.cyb3r.social/skat/goDataCollector/config"
	"code.cyb3r.social/skat/goDataCollector/services/discord"
	"code.cyb3r.social/skat/goDataCollector/services/influxdb"
	"code.cyb3r.social/skat/goDataCollector/services/netatmo"
	"code.cyb3r.social/skat/goDataCollector/services/pushover"
	"code.cyb3r.social/skat/goDataCollector/services/wellKnownIP"

	log "github.com/sirupsen/logrus"
)

var runTimer *time.Timer
var duration time.Duration
var exit bool
var conf config.DataCollectorConfig
var co2Notified bool

func configure(local bool) {
	log.Trace("Read Configuration.\n")

	if local {
		conf = config.ReadConfig("./appData/local.config.json")
	} else {
		conf = config.ReadConfig("./appData/config.json")
	}

	discord.SetConfig(conf.Discord)
	pushover.SetConfig(conf.PushOver)
	netatmo.SetConfig(conf.NetAtmo)
	influxdb.SetConfig(conf.InfluxDb)
}

func runBackgroundTasks(daemon bool, devLocal bool) {

	log.Tracef("task started with parameter %b\n", daemon)

	ip, notify := wellKnownIP.GetMyPublicIP()

	if notify {
		discord.PushIPChange(ip)
		pushover.PushIPChange(ip)
	}

	stationData, err := netatmo.GetStationsData()
	if err != nil {
		log.WithError(err).Error("Error from netatmo api")
	} else {

		log.Debugf("%v\n", stationData)
	}

	if stationData.Co2 > 750 && !co2Notified {
		pushover.PushNotification(stationData.Co2)
		co2Notified = true
	}

	if stationData.Co2 < 700 {
		co2Notified = false
	}

	influxdb.Write(stationData.ToPoint())

	if daemon {
		runTimer = time.AfterFunc(duration, func() {
			runBackgroundTasks(true, devLocal)
		})

	}

}

func Daemonize(devLocal bool) {
	log.Trace("Run on Daemon Mode\n")
	configure(devLocal)

	duration = time.Duration(20) * time.Second
	runBackgroundTasks(true, devLocal)

	for !exit {
		time.Sleep(30 * time.Second)
	}
}

func Run(devLocal bool) {
	log.Trace("Run on Single Mode\n")
	configure(devLocal)
	runBackgroundTasks(false, devLocal)
}
