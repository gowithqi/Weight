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
	stmt, _ := db.Prepare("insert WeightRecord set user_id=?, date=?, time=?, weight=?")
	stmt.Exec(user.Id,
		time.Now().Format("2006-01-02"),
		time.Now().Format("3:04PM"),
		user_weight)

	Log := log.New(os.Stdout, "User.RecordWeight: ", log.LstdFlags)

	Log.Println("record user weight successfully")
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
