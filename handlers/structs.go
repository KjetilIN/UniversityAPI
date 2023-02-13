package handlers

type ProjectInfo struct{
	Author string `json:"author"`
	Version string `json:"version"`
}

type StatusInfo struct{
	UniApi string `json:"universitiesapi"`
	CountryApi string `json:"countriesapi"`
	Version string `json:"version"`
	Uptime int `json:"uptime"`
}	
