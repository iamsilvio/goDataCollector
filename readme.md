# WiP

## Dev env

- start the influxDB 2 container `docker compose -f docker-compose.dev.yml up`
- make shure you have a config file `./appData/config.json`
- run the app with `go run . -d`
- look at `main.go` for additional options or `go run . -h`

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
