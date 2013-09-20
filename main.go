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
	"strings"
	"time"
)

func weight(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET")
		t, _ := template.ParseFiles("template/weight.html")
		t.Execute(w, nil)
		return
	}
	fmt.Println("username: ", r.FormValue("username"))
	fmt.Println("weight: ", r.FormValue("weight"))

	db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")
	fmt.Println("sql.Open")

	query := fmt.Sprintf("select * from User where username=\"%s\"", r.FormValue("username"))
	fmt.Println("query: ", query)
	row, _ := db.Query(query)

	if !row.Next() {
		stmt, _ := db.Prepare("insert User set username=?")
		stmt.Exec(r.FormValue("username"))
		fmt.Println("insertn a new user")
	}

	row, _ = db.Query(query)
	fmt.Println("query")

	for row.Next() {
		var id int
		var username string
		var pawd string
		row.Scan(&id, &username, &pawd)

		stmt, _ := db.Prepare("insert WeightRecord set user_id=?, date=?, weight=?")
		w, _ := strconv.ParseFloat(r.FormValue("weight"), 32)
		stmt.Exec(id, time.Now().String(), w)

		query = fmt.Sprintf("select date, weight from WeightRecord where user_id=%d", id)
		rows, _ := db.Query(query)

		fmt.Println("User ", username, "'s record:")
		for rows.Next() {
			var date string
			var weight float32
			rows.Scan(&date, &weight)

			fmt.Println("date: ", date, " weight: ", weight)
		}
		fmt.Println("")
	}
	t, _ := template.ParseFiles("template/weight.html")
	t.Execute(w, nil)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("* method: ", r.Method)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("key: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello, astaxie!")
}

//all the static files
func staticServe(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	file := "template/" + r.URL.Path[1:]
	fmt.Println(file)
	http.ServeFile(w, r, file)
}

/*
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("_ method: ", r.Method, "URL: ", r.URL.Path)
	if r.Method == "POST" {
		r.ParseForm()
		//have got a username
		fmt.Println("username: ", r.FormValue("username"))
		fmt.Println("password: ", r.FormValue("password"))
	} else {
		fmt.Println("__", r.FormValue("user"))
	}
	t, _ := template.ParseFiles("template/login.html")
	t.Execute(w, nil)
}
*/

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Println("_ method: ", r.Method, "URL: ", r.URL.Path)
	t, _ := template.ParseFiles("template/userpage.html")
	t.Execute(w, nil)
}

func submitWeight(w http.ResponseWriter, r *http.Request) {
	var Log *log.Logger
	Log = log.New(os.Stdout, "submitWeight: ", log.LstdFlags)
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

		Log.Println("successfully record a weight")
	}
}

func requestWeightData(w http.ResponseWriter, r *http.Request) {
	var Log *log.Logger
	Log = log.New(os.Stdout, "submitWeight: ", log.LstdFlags)
	if r.Method == "POST" {
		db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")

		_user_id, _ := strconv.ParseInt(r.FormValue("user_id"), 10, 0)
		user_id := int(_user_id)

		user, _ := userpage.GetUserWithId(db, user_id)

		user.RequestWeightData(db, r.FormValue("start_date"))
		login.OpenLoginHTML(w, user)

		Log.Println("successfully request weight history")
	}
}
func requestWeightData1(w http.ResponseWriter, r *http.Request) {
	var Log *log.Logger
	Log = log.New(os.Stdout, "submitWeight: ", log.LstdFlags)
	if r.Method == "POST" {
		db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")

		_user_id, _ := strconv.ParseInt(r.FormValue("user_id"), 10, 0)
		user_id := int(_user_id)

		user, _ := userpage.GetUserWithId(db, user_id)

		user.RequestWeightData(db, r.FormValue("start_date"))
		t := template.New("try")
		t, _ = t.Parse("FUCK") //there is some code
		t.Execute(w, nil)

		Log.Println("successfully request weight history")
	}
}
func main() {
	//db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/static/", staticServe)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/weight/submit", submitWeight)
	http.HandleFunc("/weight/requestweightdata", requestWeightData)
	http.HandleFunc("/user", user)
	http.HandleFunc("/weight/requestweightdata1", requestWeightData1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
