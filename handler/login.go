package handler

import mintcommon "mintserver/common"

// NewUser return the pointer of a new user
// TODO: implement ID generation based on database
func NewUser(_account string, _name string, _psw string) *mintcommon.AppUser {
	au := &mintcommon.AppUser{Account: _account, Name: _name, Password: _psw}
	return au
}

// LoginCheck checks the (Account, Password) pair
// TODO: use database to check
func LoginCheck(_account string, _psw string) (bool, error) {
	if _account == "yx123" && _psw == "yongxin123" {
		return true, nil
	}
	return false, nil
}
