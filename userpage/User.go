package userpage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	//"time"
)

//var db *sql.DB
//db, _ := sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/myGoWebDatabase")

type WeightAndDate struct {
	Weight float32
	Date   string
}

type User struct {
	Id             int
	Name           string
	Status         string
	Weight_delta   float32
	Current_weight float32
	Friends        []User

	HistoryWeight []WeightAndDate
}

func (user User) GetId() int {
	return user.Id
}

func (user User) GetName() string {
	return user.Name
}

func (user User) GetStatus() string {
	return user.Status
}

func (user User) GetWeightDelta() float32 {
	return user.Weight_delta
}

func GetUserWithName(db *sql.DB, username string, password string) (User, string) {
	var user User

	Log := log.New(os.Stdout, "GetUser: ", log.LstdFlags)
	row, _ := db.Query("select * from User where username=?", username)
	if !row.Next() {
		Log.Println("there is not such a user")
		return user, "NotExist"
	}

	var passwordC string //correct password

	row.Scan(&user.Id, &user.Name, &passwordC, &user.Status, &user.Weight_delta)

	if password != passwordC {
		Log.Println("password is not correct")
		return user, "PasswordWrong"
	} else {
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
		return user, "NotExist"
	}

	row.Scan(&user.Id, &user.Name, &user.Status, &user.Weight_delta)
	Log.Println("get a User successfully")
	return user, "Success"
}
