package userpage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

func (user *User) RecordWeight(db *sql.DB, user_weight float32) {
	Log := log.New(os.Stdout, "User.RecordWeight: ", log.LstdFlags)

	stmt, _ := db.Prepare("insert WeightRecord set user_id=?, date=?, time=?, weight=?")
	stmt.Exec(user.Id,
		time.Now().Format("2006-01-02"),
		time.Now().Format("3:04PM"),
		user_weight)
	Log.Println("record user weight successfully")

	d := time.Hour * (-24)
	rows, _ := db.Query("select weight from WeightRecord where date=? and user_id=?", time.Now().Add(d).Format("2006-01-02"), user.Id)
	weight_delta := -1.0
	for rows.Next() {
		rows.Scan(&weight_delta)
	}
	stmt, _ = db.Prepare("update User set status=?, weight_delta=? where user_id=?")
	stmt.Exec(time.Now().Format("2006-01-02"),
		weight_delta,
		user.Id)
	Log.Println("update user status")
}

func (user *User) RequestWeightData(db *sql.DB, start_date string) {
	rows, _ := db.Query("select date, weight from WeightRecord where user_id=? and date>=?", user.Id, start_date)

	Log := log.New(os.Stdout, "User.RequestWeightData: ", log.LstdFlags)
	Log.Println("have got requested user weight data from database")

	user.HistoryWeight = make([]WeightAndDate, 0)

	for rows.Next() {
		var weightanddate WeightAndDate
		rows.Scan(&weightanddate.Date, &weightanddate.Weight)
		fmt.Println("Date: ", weightanddate.Date, " Weight: ", weightanddate.Weight)
		user.HistoryWeight = append(user.HistoryWeight, weightanddate)
	}
}
