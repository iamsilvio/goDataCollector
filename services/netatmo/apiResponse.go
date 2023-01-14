package netatmo

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
