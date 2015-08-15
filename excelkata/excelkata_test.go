package excelkata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	TO_URL   = "http://code-challenge.afitnerd.com/v1/excel_kata/to_col_notation/%d"
	FROM_URL = "http://code-challenge.afitnerd.com/v1/excel_kata/from_col_notation/%s"
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
	req_url := fmt.Sprintf(TO_URL, number)
	resp, err := http.Get(req_url)
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
	req_url := fmt.Sprintf(FROM_URL, str)
	resp, err := http.Get(req_url)
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
