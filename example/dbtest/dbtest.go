package dbtest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestConnMySQL(t *testing.T) {
	uname := "u1"
	password := "123456"
	ipAddress := "127.0.0.1"
	port := 20000
	dbName := "test"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?", uname, password, ipAddress, port, dbName)
	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	} else {
		fmt.Println("Successful")
	}
	results, err1 := Db.Query("SELECT  VERSION();")
	if err1 != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var ver string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&ver)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		t.Log(ver)
	}
}

func TestCreateTable(t *testing.T) {
	uname := "u1"
	password := "123456"
	ipAddress := "127.0.0.1"
	port := 20000
	dbName := "minigame"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?", uname, password, ipAddress, port, dbName)
	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	} else {
		fmt.Println("Successful")
	}
	if Db != nil {
		_, err := Db.Query("CREATE TABLE muser (_id INT AUTO_INCREMENT, " +
			"_account VARCHAR(30), _name VARCHAR(30), _password VARCHAR(30), PRIMARY KEY(_id), UNIQUE KEY(_account)")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestVerDB(t *testing.T) {
	db, err := sql.Open("mysql", "u1:123456@tcp(127.0.0.1:20000)")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err1 := db.Query("SELECT  VERSION();")
	if err1 != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var ver string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&ver)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		t.Log(ver)
	}
}
