package database

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gu2list/ent/users"
	"time"

	//"gorm.io/gorm"
	//_ "github.com/mattn/go-sqlite3"
	"gu2list/ent"
	"gu2list/ent/schema"
	"log"
)

//type Client *ent.Client

var (
	sql *ent.Client
)

/*
		   	Choose Database on your own
			Here is Example Using Mysql Database.
	        Documents: https://entgo.io/docs/getting-started
*/
func CreateClient(user, pass, db, host string, port int) (client *ent.Client, err error) {
	client, err = ent.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True", user, pass, host, port, db))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	//defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	sql = client
	return sql, err
}

func UpdateUser(userID uint64, Roles *schema.Roles, Manager bool) (err error) {
	User, err := sql.Users.Query().
		Where(users.UserID(userID)).
		Only(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		err = sql.Users.Create().
			SetUserID(userID).
			SetRoles(Roles).
			SetManager(Manager).
			SetUpdateAt(time.Now().Unix()).
			OnConflict().
			UpdateUpdateAt().
			UpdateRoles().
			UpdateManager().
			Exec(context.Background())
	} else {
		err = User.Update().
			SetUserID(userID).
			SetRoles(Roles).
			SetManager(Manager).
			SetUpdateAt(time.Now().Unix()).Exec(context.Background())
	}
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserLogs(userID uint64, Events string, Accept bool, timeAt int64) (err error) {
	err = sql.DefautlDB.Create().
		SetUserID(userID).
		SetLogs(Events).
		SetAccept(Accept).
		SetTimeAt(timeAt).
		//OnConflict().
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func IsManager(UserID uint64) (result bool) {
	Users, err := sql.Users.Query().Where(users.UserID(UserID)).Only(context.Background()) //.Where(users.UserID(UserID)).Only(context.Background())
	if err != nil {
		panic(err)
	}

	if Users.Manager {
		result = true
	}
	return result
}
