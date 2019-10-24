package main

import (
	"WaterMeterApi/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
	"time"
)

//TODO: Error handler

func main() {

	http.HandleFunc("/", WaterRouterHandler)
	http.HandleFunc("/favicon.ico", Favicon)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func WaterRouterHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/w_meter")
	if db == nil {
		panic("DB in nil")
	}
	fmt.Println("DB - OK!")

	params := new(models.WaterParams)
	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	params.Id = r.FormValue("id")
	if params.Id == "" {
		fmt.Println("Error: empty ID")
		return
	}

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

	js, err := json.Marshal(params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	a, err := w.Write(js)

	if err != nil {
		fmt.Println("Error: 12: %w", a)
	}
	fmt.Print("New request: ")
	fmt.Println(string(js))

	tm := time.Unix(int64(params.Date), 0)

	sql, err := db.Exec("insert into meters_data (MeterId, WCold1, WCold2, WHot1, WHot2, Power, Date) values (?,?,?,?,?,?,?)",
		params.Id, params.WCold1, params.WCold2, params.WHot1, params.WHot2, params.Power, tm)
	if err != nil {
		fmt.Println(err)
	}

	in, err := sql.LastInsertId()
	fmt.Println("New record id:", in)
}

func Favicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
