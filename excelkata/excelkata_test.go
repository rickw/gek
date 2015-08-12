package excelkata

import (
	"net/http"
	"testing"
)

const (
	TO_URL   = "http://code-challenge.afitnerd.com/v1/excel_kata/to_col_notation/%i"
	FROM_URL = "http://code-challenge.afitnerd.com/v1/excel_kata/from_col_notation/%s"
)

func toCol(number int) string {
	req_url := sprintf(TO_URL, number)
	resp, err := http.Get(req_url)
}

func fromCol(str string) int {
	req_url := sprintf(FROM_URL, str)
	resp, err := http.Get(req_url)
}
