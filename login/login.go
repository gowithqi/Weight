package login

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"os"
)

func login(w http.ResponseWriter, r *http.Request) {
	var Log *log.Logger
	Log = log.New(os.Stdout, "login: ", log.LstdFlags)

	switch r.Method {
	case "GET":
		Log.Println("receive a Get Request")
		//put a login html
		//t, _ := template.ParseFiles()    lack of a html
		//t.Execute(w, nil)
	case "POST":
		Log.Println("receive a POST Request")

		db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")
		row, _:= db.Query("select id, password from User where username=?", r.FormValue("username"))
		//if there are not such a user
		if !row.Next() {
			//beta version will create 
			//formal version will put a sign up html
		}

		var user_id int
		var password string
		row.Scan(&user_id, &password)

		if password == r.FormValue("password") { //correct password
			//put a user page
			//t, _ := template.ParseFiles()  lack of a html
			//t, Excute(w, nil)
		} else {                                 //wrong password
			//send a alert
			//
		}
	}
}

