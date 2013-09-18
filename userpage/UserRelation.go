package userpage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func (user User) AddFriend(db *sql.DB, shou User) {
	stmt, _ := db.Prepare("insert UserRelation set user1_id=?, user2_id=?")
	stmt.Exec(user.Id, shou.GetId())
}

func (user User) DeleteFriend(db *sql.DB, shou User) {
	stmt, _ := db.Prepare("delete from UserRelation where (user1_id=? and user2_id=?) or (user1_id=? and user2_id=?)")
	stmt.Exec(user.Id, shou.GetId(), shou.GetId(), user.Id)
}

func (user *User) GetAllFriend(db *sql.DB) {
	user.Friends = make([]User, 0)

	Log := log.New(os.Stdout, "User.GetAllFriend: ", log.LstdFlags)

	rows, _ := db.Query("select user1_id, user2_id from UserRelation where (user1_id=? or user2_id=?)", user.Id, user.Id)

	for rows.Next() {
		var user1_id int
		var user2_id int
		rows.Scan(&user1_id, &user2_id)

		var newfriend User
		var res string
		if user1_id == user.Id {
			newfriend, res = GetUserWithId(db, user2_id)
		} else {
			newfriend, res = GetUserWithId(db, user1_id)
		}
		if res == "NotExist" {
			Log.Panicln("there is no such a user id")
		}
		user.Friends = append(user.Friends, newfriend)
	}
}

func (user User) IsFriendOf(db *sql.DB, shou User) bool {
	row, _ := db.Query("select * from UserRelation where (user1_id=? and user2_id=?) or (user1_id=? and user2_id=?)", user.Id, shou.GetId(), shou.GetId(), user.Id)
	return row.Next()
}

func (user User) SortFriendsWeight() {
	for i := 0; i < len(user.Friends); i++ {
		for j := i + 1; j < len(user.Friends); i++ {
			if user.Friends[i].Weight_delta < user.Friends[i].Weight_delta {
				user.Friends[i], user.Friends[j] = user.Friends[j], user.Friends[i]
			}
		}
	}
}
