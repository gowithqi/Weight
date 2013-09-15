package userpage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func RequestWeightData (db *sql.DB, user_id int, start_date string) (dates []string, times [] string, weights []float32) {
	rows, _ := db.Query("select date, time, weight from WeightRecord where user_id=?, date>=?", user_id, start_date)

	Log := log.New(os.Stdout, "RequestWeightData: ", log.LstdFlags)

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
}
