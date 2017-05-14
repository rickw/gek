package excelkata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
    toURL   = "http://afitnerd.herokuapp.com/v1/excel_kata/to_col_notation/%d"
    fromURL = "http://afitnerd.herokuapp.com/v1/excel_kata/from_col_notation/%s"
)

func TestToCol(t *testing.T) {
	for toTest := 1; toTest < 100; toTest++ {
		result := ToColNotation(toTest)
		expected := toCol(toTest)

		if result != expected {
			t.Errorf("ToCol(%d), returned(%s), expected(%s)", toTest, result, expected)
		}
	}
}

func TestFromCol(t *testing.T) {
	for toTest := 1; toTest < 100; toTest++ {
		toTestString := toCol(toTest)
		result := FromColNotation(toTestString)
		expected := fromCol(toTestString)

		if result != expected {
			t.Errorf("FromCol(%s), returned(%d), expected(%d)", toTestString, result, expected)
		}
	}
}

func toCol(number int) string {
	reqURL := fmt.Sprintf(toURL, number)
	resp, err := http.Get(reqURL)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println(err)
	}
	return dat["result"].(string)
}

func fromCol(str string) int {
	reqURL := fmt.Sprintf(fromURL, str)
	resp, err := http.Get(reqURL)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println(err)
	}
	return int(dat["result"].(float64))
}
