package handler

import (
	"fmt"
	mintcommon "mintserver/common"
	"net"

	"github.com/golang/protobuf/proto"
)

type functionType string

const (
	SIGNIN functionType = "SignIn"
	SIGNUP functionType = "SignUp"
)

const (
	SUCCESS = 0
	ERROR = 1
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
	fn := cr.GetFuction()
	switch functionType(fn) {
	case SIGNIN:
		result, err := SignIn(au.Account, au.Password)
		err2 := serverResponse(conn, au, SIGNIN, result);
		if err != nil || err2 != nil {
			return err
		}
	case SIGNUP:
		result, err := SignUp(au.Account, au.Name, au.Password)
		err2 := serverResponse(conn, au, SIGNUP, result);
		if err != nil || err2 != nil {
			return err
		}
	default:
		if err := serverResponse(conn, au, functionType(fn), false); err != nil {
			return err
		}
	}
	return nil
}

// serverResponse send messages back to clients
// The message includes
// 1. Whether the operation failed or not
// 2. Some additional information
func serverResponse(conn *net.TCPConn, au *AppUser, ft functionType, result bool) error {
	switch ft {
	case SIGNIN:
		if result {
			mintcommon.DebugPrint(true, true, "",
				fmt.Sprintf("[info] sign in succeeded, account: %s", au.Account))
			return writeResponse(conn, au, SUCCESS,
				fmt.Sprintf("Sign in success, welcome %s!", au.Name))
		} else {
			mintcommon.DebugPrint(true, true, "",
				fmt.Sprintf("[info] sign in failed, account: %s", au.Account))
			return writeResponse(conn, au, ERROR,
				fmt.Sprintf("Sign in failed, try again!"))
		}
	case SIGNUP:
		if result {
			mintcommon.DebugPrint(true, true, "",
				fmt.Sprintf("[info] sign up succeeded, account: %s", au.Account))
			return writeResponse(conn, au, SUCCESS,
				fmt.Sprintf("Sign up success, welcome %s!", au.Name))
		} else {
			mintcommon.DebugPrint(true, true, "",
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeResponse(conn, au, ERROR,
				fmt.Sprintf("Sign up failed, try again!"))
		}
	default:
		mintcommon.DebugPrint(true, true, "",
			"[info] Unknown function in ServerResponse")
		return writeResponse(conn, au, ERROR, "Unknown function, contact developer!")
	}
}

// writeResponse is the internal implementation of serverResponse
// it sends the proto message to client use net.TCPConn.Write
func writeResponse(conn *net.TCPConn, au *AppUser, isError uint32, info string) error {
	srvrsp := &SrvResponse{Error: isError, Info: info, Au: au}
	data, err := proto.Marshal(srvrsp)
	if err != nil {
		return err
	}
	if _, err := conn.Write(data); err != nil {
		return err
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
