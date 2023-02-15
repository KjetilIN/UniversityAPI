package handlers

//Default handlets return struct
type ProjectInfo struct{
	Author string `json:"author"`
	Version string `json:"version"`
}

// Diag handler return struct 
type StatusInfo struct{
	UniApi string `json:"universitiesapi"`
	CountryApi string `json:"countriesapi"`
	Version string `json:"version"`
	Uptime int `json:"uptime"`
}	

// University information structs used for the endpoints 
type UniverityInfo struct {
	Name string `json:"name"`
	Country string `json:"country"`
	Isocode string `json:"isocode"`
	Webpages []string `json:"webpages"`
	Languages []string `json:"languages"`
	MapLink string `json:"map"`
}