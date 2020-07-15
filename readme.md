# WiP

## ToDos

## references

## Netatmo api ref

### Units

| Metric | Unit |
|---|---|
| length | metric system |
| Wind | kph |
| Pressure | mbar |
| Temperature | °C |
| CO2 | ppm |
| Humidity | % |
| Noise | dB |

### /getstationsdata

sample

```http
https://api.netatmo.com/api/getstationsdata?get_favorites=false
```

#### /getstationsdata optional

- device_id
- get_favorites bool

### /getmeasure

sample:

```http
https://api.netatmo.com/api/getmeasure?device_id=70%3Aee%3A50%3A2c%3A8f%3Abc&scale=30min&type=temperature&type=humidity&type=pressure&type=co2&type=noise&optimize=false&real_time=false
```

- device_id

- scale []
    30min

- type []
  - temperature
  - humidity
  - pressure
  - co2
  - noise

#### /getmeasure optional

- date_begin
- date_end
- limit

## commands

to register a sub Package run `go install` inside the subpackage folder

- Set the date format on influx db `influx -precision rfc3339` or on a existing cli session `precision rfc3339`

## InfluxDB

[Create User in Influx db](https://docs.influxdata.com/influxdb/v1.8/administration/authentication_and_authorization/#authorization)

### Dev Env config

```sql
CREATE USER Dev WITH PASSWORD 'SecurePassword' WITH ALL PRIVILEGES
CREATE DATABASE dataCollect

```

### MacOs

`influxd -config /usr/local/etc/influxdb.conf` to start a local instance
