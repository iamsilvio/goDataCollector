package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"code.cyb3r.social/skat/goDataCollector/data"
	"code.cyb3r.social/skat/goDataCollector/discord"
	"code.cyb3r.social/skat/goDataCollector/netatmo"
	"code.cyb3r.social/skat/goDataCollector/pushover"
	"code.cyb3r.social/skat/goDataCollector/wellknownip"
)

func runBackgroundTasks() {

	runNetatmoStuff()

	if ipUpdated() {
		pushover.PushIPChange(lastIP)
		discord.PushIPChange(lastIP)
		saveIP()
	}

	runTimer = time.AfterFunc(duration, runBackgroundTasks)
}

func runNetatmoStuff() {

	device, err := netatmo.GetStationsData()
	if err == nil {
		d := data.NewDataPoint()
		d.Title = "HomeClima"
		d.Date = time.Unix(device.Time, 0)
		d.Tags["Room"] = "rabbitWarren"
		d.Tags["Service"] = "NetAtmoAPI"
		d.Fields["Temperature"] = fmt.Sprintf("%f", device.Temperature)
		d.Fields["Co2"] = fmt.Sprintf("%d", device.Co2)
		d.Fields["Humidity"] = fmt.Sprintf("%d", device.Humidity)
		d.Fields["Noise"] = fmt.Sprintf("%d", device.Noise)
		d.Fields["Pressure"] = fmt.Sprintf("%f", device.Pressure)
		d.Fields["AbsolutePressure"] = fmt.Sprintf("%f", device.AbsolutePressure)

		data.Write(d)
	}

	if device.Co2 > 750 {
		pushover.PushNotification(device.Co2)
	}
}

var lastIP string

func ipUpdated() bool {
	ip := wellknownip.GetMyPublicIP()

	if len(ip) > 0 && lastIP != ip {

		lastIP = ip
		return true
	}
	return false
}

func saveIP() {
	file, err := json.MarshalIndent(lastIP, "", " ")
	if err != nil {
		log.Printf("Failed to marshal ip: %v\n", err)
	}
	err = ioutil.WriteFile("ip.json", file, 0644)
	if err != nil {
		log.Printf("Failed to write ip file: %v\n", err)
	}
}

func loadIP() {

	file, err := ioutil.ReadFile("ip.json")
	if err != nil {
		log.Printf("Failed to read ip file: %v\n", err)
	}

	err = json.Unmarshal([]byte(file), &lastIP)
	if err != nil {
		log.Printf("Failed to unmarshal ip: %v\n", err)
	}
}

var runTimer *time.Timer
var duration time.Duration
var exit bool

type dataCollectorConfig struct {
	InfluxDb data.Config     `json:"influxdb"`
	NetAtmo  netatmo.Config  `json:"netatmo"`
	Discord  discord.Config  `json:"discord"`
	PushOver pushover.Config `json:"pushOver"`
}

func readConfig(path string) dataCollectorConfig {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read config file %s:%v\n", path, err)
	}

	var conf dataCollectorConfig

	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Printf("Failed to parse config file %s:%v\n", path, err)
	}
	return conf
}

func dev() {
	if ipUpdated() {
		pushover.PushIPChange(lastIP)
		discord.PushIPChange(lastIP)
		saveIP()
	}
}

func main() {

	daemonPtr := flag.Bool("daemon", false, "run the modules as Daemon")
	flag.Parse()

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := path.Dir(ex)

	loadIP()

	if *daemonPtr == true {

		conf := readConfig(dir + "/config.json")

		data.SetConfig(conf.InfluxDb)
		netatmo.SetConfig(conf.NetAtmo)
		pushover.SetConfig(conf.PushOver)
		discord.SetConfig(conf.Discord)

		duration = time.Duration(20) * time.Second
		runBackgroundTasks()

		for exit != true {
			time.Sleep(30 * time.Second)
		}

	} else {
		conf := readConfig("config.local.json")

		data.SetConfig(conf.InfluxDb)
		netatmo.SetConfig(conf.NetAtmo)
		pushover.SetConfig(conf.PushOver)
		discord.SetConfig(conf.Discord)

		dev()

	}
}
