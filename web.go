package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rickw/gek/excelkata"
	"net/http"
	"os"
	"strconv"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/to_col_notation/:num", ToCol)
	router.GET("/from_col_notation/:str", FromCol)
	router.GET("/to_timmys_brain", TimmyBrain)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		panic(err)
	}
}

// added as a joke for "dogeared"
func TimmyBrain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Timmy Has No Brain!")
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "index.html")
}

func ToCol(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// ps.ByName("num")
	num, _ := strconv.Atoi(ps.ByName("num"))
	json.NewEncoder(w).Encode(map[string]string{"result": excelkata.ToColNotation(num)})
}

func FromCol(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// ps.ByName("str")
	json.NewEncoder(w).Encode(map[string]int{"result": excelkata.FromColNotation(ps.ByName("str"))})
}
