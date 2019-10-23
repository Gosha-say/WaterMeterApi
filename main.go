package main

import (
	"WaterMeterApi/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

//TODO: Error handler

func WaterRouterHandler(w http.ResponseWriter, r *http.Request) {
	params := new(models.WaterParams)
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Form)
	params.Id = r.FormValue("id")

	cval1, err := strconv.ParseInt(r.FormValue("cval1"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	cval2, err := strconv.ParseInt(r.FormValue("cval2"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	hval1, err := strconv.ParseInt(r.FormValue("hval1"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	hval2, err := strconv.ParseInt(r.FormValue("hval2"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	date, err := strconv.ParseInt(r.FormValue("dt"), 10, 64)
	if err != nil {
		date = time.Now().Unix()
	}
	power, err := strconv.ParseInt(r.FormValue("vp"), 10, 64)
	params.WCold1 = cval1
	params.WCold2 = cval2
	params.WHot1 = hval1
	params.WHot2 = hval2
	params.Date = int32(date)
	params.Power = power
	valid, comment := params.Validate()
	if !valid {
		fmt.Println(comment)
	}
	fmt.Println(params)
	a, err := fmt.Fprintf(w, "Version: %s", "0.2")
	if err != nil {
		fmt.Println("Error: ", a)
	}
}

func main() {
	http.HandleFunc("/", WaterRouterHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
