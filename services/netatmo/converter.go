package netatmo

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	write "github.com/influxdata/influxdb-client-go/v2/api/write"
)

func (net *Dashboard) ToPoint() *write.Point {

	p := influxdb2.NewPointWithMeasurement("HomeClima").
		AddTag("Room", "rabbitWarren").
		AddTag("Service", "NetAtmoAPI").
		AddField("Temperature", net.Temperature).
		AddField("Co2", net.Co2).
		AddField("Humidity", net.Humidity).
		AddField("Noise", net.Noise).
		AddField("Pressure", net.Pressure).
		AddField("AbsolutePressure", net.AbsolutePressure).
		SetTime(time.Unix(net.Time, 0))

	return p
}
