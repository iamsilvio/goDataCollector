package data

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"

	influxdb "github.com/influxdata/influxdb-client-go"
)

var config Config
var roots *x509.CertPool
var client influxdb.Client

// SetConfig sets the influxDB configuration
func SetConfig(conf Config) {
	config = conf
	roots = x509.NewCertPool()

	client = influxdb.NewClientWithOptions(config.ServerURL, config.UserName+":"+config.Password,
		influxdb.DefaultOptions().SetUseGZip(true).SetTlsConfig(&tls.Config{
			RootCAs:            roots,
			InsecureSkipVerify: true}))
}

// Write write dataPoint to influx db
func Write(data *Point) {

	writeAPI := client.WriteApiBlocking("", config.Bucket)
	err := writeAPI.WriteRecord(context.Background(), data.ToLineProtocol())
	if err != nil {
		log.Printf("Failed to write to Influxdb : %s\n", err)
	}

	client.Close()

}
