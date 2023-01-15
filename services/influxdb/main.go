package influxdb

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	write "github.com/influxdata/influxdb-client-go/v2/api/write"
	log "github.com/sirupsen/logrus"
)

var config Config

var client influxdb2.Client

// SetConfig sets the influxDB configuration
func SetConfig(conf Config) {
	config = conf

	client = influxdb2.NewClient(conf.ServerURL, conf.Token)

}

// Write write dataPoint to influx db
func Write(data *write.Point) {
	defer client.Close()

	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(config.Organisation, config.Bucket)

	err := writeAPI.WritePoint(context.Background(), data)
	if err != nil {
		log.WithError(err).Error("Failed to write to Influxdb\n")
		panic(err)
	}

}
