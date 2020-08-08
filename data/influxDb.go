package data

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

var config InfluxDbConfig
var roots *x509.CertPool

func SetConfig(conf InfluxDbConfig) {
	config = conf
	roots = x509.NewCertPool()

}

func Write(data *DataPoint) {

	client := influxdb2.NewClientWithOptions(config.ServerURL, config.UserName+":"+config.Password,
		influxdb2.DefaultOptions().SetUseGZip(true).SetTlsConfig(&tls.Config{
			RootCAs:            roots,
			InsecureSkipVerify: true}))

	writeApi := client.WriteApiBlocking("", config.Bucket)
	writeApi.WriteRecord(context.Background(), data.ToLineProtocol())

	client.Close()

}
