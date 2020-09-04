package data

import (
	"fmt"
	"strings"
	"time"
)

// Config of the influxDb connection
type Config struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	ServerURL string `json:"serverurl"`
	Bucket    string `json:"bucket"`
}

// Point influxDB DataPoint
type Point struct {
	Title  string //measurement
	Tags   map[string]string
	Fields map[string]string
	Date   time.Time
}

// NewDataPoint init a new DataPoint
func NewDataPoint() *Point {
	var d Point
	d.Tags = make(map[string]string)
	d.Fields = make(map[string]string)
	return &d
}

// ToLineProtocol converts a dataPoint to lineProtocol String
func (p *Point) ToLineProtocol() string {

	result := fmt.Sprintf("%s", p.Title)

	for k, v := range p.Tags {
		result = fmt.Sprintf("%s,%s=%s", result, k, v)
	}

	result = fmt.Sprintf("%s ", result)

	for k, v := range p.Fields {
		result = fmt.Sprintf("%s%s=%s,", result, k, v)
	}

	result = strings.TrimSuffix(result, ",")

	result = fmt.Sprintf("%s %d", result, p.Date.UnixNano())

	return result
}
