package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func readConfig(path string) InfluxDbConfig {

	// read file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	var config InfluxDbConfig
	// unmarshall it
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error:", err)
	}

	return config
}

var config InfluxDbConfig

func SetConfig(conf InfluxDbConfig) {
	config = conf

}

func Write(data *DataPoint) {

	client := influxdb2.NewClient(config.ServerURL, config.UserName+":"+config.Password)
	writeApi := client.WriteApiBlocking("", config.Bucket)

	//	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	writeApi.WriteRecord(context.Background(), data.ToLineProtocol())

	fmt.Println(data.ToLineProtocol())

	client.Close()
}
