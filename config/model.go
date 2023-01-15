package config

import (
	"code.cyb3r.social/skat/goDataCollector/services/discord"
	"code.cyb3r.social/skat/goDataCollector/services/influxdb"
	"code.cyb3r.social/skat/goDataCollector/services/netatmo"
	"code.cyb3r.social/skat/goDataCollector/services/pushover"
)

type DataCollectorConfig struct {
	InfluxDb influxdb.Config `json:"influxdb"`
	NetAtmo  netatmo.Config  `json:"netatmo"`
	Discord  discord.Config  `json:"discord"`
	PushOver pushover.Config `json:"pushOver"`
}
