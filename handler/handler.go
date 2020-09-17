package handler

import (
	"fmt"
	mintcommon "mintserver/common"
	"net"

	"github.com/golang/protobuf/proto"
)

// MainHanlder handles functions
// For example, Login function is received, call LoginCheck
// or Register function is received, call Register
// Now only login is implemented here
// TODO: more functions to implement and global parameters
func MainHandler(conn *net.TCPConn, data []byte, cnt int) error {
	cr := &ClientRequest{}
	if err := proto.Unmarshal(data[:cnt], cr); err != nil {
		return err
	}
	au := cr.GetAu()
	switch cr.GetFuction() {
	case "SignIn":
		result, err := SignIn(au.Account, au.Password)
		if err != nil {
			return err
		}
		if result {
			mintcommon.DebugPrint(true, true, "",
				fmt.Sprintf("[info] sign in succeeded, account: %s", au.Account))
			if _, err := conn.Write([]byte("sign in success, welcome yongxin!")); err != nil {
				return err
			}
		} else {
			mintcommon.DebugPrint(true, true, "",
				fmt.Sprintf("[info] signin failed, account: %s", au.Account))
			if _, err := conn.Write([]byte("signin failed, try again!")); err != nil {
				return err
			}
		}
	case "SignUp":
		result, err := SignUp(au.Account, au.Name, au.Password)
		if err != nil {
			return err
		}
		if result != -1 {
			mintcommon.DebugPrint(true, true, "",
				fmt.Sprintf("[info] sign up succeeded, account: %s", au.Account))
			if _, err := conn.Write([]byte("Sign up success, welcome yongxin!")); err != nil {
				return err
			}
		} else {
			mintcommon.DebugPrint(true, true, "",
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			if _, err := conn.Write([]byte("Sign up failed, try again!")); err != nil {
				return err
			}
		}
	default:
		if _, err := conn.Write([]byte("Unknown Function!")); err != nil {
			return err
		}
	}

	return nil
}

// WrapAccountFunction wrap a AppUser message and a function string into ClientRequest
// this involves "SignIn" and "SignUp" function
func WrapAccountRequest(user *AppUser, fname string) *ClientRequest {
	cr := &ClientRequest{}
	cr.Fuction = fname
	cr.Au = user
	return cr
}
