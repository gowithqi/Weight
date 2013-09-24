package main

import (
	"database/sql"
	"fmt"
	"github.com/Weight/login"
	"github.com/Weight/userpage"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	//"strings"
	//"time"
)

//all the static files
func staticServe(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Path)
	file := "template/" + r.URL.Path[1:]
	//fmt.Println(file)
	http.ServeFile(w, r, file)
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Println("_ method: ", r.Method, "URL: ", r.URL.Path)
	t, _ := template.ParseFiles("template/userpage.html")
	t.Execute(w, nil)
}

func submitWeight(w http.ResponseWriter, r *http.Request) {
	var Log *log.Logger
	Log = log.New(os.Stdout, "Main.submitWeight: ", log.LstdFlags)
	if r.Method == "POST" {
		db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")

		_user_weight, err := strconv.ParseFloat(r.FormValue("weight"), 32)
		user_weight := float32(_user_weight)
		if err != nil {
			Log.Println("format of weight is wrong, weight is ", r.FormValue("weight"))
		}
		_user_id, err := strconv.ParseInt(r.FormValue("user_id"), 10, 0)
		user_id := int(_user_id)
		if err != nil {
			Log.Println("format of useid is wrong, useid is ", r.FormValue("user_id"))
		}

		user, _ := userpage.GetUserWithId(db, user_id)
		user.RecordWeight(db, user_weight)

		login.OpenLoginHTML(w, user)

		Log.Println("User: ", user.Name, ", successfully record a weight, ", user_weight, "kg")
	}
}

func requestWeightData(w http.ResponseWriter, r *http.Request) {
	var Log *log.Logger
	Log = log.New(os.Stdout, "Main.requestWeightData: ", log.LstdFlags)
	if r.Method == "GET" {
		db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")

		_user_id, _ := strconv.ParseInt(r.FormValue("user_id"), 10, 0)
		user_id := int(_user_id)

		user, _ := userpage.GetUserWithId(db, user_id)

		user.RequestWeightData(db, r.FormValue("start_date"))

		var data string = ""
		for i := 0; i < len(user.HistoryWeight); i++ {
			data += user.HistoryWeight[i].Date
			data += "-" + strconv.FormatFloat(float64(user.HistoryWeight[i].Weight), 'f', 1, 32)
			if i != len(user.HistoryWeight)-1 {
				data += "\n"
			}
		}
		Log.Println("user id: ", r.FormValue("user_id"), "    start date: ", r.FormValue("start_date"))
		t := template.New("response")
		t, _ = t.Parse(data)
		t.Execute(w, nil)

		Log.Println("User: ", user.Name, ", successfully request weight history")
	}
}

func requestWeightDataFALSE(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	//db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")
	http.HandleFunc("/", login.Login)
	http.HandleFunc("/static/", staticServe)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/weight/submit", submitWeight)
	http.HandleFunc("/weight/requestweightdata", requestWeightData)
	http.HandleFunc("/user", user)
	http.HandleFunc("/weight/requestweightdataFALSE", requestWeightDataFALSE)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
