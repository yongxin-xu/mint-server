package handler

import (
	"fmt"
	mintcommon "mintserver/common"
	"mintserver/config"
	"net"

	"github.com/golang/protobuf/proto"
)

type functionType string

const (
	SIGNIN functionType = "SignIn"
	SIGNUP functionType = "SignUp"
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
		result := signIn(au)
		err2 := serverResponse(conn, au, SIGNIN, result);
		if err2 != nil {
			return err2
		}
	case SIGNUP:
		result := signUp(au)
		err2 := serverResponse(conn, au, SIGNUP, result);
		if err2 != nil {
			return err2
		}
	default:
		if err := serverResponse(conn, au, functionType(fn), UNKNOWN_FUNC); err != nil {
			return err
		}
	}
	return nil
}

// serverResponse send messages back to clients
// The message includes
// 1. Whether the operation failed or not
// 2. Some additional information
func serverResponse(conn *net.TCPConn, au *AppUser, ft functionType, result uint32) error {
	switch ft {
	case SIGNIN:
		switch result {
		case OK:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
				config.GlobalConfiguration.LogToConsole,
				config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in succeeded, account: %s", au.Account))
			return writeResponse(conn, au, OK,
				fmt.Sprintf("Sign in success, welcome %s!", au.Name))
		case ACC_PSW_NO_MATCH:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in failed, account: %s", au.Account))
			return writeResponse(conn, au, ACC_PSW_NO_MATCH,
				fmt.Sprintf("Account and Password not matched, or may not exist!"))
		case ACC_INVALID:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in failed, account: %s", au.Account))
			return writeResponse(conn, au, ACC_INVALID,
				fmt.Sprintf("Account not valid, only alphabet and number are allowed, " +
					"should be less than 25 characters"))
		case PSW_INVALID:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in failed, account: %s", au.Account))
			return writeResponse(conn, au, PSW_INVALID,
				fmt.Sprintf("Password not valid, only alphabet and number are allowed, " +
					"should be less than 25 characters."))
		case DBFAIL:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in failed, account: %s", au.Account))
			return writeResponse(conn, au, DBFAIL,
				fmt.Sprintf("Server error, contact admin!"))
		default:
			break
		}
	case SIGNUP:
		switch result {
		case OK:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up succeeded, account: %s", au.Account))
			return writeResponse(conn, au, OK,
				fmt.Sprintf("Sign up success, welcome %s!", au.Name))
		case ACC_INVALID:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeResponse(conn, au, ACC_INVALID,
				fmt.Sprintf("Account not valid, only alphabet and number are allowed, " +
					"should be less than 25 characters"))
		case PSW_INVALID:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeResponse(conn, au, PSW_INVALID,
				fmt.Sprintf("Password not valid, only alphabet and number are allowed, " +
					"should be less than 25 characters."))
		case NAME_INVALID:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeResponse(conn, au, NAME_INVALID,
				fmt.Sprintf("Name not valid, should be less than 25 characters"))
		case ACC_EXISTED:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeResponse(conn, au, PSW_INVALID,
				fmt.Sprintf("Account already existed"))
		case DBFAIL:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeResponse(conn, au, DBFAIL,
				fmt.Sprintf("Server error, contact admin!"))
		default:
			break
		}
	default:
		mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
			"[info] Unknown function in ServerResponse")
		return writeResponse(conn, au, UNKNOWN_FUNC, "Unknown function, contact developer!")
	}
	return nil
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
// this involves "signIn" and "signUp" function
func WrapAccountRequest(user *AppUser, fname string) *ClientRequest {
	cr := &ClientRequest{}
	cr.Fuction = fname
	cr.Au = user
	return cr
}
