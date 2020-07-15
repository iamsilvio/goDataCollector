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
	"code.cyb3r.social/skat/goDataCollector/netAtmo"
)

func runNetatmoStuff() {
	fmt.Printf("%s:/n", time.Now())

	device, err := netAtmo.GetStationsData()
	if err != nil {
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
	runTimer = time.AfterFunc(duration, runNetatmoStuff)
}

var runTimer *time.Timer
var duration time.Duration
var exit bool

type dataCollectorConfig struct {
	InfluxDb data.InfluxDbConfig `json:"influxdb"`
	NetAtmo  netAtmo.ApiConfig   `json:"netatmo"`
}

func readConfig(path string) dataCollectorConfig {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read config file %s", path)
		log.Fatal(err)
	}

	var conf dataCollectorConfig

	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Printf("Failed to parse config file %s", path)
		log.Fatal(err)
	}
	return conf
}

func main() {

	daemonPtr := flag.Bool("daemon", false, "run the modules as Daemon")
	flag.Parse()

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := path.Dir(ex)

	if *daemonPtr == true {

		conf := readConfig(dir + "/config.json")

		data.SetConfig(conf.InfluxDb)
		netAtmo.SetConfig(conf.NetAtmo)

		duration = time.Duration(5) * time.Minute
		runTimer = time.AfterFunc(duration, runNetatmoStuff)

		for exit != true {
			time.Sleep(30 * time.Second)
		}

	} else {
		conf := readConfig("config.dev.json")

		data.SetConfig(conf.InfluxDb)
		netAtmo.SetConfig(conf.NetAtmo)
	}

}
