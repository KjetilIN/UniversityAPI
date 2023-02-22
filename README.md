# University API

The UniversityAPI service provides a RESTful API for retrieving information about universities in various countries. It allows users to query for universities by name or country, and returns data on the university's country, website, languages spoken and more. It uses third-party APIs to gather data from universities and countries. The service is built using the Go programming language and follows best practices for scalable and maintainable code, including a modular architecture and well-documented API endpoints. Whether you're a student researching potential universities, or a developer looking for a lightweight API to integrate into your application, the UniversityAPI service provides an easy-to-use and efficient solution.

> Assignment 1 <br>
> Version: V1 <br>
> PROG2005 Cloud Technologies (2023 VÅR)<br>


## End Points & How to use them 

The service provide a set of endpoints:

```
/unisearcher/{VERSION}/uniinfo/
/unisearcher/{VERSION}/neighbourunis/
/unisearcher/{VERSION}/diag/
```

Each endpoint has to be used in a certain way. There will be a response that can give you a hint of what is done wrong.

## University Information 
> Method: GET <br>
> Path: unisearcher/{VERSION}/uniinfo/{:partial_or_complete_university_name}/


This endpoint is uses the partial or complete name for the university. 
The response could be of different length but will always have JSON objects in the same structure:


```
[
	{
		"name": <Name of university>,
		"country": <The country the university is in>,
		"alpha_two_code": <Alpha two code of the country>,
		"web_pages": [<List of website links to the university>],
		"languages": {<Map of languages spoken by the university>},
		"region": <Region of the university>
	},
	....
]

```

Example response: 

```
[
	{
		"name": "Häme University of Applied Sciences",
		"country": "Finland",
		"alpha_two_code": "FI",
		"web_pages": [
			"https://www.hamk.fi/"
		],
		"languages": {
			"fin": "Finnish",
			"swe": "Swedish"
		},
		"region": "Europe"
	},
	....

```

## University Neighbors 

>Method: GET <br>
>Path: unisearcher/{VERSION}/neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}

Given a country and a partial name of the university, this endpoints find information for all countries that are a neighbor country.
This does not include the universities from the country given. This endpoint will be the same structure as the `uniinfo` endpoint. 


## Diagnostic Endpoint

> Method: GET
> Path: unisearcher/{VERSION}/diag/

To check if all third party endpoints are up, and some additional server information, use this endpoint.
Additionally some information on the server. 

Response: 

```
{
   "universitiesapi": "<http status code for universities API>",
   "countriesapi": "<http status code for restcountries API>",
   "version": <Version of the service>,
   "uptime": <time in seconds from the last service restart>
}

```


# Third Party APIs. 

 This service utilizes third-party APIs to retrieve and process data for its endpoints. See the `diag` endpoint for checking if all endpoints works as expected. The API integrates with external services to provide information on university data and country data. The use of third-party APIs enables the service to offer a wider range of features and functionalities while reducing the need for extensive data storage and processing on its servers.

### UNI API


The University API is a RESTful service that allows you to get information about universities in Norway. With this API, you can get the names of universities, the number of students enrolled, the number of faculties, and other details. The service uses third-party APIs to collect data about universities and then formats it into a JSON response.

To use the University API, you can make GET requests to the various endpoints provided by the service. The API provides endpoints to get the list of universities, detailed information about a specific university, and information about neighboring universities. You can use this information to create web applications or dashboards that display information about universities in Norway.

The code for the University API is available on GitHub at https://github.com/username/UniversityAPI. You can clone the repository and run the service on your local machine or deploy it to a cloud platform like AWS or Heroku. The API is written in Golang and uses popular libraries like Gorilla Mux for routing and handling HTTP requests.


- Link: `http://universities.hipolabs.com/search?name={NAME}&country={COUNTRY}`


Response as JSON (University in Turkey): 

```
[
	...
	{
	    "alpha_two_code": "TR",
	    "country": "Turkey",
	    "state-province": null,
	    "domains": [
	        "sabanciuniv.edu",
	        "sabanciuniv.edu.tr"
	    ],
	    "name": "Sabanci University",
	    "web_pages": [
	        "http://www.sabanciuniv.edu/",
	        "http://www.sabanciuniv.edu.tr/"
	    ],
	},
	...
]

```


### Country API

- Link use: 
    - `https://restcountries.com/v3.1/name/{COUNTRY_NAME}`
    - `https://restcountries.com/v3.1/aplha/{COUNTRY_ALPHA}`


JSON Response of Country (In this example Colombia): 

```
[[{
	"name": "Colombia",
	"topLevelDomain": [".co"],
	"alpha2Code": "CO",
	"alpha3Code": "COL",
	"callingCodes": ["57"],
	"capital": "Bogotá",
	"altSpellings": ["CO", "Republic of Colombia", "República de Colombia"],
	"region": "Americas",
	"subregion": "South America",
	"population": 48759958,
	"latlng": [4.0, -72.0],
	"demonym": "Colombian",
	"area": 1141748.0,
	"gini": 55.9,
	"timezones": ["UTC-05:00"],
	"borders": ["BRA", "ECU", "PAN", "PER", "VEN"],
	"nativeName": "Colombia",
	"numericCode": "170",
	"currencies": [{
		"code": "COP",
		"name": "Colombian peso",
		"symbol": "$"
	}],
	"languages": [{
		"iso639_1": "es",
		"iso639_2": "spa",
		"name": "Spanish",
		"nativeName": "Español"
	}],
	"translations": {
		"de": "Kolumbien",
		"es": "Colombia",
		"fr": "Colombie",
		"ja": "コロンビア",
		"it": "Colombia",
		"br": "Colômbia",
		"pt": "Colômbia"
	},
	"flag": "https://restcountries.com/data/col.svg",
	"regionalBlocs": [{
		"acronym": "PA",
		"name": "Pacific Alliance",
		"otherAcronyms": [],
		"otherNames": ["Alianza del Pacífico"]
	}, {
		"acronym": "USAN",
		"name": "Union of South American Nations",
		"otherAcronyms": ["UNASUR", "UNASUL", "UZAN"],
		"otherNames": ["Unión de Naciones Suramericanas", "União de Nações Sul-Americanas", "Unie van Zuid-Amerikaanse Naties", "South American Union"]
	}]
}]

