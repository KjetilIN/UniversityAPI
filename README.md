# University API

The UniversityAPI service provides a RESTful API for retrieving information about universities in various countries. It allows users to query for universities by name or country, and returns data on the university's country, website, languages spoken and more. It uses third-party APIs to gather data from universities and countries. The service is built using the Go programming language and follows best practices for scalable and maintainable code, including a modular architecture and well-documented API endpoints. Whether you're a student researching potential universities, or a developer looking for a lightweight API to integrate into your application, the UniversityAPI service provides an easy-to-use and efficient solution.

<br>

To run the project simply run the `main.go` file, with the following command:

```terminal 
	go run ./cmd/univerity-api/main.go 
```


> Assignment 1 <br>
> Version: v1 <br>
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


The University API is a RESTful service that allows you to get information about universities in the world. With this API, you can get the names of universities, website links, and other details. The service uses third-party APIs to collect data about universities and then formats it into a JSON response.

To use the University API, you can make GET requests to the various endpoints provided by the service. The API provides endpoints to get the list of universities and detailed information about a specific university. You can use this information to create web applications or dashboards that display information about universities.

The code for the University API is available on GitHub at https://github.com/Hipo/university-domains-list <br>

✅ Usage: <br>
	- `http://universities.hipolabs.com/search?name={NAME}&country={COUNTRY}`


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


The Country API is a third-party API that provides information on countries all over the world. With this API, users can retrieve various data points about countries, such as their name, capital, and official languages. This API can be useful in a variety of applications, from educational tools to travel planning apps. The API is free to use and does not require authentication. To use the Country API, users can send HTTP requests to the API endpoints and receive JSON responses in return. The documentation for the Country API can be found on their website: https://restcountries.com/

✅ How to use: <br>
    - `https://restcountries.com/v3.1/name/{COUNTRY_NAME}` <br>
    - `https://restcountries.com/v3.1/aplha/{COUNTRY_ALPHA}` <br>


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
```

## Deployment
<br>
The service has been successfully deployed using Render. Render is a cloud computing platform that simplifies the deployment and scaling of web applications and services. Render's intuitive interface for configuring and managing infrastructure made it easy for to deploy the REST service and make it available to our users. Here is the render link: <br>
<br>
https://prog2005-universityapi.onrender.com


## Mock Endpoints 

The UniversityAPI project also includes two mock endpoints, one for universities and one for countries. These mock endpoints can be used for testing purposes and to avoid making real API calls during development. The mock endpoints return sample data in the same format as the actual API responses, allowing developers to test the code with realistic data. The sample data for the mock endpoints is stored in JSON files located in the `./internal/res` folder of the project. These mock endpoints can be accessed at: <br>
- `/mock/uni` <br>
- `/mock/country` <br>


## Testing 

The service also provide a set of tests for the code. These tests are designed to check the behavior of the code and to ensure that it meets the expected requirements. We encourage you to run these tests and contribute to the project by reporting any issues or bugs that you may find. You can run the tests by using the appropriate commands provided<br>

To run tests: <br>

```terminal
	go test ./...
```

And to see the coverage rate, use: <br>

```terminal
	go test -coverpkg=./... ./...
```


## A note on status codes used in the service:

It is important to note that receiving a 400 status code from the service does not always indicate an internal error. It is possible that the server processed the request successfully but did not have any content to return to the client. Therefore, it is recommended to always check the status message associated with the response to determine the exact nature of the error. The status message can provide valuable information to help troubleshoot the issue and determine the appropriate course of action. 


Instead of a **203**, that you would expect to get, you get a **4xx** status code:

Why? Because this is interpreted as a client error. The client did not use the endpoint correctly and got no content back. (In contrast to another interpretation, where you would get **203** because the client did nothing wrong, **we blame the client**)