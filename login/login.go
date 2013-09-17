package login

import (
	"database/sql"
	//"fmt"
	"github.com/Weight/userpage"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	//"strconv"
	//"strings"
	//"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var Log *log.Logger
	Log = log.New(os.Stdout, "login: ", log.LstdFlags)

	switch r.Method {
	case "GET":
		Log.Println("receive a Get Request")
		//put a login html
		t, _ := template.ParseFiles("template/login.html")
		t.Execute(w, nil)
	case "POST":
		Log.Println("receive a POST Request")

		db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")
		row, _ := db.Query("select id, password from User where username=?", r.FormValue("username"))

		//if there are not such a user
		if !row.Next() {
			//beta version will create
			//formal version will put a sign up html

			//add a user currently
			stmt, _ := db.Prepare("insert User set username=?, password=?")
			stmt.Exec(r.FormValue("username"), r.FormValue("password"))
		}

		var user_id int
		var password string
		row.Scan(&user_id, &password)

		user, res := userpage.GetUserWithName(db, r.FormValue("username"), r.FormValue("password"))
		switch res {
		case "Success":
			t, _ := template.ParseFiles("template/userpage.html") //there is some code
			t.Execute(w, user)
		case "PasswordWrong":
			t, _ := template.ParseFiles("template/login.html")
			t.Execute(w, nil)
		}
	}
}

func OpenLoginHTML(w http.ResponseWriter, user userpage.User) {
	t, _ := template.ParseFiles("template/userpage.html") //there is some code
	t.Execute(w, user)
}
