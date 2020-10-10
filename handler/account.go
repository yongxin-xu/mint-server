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
func NewUser(_account string, _psw string) *PlayerInfo {
	au := &PlayerInfo{Account: _account, Password: _psw}
	return au
}

// signIn checks the (Account, Password) pair
// Results of SignIn
//  0. user_id only valid if OK
//	1. ServerReturnCode
//		a. Account and Password not matched, or may not exist. (ACC_PSW_NO_MATCH)
//		b. OK
//		c. DBFAIL
//	2. progress of user only valid if OK
func signIn(_au *PlayerInfo) (ServerReturnCode, int) {
	defer func(){_au.Password = ""}() // mask password when finished
	// check account
	if len(_au.Account) < 8 {
		return ServerReturnCode_ACC_TOO_SHORT, 0
	} else if len(_au.Account) > 25 {
		return ServerReturnCode_ACC_TOO_LONG, 0
	} else if !isAlphaNum(_au.Account) {
		return ServerReturnCode_ACC_INVALID, 0
	}

	// check password
	if len(_au.Password) < 8 {
		return ServerReturnCode_PSW_TOO_SHORT, 0
	} else if len(_au.Password) > 25 {
		return ServerReturnCode_ACC_TOO_LONG, 0
	} else if !isAlphaNum(_au.Password) {
		return ServerReturnCode_PSW_INVALID, 0
	}

	result, id, err := signInTry(_au.Account, _au.Password)
	if err != nil {
		mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
			fmt.Sprintf("[info] Database failed %s", err))
	}
	return result, id
}

// signUp an account
// 1. Check if (Account, Name, Password) are valid
// 2. Check if Account duplicated in database
// 3. Register the user, and fetch user id
// 4. Return Result
// Results of SignUp:
// 1. Return Code
// 		a. Account not valid (ACC_INVALID)
//		b. Account already existed (ACC_EXISTED)
//		c. Password not valid (PSW_INVALID)
//		d. Name not valid (NAME_INVALID)
//		e. OK
//		f. DBFAIL
// 2. UserID
func signUp(_au *PlayerInfo) (ServerReturnCode, int) {
	defer func(){_au.Password = ""}() // mask password when finished

	// check account
	if len(_au.Account) < 8 {
		return ServerReturnCode_ACC_TOO_SHORT, 0
	} else if len(_au.Account) > 25 {
		return ServerReturnCode_ACC_TOO_LONG, 0
	} else if !isAlphaNum(_au.Account) {
		return ServerReturnCode_ACC_INVALID, 0
	}

	// check password
	if len(_au.Password) < 8 {
		return ServerReturnCode_PSW_TOO_SHORT, 0
	} else if len(_au.Password) > 25 {
		return ServerReturnCode_ACC_TOO_LONG, 0
	} else if !isAlphaNum(_au.Password) {
		return ServerReturnCode_PSW_INVALID, 0
	}

	result, _id, err := signUpTry(_au.Account, _au.Password)
	if err != nil {
		mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
			fmt.Sprintf("[info] Database failed %s", err))
	}
	return result, _id
}

func isAlphaNum(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}