package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/to_col_notation/:num", ToCol)
	router.GET("/from_col_notation/:str", FromCol)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "index.html")
}

func ToCol(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// ps.ByName("num")
	fmt.Fprintf(w, "sorry, %s!\n", ps.ByName("num"))
}

func FromCol(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// ps.ByName("str")
	fmt.Fprintf(w, "sorry, %s!\n", ps.ByName("str"))
}
