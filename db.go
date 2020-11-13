package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os"
)

type bar struct {
	id  string
	age int
}

type Login struct {
	DbUser   string `json:"dbUser"`
	DbPasswd string `json:"dbPasswd"`
}

func uidExists(uid int64) bool {
	db := dbConn()
	defer db.Close()
	rows, err := db.Query("SELECT COUNT(uid) FROM bar WHERE uid = ?", uid)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	rows.Next()
	var matchNum int
	err = rows.Scan(&matchNum)
	if matchNum == 0 {
		return false
	}
	return true
}

func getLoginInfo(file string) (string, string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println("error: Requires a database configuration file!")
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var login Login
	json.Unmarshal([]byte(byteValue), &login)

	return login.DbUser, login.DbPasswd
}

func dbConn() *sql.DB {
	dbDriver := "mysql"
	dbUser, dbPasswd := getLoginInfo("conf.json")
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

func dbRetrieve(uid int64) (int, *Content) {
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
		id  string
		age int
	)
	err = rows.Scan(&age, &id)
	if err != nil {
		return http.StatusNotFound, nil
	}
	return http.StatusOK, &Content{uid, id, age}
}

func dbUpdate(dst *Content) (int, *Content) {
	db := dbConn()
	defer db.Close()

	if !uidExists(dst.Uid) {
		return http.StatusNotFound, nil
	}

	_, err := db.Exec("UPDATE bar SET id = ?, age = ? WHERE uid = ?", dst.Id, dst.Age, dst.Uid)
	if err != nil {
		panic(err.Error())
	}
	return http.StatusOK, dst
}

func dbDelete(uid int64) int {
	db := dbConn()
	defer db.Close()

	if !uidExists(uid) {
		return http.StatusNotFound
	}

	_, err := db.Exec("DELETE FROM bar WHERE uid = ?", uid)
	if err != nil {
		return http.StatusNotFound
	}
	return http.StatusOK
}
