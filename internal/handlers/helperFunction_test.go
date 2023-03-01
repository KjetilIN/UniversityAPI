package handlers

import (
	"encoding/json"
	"testing"
	"uniapi/internal/constants"
)

func TestGetFromUniApi_Positive(t *testing.T) {

	//Search word
	search := "Norway"

	resp, err := getFromUniAPI(search)

	if(err != nil){
		t.Fatal("TestGetFromUniApiPositive: Error on Positive Test for ", search)
		t.Error()

	}

	if (resp == nil){
		t.Fatal("TestGetFromUniApiPositive: Error no response to ", search)
		t.Error()

	}

	//Decoding into struct 
	var uniStruct []constants.UniStruct

	// Decode struct
	decodeErr := json.NewDecoder(resp.Body).Decode(&uniStruct)
	if(decodeErr != nil){
		t.Fatal("TestGetFromUniApiPositive: Error during Positive test Decoding")
		t.Error()
	}

	if(len(uniStruct) == 0){
		t.Fatal("TestGetFromUniApiPositive: Got list of structs on Positive test")
	}

}

func TestGetFromUniApi_Negative(t *testing.T) {

	//Search word is random and no result is aspected 
	search := "swdfghjkjhgfdsasdfghjkjhgfdsasdfghjkjhgfdsasdfghjk";

	resp, _ := getFromUniAPI(search);

	//Decoding into struct 
	var uniStruct []constants.UniStruct

	// Decode struct
	decodeErr := json.NewDecoder(resp.Body).Decode(&uniStruct)
	if(decodeErr != nil){
		t.Fatal("TestGetFromUniApiNegative: Error during Negative test Decoding")
		t.Error()
	}

	if(len(uniStruct) != 0){
		t.Fatal("TestGetFromUniApiNegative: Got list of structs on Negative test")
	}

}


func TestGetCountryFromAlphaCode_Positive(t *testing.T) {
	//Code that should give positive result:
	code :="NO"

	resp, err := getCountryFromAlphaCode(code)

	//No error
	if(err != nil){
		t.Fatal("TestGetFromUniApiPositive: Error on Positive Test for ", code)
		t.Error()

	}

	if (resp != "Norway"){
		t.Fatal("TestGetFromUniApiPositive: Not correct response, got ",resp)
		t.Error()

	}
}



func TestGetCountryFromAlphaCode_Negative(t *testing.T) {
	//Code that should give Negative result:
	code :="PYTHONISBEST"

	resp, err := getCountryFromAlphaCode(code)

	//Expect no error 
	if(err != nil){
		t.Fatal("TestGetFromUniApiPositive: Error on Negative Test for ", code)
		t.Error()

	}
	
	//Aspect no response 
	if (resp != ""){
		t.Fatal("TestGetFromUniApiPositive: Not correct response, got ",resp)
		t.Error()

	}
}


func TestGetBorderCountry_Positive(t *testing.T) {
	//Country we want to get border country from 
	country:= "Norway"

	resp, err := getBorderCountry(country)

	if(err != nil){
		t.Fatal("TestGetBorderCountryPositive: Got error")
		t.Error()
	}

	if(len(resp) != 3){
		t.Fatal("TestGetBorderCountryPositive: Did not get correct length of strings")
		t.Error()
	}


}


func TestGetBorderCountry_Negative(t *testing.T) {
	//Country we want to get border country from 
	country:= "THIS_IS_NO_COUNTRY"

	resp, err := getBorderCountry(country)

	if(err != nil){
		t.Fatal("TestGetBorderCountryNegative: Did get an error")
		t.Error()
	}

	if(len(resp) != 0){
		t.Fatal("TestGetBorderCountryNegative: Got results, that are not wanted: ")
		t.Error()
	}


}

func TestGetAllFromUniAPI_Positive(t *testing.T) {
	//Keywords that should get results 
	middle :="science"
	country :="Norway"

	resp, err := getAllFromUniAPI(country,middle)

	if (err != nil){
		t.Fatal("TestGetAllFromUniAPI_Positive: Error on method, should be none.")
		t.Error()
	}

	//Decoding into struct 
	var uniStruct []constants.UniStruct

	// Decode struct
	decodeErr := json.NewDecoder(resp.Body).Decode(&uniStruct)
	if(decodeErr != nil){
		t.Fatal("TestGetAllFromUniAPI_Positive: Error during Positive test Decoding")
		t.Error()
	}

	if(len(uniStruct) == 0){
		t.Fatal("TestGetAllFromUniAPI_Positive: No result in struct after decoding")
		t.Error()
	}


}

func TestGetAllFromUniAPI_Negative(t *testing.T) {
	//Keywords that should get results 
	middle :="A_MASSIVE_WORD_THAT_DONT_GIVE_RESULTS"
	country :="Norway"

	resp, err := getAllFromUniAPI(country,middle)

	if (err != nil){
		t.Fatal("TestGetAllFromUniAPI_Negative: Error on method, should be none.")
		t.Error()
	}

	//Decoding into struct 
	var uniStruct []constants.UniStruct

	// Decode struct
	decodeErr := json.NewDecoder(resp.Body).Decode(&uniStruct)
	if(decodeErr != nil){
		t.Fatal("TestGetAllFromUniAPI_Negative: Error during Negative test Decoding")
		t.Error()
	}

	if(len(uniStruct) != 0){
		t.Fatal("TestGetAllFromUniAPI_Negative: Got result after decoding, expected none.")
		t.Error()
	}

}