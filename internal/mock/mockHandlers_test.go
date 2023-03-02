package mock

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"uniapi/internal/constants"
)

func TestCountryMockHandler(t *testing.T) {
    // Testing a get request on local host 
    getRequest, _ := http.NewRequest("GET", constants.MOCK_COUNTRY_API_URL, nil)
    response := httptest.NewRecorder()
	//Executing the handler 
    CountryMockHandler(response, getRequest)
    resultGet := response.Result()
    defer resultGet.Body.Close()

	//Error if not implemented or not correct 
    if resultGet.StatusCode != http.StatusOK {
        t.Error("Test case on GET failed, should be 200")
    }
    expected1 := "application/json"
    resultGetHeader := resultGet.Header.Get("content-type")
    if resultGetHeader != expected1 {
        t.Errorf("Test case failed on GET: wrong header information")
    }

    // Test case 2: POST request
    postRequest, _ := http.NewRequest("POST", constants.MOCK_COUNTRY_API_URL, nil)
    postResponse := httptest.NewRecorder()
    CountryMockHandler(postResponse, postRequest)
    resultPost := postResponse.Result()
    defer resultPost.Body.Close()
    if resultPost.StatusCode != http.StatusNotImplemented {
        t.Errorf("Test case POST failed: Not marked as not implemented")
    }
}

func TestUniMockHandler(t *testing.T) {
    // Test case 1: GET request
    req1, _ := http.NewRequest("GET", constants.MOCK_UNI_API_URL, nil)
    rr1 := httptest.NewRecorder()
    UniMockHandler(rr1, req1)
    result1 := rr1.Result()
    defer result1.Body.Close()
    if result1.StatusCode != http.StatusOK {
        t.Errorf("Test case 1 failed: expected status OK, but got %v", result1.StatusCode)
    }
    expected1 := "application/json"
    result1Header := result1.Header.Get("content-type")
    if result1Header != expected1 {
        t.Errorf("Test case 1 failed: expected header %v, but got %v", expected1, result1Header)
    }

    // Test case 2: POST request
    req2, _ := http.NewRequest("POST", constants.MOCK_UNI_API_URL, nil)
    rr2 := httptest.NewRecorder()
    UniMockHandler(rr2, req2)
    result2 := rr2.Result()
    defer result2.Body.Close()
    if result2.StatusCode != http.StatusNotImplemented {
        t.Errorf("Test case 2 failed: expected status not implemented, but got %v", result2.StatusCode)
    }
}
