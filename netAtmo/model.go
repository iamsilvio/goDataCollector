package netatmo

// Config Netatwo API configuration
type Config struct {
	UserName     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"clientsecret"`
}

// Place Netatmo API Place
type Place struct {
	Altitude int       `json:"altitude"`
	City     string    `json:"city"`
	Country  string    `json:"country"`
	Timezone string    `json:"timezone"`
	Location []float32 `json:"location"`
}

// Dashboard Netatmo API Dashboard
type Dashboard struct {
	Time             int64   `json:"time_utc"`
	Temperature      float32 `json:"Temperature"`
	Co2              int     `json:"CO2"`
	Humidity         int     `json:"Humidity"`
	Noise            int     `json:"Noise"`
	Pressure         float32 `json:"Pressure"`
	AbsolutePressure float32 `json:"AbsolutePressure"`
	MinTemp          float32 `json:"min_temp"`
	MaxTemp          float32 `json:"max_temp"`
	DateMinTemp      int64   `json:"date_min_temp"`
	DateMaxTemp      int64   `json:"date_max_temp"`
	TempTrend        string  `json:"temp_trend"`
	PressureTrend    string  `json:"pressure_trend"`
}

// Module Netatmo API Module
type Module struct {
	ID             string   `json:"_id"`
	Type           string   `json:"type"`
	ModuleName     string   `json:"module_name"`
	LastSetup      int64    `json:"last_setup"`
	DataType       []string `json:"data_type"`
	BatteryPercent int      `json:"battery_percent"`
	Reachable      bool     `json:"reachable"`
	Firmware       int      `json:"firmware"`
	LastMessage    int64    `json:"last_message"`
	LastSeen       int64    `json:"last_seen"`
	RfStatus       int      `json:"rf_status"`
	BatteryVp      int      `json:"battery_vp"`
}

// Device Netatmo API Device
type Device struct {
	ID              string    `json:"_id"`
	StationName     string    `json:"station_name"`
	DateSetup       int64     `json:"date_setup"`
	LastSetup       int64     `json:"last_setup"`
	Type            string    `json:"type"`
	LastStatusStore int64     `json:"last_status_store"`
	ModuleName      string    `json:"module_name"`
	Firmware        int       `json:"firmware"`
	LastUpgrade     int64     `json:"last_upgrade"`
	WifiStatus      int       `json:"wifi_status"`
	Reachable       bool      `json:"reachable"`
	Co2Calibrating  bool      `json:"co2_calibrating"`
	DataType        []string  `json:"data_type"`
	Place           Place     `json:"place"`
	Dashboard       Dashboard `json:"dashboard_data"`
	Modules         []Module  `json:"modules"`
}

// AdminSetting Netatmo API AdminSetting
type AdminSetting struct {
	Lang         string `json:"lang"`
	RegLocale    string `json:"reg_locale"`
	Unit         int    `json:"unit"`
	Windunit     int    `json:"windunit"`
	Pressureunit int    `json:"pressureunit"`
	FeelLikeAlgo int    `json:"feel_like_algo"`
}

// User Netatmo API User
type User struct {
	Mail           string       `json:"mail"`
	Administrative AdminSetting `json:"administrative"`
}

// APIResponseBody Netatmo API ResponseBody
type APIResponseBody struct {
	Devices []Device `json:"devices"`
	User    User     `json:"user"`
}

// APIResponse Netatmo API Response
type APIResponse struct {
	Body       APIResponseBody `json:"body"`
	Status     string          `json:"status"`
	TimeExec   float64         `json:"time_exec"`
	TimeServer int64           `json:"time_server"`
}
