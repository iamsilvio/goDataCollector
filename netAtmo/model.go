package netAtmo

type ApiConfig struct {
	UserName     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"clientsecret"`
}

type Place struct {
	Altitude int       `json:"altitude"`
	City     string    `json:"city"`
	Country  string    `json:"country"`
	Timezone string    `json:"timezone"`
	Location []float32 `json:"location"`
}

type Dashboard struct {
	Time             int64   `json:"time_utc"`
	Temperature      float32 `json:"Temperature"`
	Co2              int     `json:"CO2"`
	Humidity         int     `json:"Humidity"`
	Noise            int     `json:"Noise"`
	Pressure         float32 `json:"Pressure"`
	AbsolutePressure float32 `json:"AbsolutePressure"`
	Min_temp         float32 `json:"min_temp"`
	Max_temp         float32 `json:"max_temp"`
	Date_min_temp    int64   `json:"date_min_temp"`
	Date_max_temp    int64   `json:"date_max_temp"`
	Temp_trend       string  `json:"temp_trend"`
	Pressure_trend   string  `json:"pressure_trend"`
}

type Module struct {
	ID              string   `json:"_id"`
	Type            string   `json:"type"`
	Module_name     string   `json:"module_name"`
	Last_setup      int64    `json:"last_setup"`
	Data_type       []string `json:"data_type"`
	Battery_percent int      `json:"battery_percent"`
	Reachable       bool     `json:"reachable"`
	Firmware        int      `json:"firmware"`
	Last_message    int64    `json:"last_message"`
	Last_seen       int64    `json:"last_seen"`
	Rf_status       int      `json:"rf_status"`
	Battery_vp      int      `json:"battery_vp"`
}

type Device struct {
	ID                string    `json:"_id"`
	Station_name      string    `json:"station_name"`
	Date_setup        int64     `json:"date_setup"`
	Last_setup        int64     `json:"last_setup"`
	Type              string    `json:"type"`
	Last_status_store int64     `json:"last_status_store"`
	Module_name       string    `json:"module_name"`
	Firmware          int       `json:"firmware"`
	Last_upgrade      int64     `json:"last_upgrade"`
	Wifi_status       int       `json:"wifi_status"`
	Reachable         bool      `json:"reachable"`
	Co2_calibrating   bool      `json:"co2_calibrating"`
	Data_type         []string  `json:"data_type"`
	Place             Place     `json:"place"`
	Dashboard         Dashboard `json:"dashboard_data"`
	Modules           []Module  `json:"modules"`
}

type AdminSetting struct {
	Lang           string `json:"lang"`
	Reg_locale     string `json:"reg_locale"`
	Unit           int    `json:"unit"`
	Windunit       int    `json:"windunit"`
	Pressureunit   int    `json:"pressureunit"`
	Feel_like_algo int    `json:"feel_like_algo"`
}

type User struct {
	Mail           string       `json:"mail"`
	Administrative AdminSetting `json:"administrative"`
}

type ApiResponseBody struct {
	Devices []Device `json:"devices"`
	User    User     `json:"user"`
}

type ApiResponse struct {
	Body        ApiResponseBody `json:"body"`
	Status      string          `json:"status"`
	Time_Exec   float64         `json:"time_exec"`
	Time_Server int64           `json:"time_server"`
}
