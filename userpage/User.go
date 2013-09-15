package userpage 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"time"
)

//var db *sql.DB
//db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")

type User struct {
	id int
	name string
	status string
	weight_delta float32
}

func (user User) GetId() (int) {
	return user.id
}

func (user User) GetName() (string) {
	return user.name
}

func (user User) GetStatus() (string) {
	return user.status
}

func (user User) GetWeightDelta() (float32) {
	return user.weight_delta
}

func GetUserWithName(db *sql.DB, username string, password string) (User, string) {
	var user User

	Log := log.New(os.Stdout, "GetUser: ", log.LstdFlags)
	row, _ := db.Query("select * from User where username=?", username)
	if !row.Next() {
		Log.Println("there is not such a user")
		return nil, "NotExist"
	}

	var passwordC string //correct password

	row.Scan(&user.id, &user.name, &passwordC, &user.status, &user.weight_delta)

	if password != passwordC {
		Log.Println("password is not correct")
		return nil, "PasswordWrong"
	}
	else {
		Log.Println("get a User successfully")
		return user, "Success"
	}
}

func GetUserWithId(db *sql.DB, id int) (User, string) {
	var user User

	Log := log.New(os.Stdout, "GetUser: ", log.LstdFlags)
	row, _ := db.Query("select id, username, status, weight_delta from User where id=?", id)
	if !row.Next() {
		Log.Println("there is not such a user")
		return nil, "NotExist"
	}

	row.Scan(&user.id, &user.name, &user.status, &user.weight_delta)
	Log.Println("get a User successfully")
	return user, "Success"
}
