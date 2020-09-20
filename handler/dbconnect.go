package handler

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	mintcommon "mintserver/common"
	"mintserver/config"
)

// Currently, we use short connection to DB
// TODO, in the future, long connection to DB may be supported

// signUpTry use a db connection to try signing up.
// return vals:
// 		1. status of the result
//		2. id of the account generated by the database, if failed, id is 0
//		3. error
func signUpTry(account string, name string, password string) (uint32, uint32, error) {
	dbconn, err := getDBConn()
	if err != nil {
		return DBFAIL, 0, err
	}
	// 1. set transaction level
	if _, err := dbconn.Query("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE"); err != nil {
		return DBFAIL, 0, err
	}

	// 2. begin transaction
	if _, err := dbconn.Query("BEGIN"); err != nil {
		return DBFAIL, 0, err
	}

	// 3. insert into ... values => insert data
	insert_stmt, err := dbconn.Prepare("INSERT INTO muser(_account, _name, _password) VALUES(?, ?, ?)")
	if err != nil {
		return DBFAIL, 0, err
	}
	_, err = insert_stmt.Exec(account, name, password)
	me, ok := err.(*mysql.MySQLError)
	if ok && me.Number == 1062 {
		return ACC_EXISTED, 0, nil
	}

	// 4. select ... => get id
	select_sql := fmt.Sprintf("SELECT _id FROM muser WHERE _account = \"%s\"",
		account)

	rets, err := dbconn.Query(select_sql)
	if err != nil {
		return DBFAIL, 0, err
	}
	var ret_id int = 0
	for rets.Next() {
		// for each row, scan the result into our tag composite object
		err = rets.Scan(&ret_id)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		break
	}

	// 5. commit()
	if _, err := dbconn.Query("COMMIT"); err != nil {
		return DBFAIL, 0, err
	}
	return OK, uint32(ret_id), nil
}

// signInTry use a db connection to try signing in
// return vals:
//		1. status of the result
//		2. error
func signInTry(account string, id uint32, password string) (uint32, error) {
	var select_sql string
	if id == 0 {
		// use account to sign in
		select_sql = fmt.Sprintf("SELECT _id FROM muser WHERE _account = '%s' and _password = '%s'",
			account, password)
	} else {
		// use id to sign in
		select_sql = fmt.Sprintf("SELECT _id FROM muser WHERE _id = %d and _password = '%s'",
			id, password)
	}
	dbconn, err := getDBConn()
	if err != nil {
		return DBFAIL, err
	}
	// 1. set transaction level
	if _, err := dbconn.Query("SET TRANSACTION ISOLATION LEVEL REPEATABLE READ"); err != nil {
		return DBFAIL, err
	}

	// 2. begin transaction
	if _, err := dbconn.Query("BEGIN"); err != nil {
		return DBFAIL, err
	}

	// 3. select ... => get id
	rets, err := dbconn.Query(select_sql)
	if err != nil {
		return DBFAIL, err
	}
	count := 0
	for rets.Next() {
		count++
		break
	}
	if count == 0 {
		return ACC_PSW_NO_MATCH, err
	}

	// 4. commit()
	if _, err := dbconn.Query("COMMIT"); err != nil {
		return DBFAIL, err
	}
	return OK, nil
}

func getDBConn() (*sql.DB, error) {
	uname := config.GlobalConfiguration.DBUser
	password := config.GlobalConfiguration.DBPassword
	ipAddress := config.GlobalConfiguration.DBHost
	port := config.GlobalConfiguration.DBPort
	dbName := config.GlobalConfiguration.DBSchemaName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?", uname, password, ipAddress, port, dbName)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
			fmt.Sprintf("mysql connect failed, detail is [%v]", err.Error()))
		return nil, err
	}
	return DB, nil
}