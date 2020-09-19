package handler

// NewUser return the pointer of a new user
// TODO: implement ID generation based on database
func NewUser(_account string, _name string, _psw string) *AppUser {
	au := &AppUser{Account: _account, Name: _name, Password: _psw}
	return au
}

// SignIn checks the (Account, Password) pair
// TODO: use database to check
func SignIn(_account string, _psw string) (bool, error) {
	if _account == "yx123" && _psw == "yongxin123" {
		return true, nil
	}
	return false, nil
}

// SignUp an account
// 1. Check if (Account, Name, Password) are valid
// 2. Check if Account duplicated in database
// 3. Register the user, and fetch user id
// 4. Return Result
// TODO: use database to check
func SignUp(_account string, _name string, _psw string) (bool, error) {
	if _account == "yx123" && _name == "yongxin" && _psw == "yongxin123" {
		return true, nil
	}
	return false, nil
}