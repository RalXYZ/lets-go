package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type bar struct {
	id  string
	age int
}

func dbConn() *sql.DB {
	dbDriver := "mysql"
	dbUser := "admin"
	dbPasswd := "admin"
	dbName := "foo"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPasswd+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbCreate(x *bar) *Content {
	db := dbConn()
	defer db.Close()
	result, err := db.Exec("INSERT INTO bar(id, age) VALUES(?, ?)", x.id, x.age)
	if err != nil {
		panic(err.Error())
	}
	uid, _ := result.LastInsertId()
	return &Content{uid, x.id, x.age}
}

func dbRetrieve(uid int64) *Content {
	db := dbConn()
	defer db.Close()

	stmt, err := db.Prepare("SELECT age, id FROM bar WHERE uid = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	rows.Next()
	var (
		age int
		id  string
	)
	err = rows.Scan(&age, &id)
	if err != nil {
		panic(err.Error())
	}

	return &Content{uid, id, age}
}

func dbUpdate(dst *Content) *Content {
	db := dbConn()
	defer db.Close()
	_, err := db.Exec("UPDATE bar SET id = ?, age = ? WHERE uid = ?", dst.Id, dst.Age, dst.Uid)
	if err != nil {
		panic(err.Error())
	}
	return dst
	// num, _ := result.RowsAffected()
}

func dbDelete(uid int64) {
	db := dbConn()
	defer db.Close()
	_, err := db.Exec("DELETE FROM bar WHERE uid = ?", uid)
	if err != nil {
		panic(err.Error())
	}
	// num, _ := result.RowsAffected()
}
