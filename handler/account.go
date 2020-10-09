package protocol

import (
	"fmt"
	mintcommon "mintserver/common"
	"mintserver/config"
	"unicode"
)

// NewUser return the pointer of a new player
// The function is for user to call
// This shall never be called by the server
func NewUser(_account string, _name string, _psw string) *PlayerInfo {
	au := &PlayerInfo{Account: _account, Password: _psw}
	return au
}

// signIn checks the (Account, Password) pair
// Results of SignIn
//  0. user_id only valid if OK
//	1. progress of user only valid if OK
//	2. ServerReturnCode
//		a. Account and Password not matched, or may not exist. (ACC_PSW_NO_MATCH)
//		b. OK
//		c. DBFAIL
func signIn(_au *PlayerInfo) (int, ServerReturnCode) {
	defer func(){_au.Password = ""}() // mask password when finished
	if len(_au.Account) == 0 || len(_au.Account) > 25 || !isAlphaNum(_au.Account) {
		return 0, ServerReturnCode_ACC_INVALID
	}
	if len(_au.Password) == 0 || len(_au.Password) > 25 || !isAlphaNum(_au.Password) {
		return 0, ServerReturnCode_PSW_INVALID
	}
	result, id, err := signInTry(_au.Account, _au.Password)
	if err != nil {
		mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
			fmt.Sprintf("[info] Database failed %s", err))
	}
	return id, result
}

// signUp an account
// 1. Check if (Account, Name, Password) are valid
// 2. Check if Account duplicated in database
// 3. Register the user, and fetch user id
// 4. Return Result
// Results of SignUp:
// 		a. Account not valid (ACC_INVALID)
//		b. Account already existed (ACC_EXISTED)
//		c. Password not valid (PSW_INVALID)
//		d. Name not valid (NAME_INVALID)
//		e. OK
//		f. DBFAIL
func signUp(_au *PlayerInfo) (ServerReturnCode) {
	defer func(){_au.Password = ""}() // mask password when finished
	if len(_au.Account) == 0 || len(_au.Account) > 25 || !isAlphaNum(_au.Account) {
		return ServerReturnCode_ACC_INVALID
	}
	if len(_au.Password) == 0 || len(_au.Password) > 25 || !isAlphaNum(_au.Password) {
		return ServerReturnCode_PSW_INVALID
	}
	result, err := signUpTry(_au.Account, _au.Password)
	if err != nil {
		mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
			fmt.Sprintf("[info] Database failed %s", err))
	}
	return result
}

func isAlphaNum(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}