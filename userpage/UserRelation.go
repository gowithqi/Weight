package userpage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func (user User) AddFriend(db *sql.DB, shou User) {
	stmt, _ := db.Prepare("insert UserRelation set user1_id=?, user2_id=?")
	stmt.Exec(user.id, shou.GetId())
}

func (user User) DeleteFriend(db *sql.DB, shou User) {
	stmt, _ := db.Prepare("delete from UserRelation where (user1_id=? and user2_id=?) or (user1_id=? and user2_id=?)")
	stmt.Exec(user.id, shou.GetId(), shou.GetId(), user.id)
}

func (user *User) GetAllFriend(db *sql.DB) {
	user.friends = make([]User, 0)

	Log := log.New(os.Stdout, "User.GetAllFriend: ", log.LstdFlags)

	rows, _ := db.Query("select user1_id, user2_id from UserRelation where (user1_id=? or user2_id=?)", user.id, user.id)

	for rows.Next() {
		var user1_id int
		var user2_id int
		rows.Scan(&user1_id, &user2_id)

		var newfriend User
		var res string
		if user1_id == user.id {
			newfriend, res = GetUserWithId(db, user2_id)
		} else {
			newfriend, res = GetUserWithId(db, user1_id)
		}
		if res == "NotExist" {
			Log.Panicln("there is no such a user id")
		}
		user.friends = append(user.friends, newfriend)
	}
}

func (user User) IsFriendOf(db *sql.DB, shou User) bool {
	row, _ := db.Query("select * from UserRelation where (user1_id=? and user2_id=?) or (user1_id=? and user2_id=?)", user.id, shou.GetId(), shou.GetId(), user.id)
	return row.Next()
}
