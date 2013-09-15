package userpage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func (user *User) RecordWeight (db *sql.DB, user_weight float32) {
	stmt, _:= db.Prepare("insert WeightRecord set user_id=?, date?, time=?, weight=?")
	stmt.Exec(user.id,
		  time.Now().Format("2006-01-02"),
		  time.Now().Format("3:04PM"),
		  user_weight)

	Log := log.New(os.Stdout, "User.RecordWeight: ", log.LstdFlags)

	Log.Println("record user weight successfully")
}

func (user User) RequestWeightData(db *sql.DB, start_date string) ([]string, []string, []float32) {
	rows, _ := db.Query("select date, time, weight from WeightRecord where user_id=?, date>=?", user.id, start_date)

	Log:= log.New(os.Stdout, "User.RequestWeightData: ", log.LstdFlags)
	Log.Println("have got requested user weight data from database")

	dates := make([]string)
	times := make([]string)
	weights := make([]float32)

	for rows.Next() {
		var date string
		var time string
		var weight float32
		rows.Scan(&date, &time, &weight)
		dates = append(dates, date)
		times = append(times, time)
		weights = append(weights, weight)
	}

	return dates, times, weights
}


