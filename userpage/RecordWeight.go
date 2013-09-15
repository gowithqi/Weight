package userpage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

//maybe this function will return something 
//but it is not decided
func RecordWeight (db *sql.DB, user_id int, user_weight float32) {
		stmt, _ := db.Prepare("insert WeightRecord set user_id=?, date=?, time=?, weight=?")
		stmt.Exec(user_id,
			  time.Now().Format("2006-01-02"),
			  time.Now().Format("3:04PM"),
			  user_weight)

		Log := log.New(os.Stdout, "RecordWeight: ", log.LstdFlags)

		Log.Println("record user weight successfully")
}
