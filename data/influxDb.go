package data

import (

	"log"
	"context"
	"crypto/tls"
	"crypto/x509"

 influxdb "github.com/influxdata/influxdb-client-go"
)

var config InfluxDbConfig
var roots *x509.CertPool
var client influxdb.Client

func SetConfig(conf InfluxDbConfig) {
	config = conf
	roots = x509.NewCertPool()

	client = influxdb.NewClientWithOptions(config.ServerURL, config.UserName+":"+config.Password,
		influxdb.DefaultOptions().SetUseGZip(true).SetTlsConfig(&tls.Config{
			RootCAs:            roots,
			InsecureSkipVerify: true}))
}

func Write(data *DataPoint) {

	writeApi := client.WriteApiBlocking("", config.Bucket)
	err := writeApi.WriteRecord(context.Background(), data.ToLineProtocol())
	if err != nil {
			log.Printf("Failed to write to Influxdb : %s\n", err)
	}

	client.Close()

}
