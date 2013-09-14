package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"log"
	"time"
	"strconv"
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
	row, _:= db.Query(query)

	if  !row.Next() {
		stmt, _  := db.Prepare("insert User set username=?")
		stmt.Exec(r.FormValue("username"))
		fmt.Println("insertn a new user")
	}

	row, _= db.Query(query)
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
		rows, _:=db.Query(query)

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
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("key: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello, astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("_ method: ", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
	 t, _ := template.ParseFiles("template/login.html")
	 t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/weight", weight)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
