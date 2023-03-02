package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"uniapi/internal/constants"
)

func TestGetFromUniApi_Positive(t *testing.T) {

	//Search word
	search := "Norway"

	resp, err := getUniversitiesWithName(search)

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

	//Search word is random and no result is aspects 
	search := "swdfghjkjhgfdsasdfghjkjhgfdsasdfghjkjhgfdsasdfghjk";

	resp, _ := getUniversitiesWithName(search);

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

	resp, err := getCountryNameFromAlphaCode(code)

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

	resp, err := getCountryNameFromAlphaCode(code)

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

	resp, err := getBorderCountries(country)

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

	resp, err := getBorderCountries(country)

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

	resp, err := getUniversitiesWithNameAndMiddle(country,middle)

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

	resp, err := getUniversitiesWithNameAndMiddle(country,middle)

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


func TestRemoveEmptyStrings(t *testing.T) {
    // Testing if given empty list of strings 
    emptyList := []string{}
    expected := []string{}
    result := removeEmptyStrings(emptyList)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, but got %v", expected, result)
    }

    // Test with some empty strings 
    listWithEmpty := []string{"", "hello", "", "world", ""}
    expected = []string{"hello", "world"}
    result = removeEmptyStrings(listWithEmpty)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, but got %v", expected, result)
    }

    // Test with no empty strings 
    noEmptyList := []string{"hello", "world"}
    expected = []string{"hello", "world"}
    result = removeEmptyStrings(noEmptyList)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, but got %v", expected, result)
    }
}


func TestIsOfValidLength(t *testing.T) {
    // Test case with valid length 
    strList1 := []string{"hello", "world"}
    required1 := 2
    message1 := "Invalid request path."
    w1 := httptest.NewRecorder()
    result1 := isOfValidLength(strList1, required1, message1, w1)
    if result1 != true {
        t.Errorf("Test case 1 failed: expected true, but got %v", result1)
    }


    // Test case with empty strings
    strList3 := []string{"", "world", ""}
    required3 := 2
    message3 := "Invalid request path."
    w3 := httptest.NewRecorder()
    result3 := isOfValidLength(strList3, required3, message3, w3)
    if result3 != false {
        t.Errorf("Test case 3 failed: expected false, but got %v", result3)
    }
}


func TestIsCorrectRequestMethod(t *testing.T) {
    // Test for a GET request 
    req1, _ := http.NewRequest("GET", "/", nil)
    result1 := isCorrectRequestMethod(req1)
    if result1 != true {
        t.Error("Test case with GET failed!")
    }

    // Test for a POST request 
    req2, _ := http.NewRequest("POST", "/", nil)
    result2 := isCorrectRequestMethod(req2)
    if result2 != false {
        t.Error("Test case with POST failed!")
    }
}


func TestReplaceSpaces(t *testing.T) {
    // Test case with a single space
    url1 := "http://localhost/hello%20world"
    expected1 := "http://localhost/hello world"
    result1 := replaceSpaces(url1)
    if result1 != expected1 {
        t.Error("Test case with single space failed")
    }

    // Test case with multiple spaces
    url2 := "http://localhost/hello%20world%20and%20universe"
    expected2 := "http://localhost/hello world and universe"
    result2 := replaceSpaces(url2)
    if result2 != expected2 {
        t.Error("Test case with multiple spaces failed")
    }

    // Test case with no spaces
    url3 := "http://localhost/helloworld"
    expected3 := "http://localhost/helloworld"
    result3 := replaceSpaces(url3)
    if result3 != expected3 {
        t.Errorf("Test case with no spaces spaces failed")
    }
}
