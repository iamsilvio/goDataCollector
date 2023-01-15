package netatmo

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	write "github.com/influxdata/influxdb-client-go/v2/api/write"
)

func (net *Dashboard) ToPoint() *write.Point {

	p := influxdb2.NewPointWithMeasurement("HomeClima").
		AddTag("Room", "rabbitWarren").
		AddTag("Service", "NetAtmoAPI").
		AddField("Temperature", fmt.Sprintf("%f", net.Temperature)).
		AddField("Co2", fmt.Sprintf("%d", net.Co2)).
		AddField("Humidity", fmt.Sprintf("%d", net.Humidity)).
		AddField("Noise", fmt.Sprintf("%d", net.Noise)).
		AddField("Pressure", fmt.Sprintf("%f", net.Pressure)).
		AddField("AbsolutePressure", fmt.Sprintf("%f", net.AbsolutePressure)).
		SetTime(time.Unix(net.Time, 0))

	return p
}
