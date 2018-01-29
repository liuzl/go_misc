package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/durecorder")
	if err != nil {
		glog.Fatal(err)
	}
	/*
		result, err := db.Exec(
			"INSERT INTO revival_card(uid,code,count) VALUES (?,?,?)", 1, "abc", 5)
		if err != nil {
			glog.Fatal(err)
		}
		glog.Infof("%+v\n", result)
	*/
	rows, err := db.Query("SELECT uid, count FROM revival_card WHERE count > ?", 1)
	if err != nil {
		glog.Fatal(err)
	}
	for rows.Next() {
		var uid, count int
		if err := rows.Scan(&uid, &count); err != nil {
			glog.Fatal(err)
		}
		glog.Infof("%d is %d\n", uid, count)
	}
	if err := rows.Err(); err != nil {
		glog.Fatal(err)
	}
	r, err := db.Exec("UPDATE revival_card set count=count-1 where count>0 and uid=?", 1)
	if err != nil {
		glog.Fatal(err)
	}
	affected, err := r.RowsAffected()
	if err != nil {
		glog.Fatal(err)
	}
	glog.Infof("r:%+v\n", affected)
}
