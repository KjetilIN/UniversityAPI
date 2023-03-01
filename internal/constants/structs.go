package constants

// Diag handler return struct 
type StatusInfo struct{
	UniApi string `json:"universitiesapi"`
	CountryApi string `json:"countriesapi"`
	Version string `json:"version"`
	Uptime int `json:"uptime"`
}	

//Struct with the country info
type CountryInfo struct{
	Languages map[string]string `json:"languages"`
	Region string `json:"region"`
}

type BorderCountries struct{
	Borders []string  `json:"borders"`
}

type CountryName struct{
	Name struct{
		Common string `json:"common"`
	} `json:"name"`
}



// Struct used by encoder to get the information from the university api
type UniStruct struct{
	Name string `json:"name"`
	Country string `json:"country"`
	Isocode string `json:"alpha_two_code"`
	Webpages []string `json:"web_pages"`
}

// University information structs used for the endpoints 
type UniversityInfo struct {
	UniStruct
	CountryInfo
}
