package protocol

import (
	"fmt"
	mintcommon "mintserver/common"
	"mintserver/config"
	"net"

	"github.com/golang/protobuf/proto"
)

type functionType string

const (
	WELCOME functionType = "Welcome"
	SIGNIN functionType = "SignIn"
	SIGNUP functionType = "SignUp"
	UNKNOWN functionType = "Unknown"
)

const (
	HANDSHAKE = 0
	REQUSET_SIGNIN = 1
	REQUEST_SIGNUP = 2
	RESPONSE_SIGNIN = 3
	RESPONSE_SIGNUP = 4
)

// MainHanlder handles functions
// For example, Login function is received, call LoginCheck
// or Register function is received, call Register
// Now only login is implemented here
func MainHandler(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println(data)
	if len(data) <= 4 {
		if err := serverResponse(conn, nil, UNKNOWN, ServerReturnCode_UNKNOWN_FUNC); err != nil {
			return err
		}
		return nil
	}
	proto_len := mintcommon.BytesToUint16(data[0:2])
	fn_type := mintcommon.BytesToUint16(data[2:4])

	fmt.Println(fn_type, proto_len)

	rl := &ReqLogin{}
	rr := &ReqRegister{}

	var fn functionType

	switch fn_type {
	case HANDSHAKE:
		fn = WELCOME
	case REQUSET_SIGNIN:
		fn = SIGNIN
	case REQUEST_SIGNUP:
		fn = SIGNUP
	default:
		fn = UNKNOWN
	}

	au := &PlayerInfo{}
	switch fn {
	case WELCOME:
		if err := writeShakehandResponse(conn); err != nil {
			return err
		}
	case SIGNIN:
		if err := proto.Unmarshal(data[5:cnt], rl); err != nil {
			return err
		}
		au.Account = rl.GetAccount()
		au.Password = rl.GetPassword()
		fmt.Println(rl)
		result := signIn(au)
		err2 := serverResponse(conn, au, SIGNIN, result);
		if err2 != nil {
			return err2
		}
	case SIGNUP:
		if err := proto.Unmarshal(data[5:cnt], rr); err != nil {
			return err
		}
		au = rr.GetPlayerInfo()
		fmt.Println(rr)
		result := signUp(au)
		err2 := serverResponse(conn, au, SIGNUP, result);
		if err2 != nil {
			return err2
		}
	case UNKNOWN:
		if err := serverResponse(conn, au, fn, ServerReturnCode_UNKNOWN_FUNC); err != nil {
			return err
		}
	default:
		if err := serverResponse(conn, au, fn, ServerReturnCode_UNKNOWN_FUNC); err != nil {
			return err
		}
	}
	return nil
}

// serverResponse send messages back to clients
// The message includes
// 1. Whether the operation failed or not
// 2. Some additional information
func serverResponse(conn *net.TCPConn, au *PlayerInfo, ft functionType, result ServerReturnCode) error {
	switch ft {
	case SIGNIN:
		switch result {
		case ServerReturnCode_OK:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
				config.GlobalConfiguration.LogToConsole,
				config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in succeeded, account: %s", au.Account))
			return writeSignInResponse(conn, ServerReturnCode_OK)
		case ServerReturnCode_ACC_PSW_NO_MATCH:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in failed, account: %s", au.Account))
			return writeSignInResponse(conn, ServerReturnCode_ACC_PSW_NO_MATCH)
		case ServerReturnCode_ACC_INVALID:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in failed, account: %s", au.Account))
			return writeSignInResponse(conn, ServerReturnCode_ACC_INVALID)
		case ServerReturnCode_PSW_INVALID:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in failed, account: %s", au.Account))
			return writeSignInResponse(conn, ServerReturnCode_PSW_INVALID)
		case ServerReturnCode_DBFAIL:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign in failed, account: %s", au.Account))
			return writeSignInResponse(conn, ServerReturnCode_DBFAIL)
		default:
			break
		}
	case SIGNUP:
		switch result {
		case ServerReturnCode_OK:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up succeeded, account: %s", au.Account))
			return writeSignUpResponse(conn, ServerReturnCode_OK)
		case ServerReturnCode_ACC_INVALID:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeSignUpResponse(conn, ServerReturnCode_ACC_INVALID)
		case ServerReturnCode_PSW_INVALID:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeSignUpResponse(conn, ServerReturnCode_PSW_INVALID)
		case ServerReturnCode_ACC_EXISTED:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeSignUpResponse(conn, ServerReturnCode_PSW_INVALID)
		case ServerReturnCode_DBFAIL:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] sign up failed, account: %s", au.Account))
			return writeSignUpResponse(conn, ServerReturnCode_DBFAIL)
		default:
			break
		}
	default:
		mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
			config.GlobalConfiguration.LogToConsole,
			config.GlobalConfiguration.LogPath,
			"[info] Unknown function in ServerResponse")
		return nil
	}
	return nil
}

// writeSignInResponse is the internal implementation of serverResponse
// it sends the proto message to client use net.TCPConn.Write
func writeSignInResponse(conn *net.TCPConn, isError ServerReturnCode) error {
	rsp := &RetLogin{Code: isError}
	data, err := proto.Marshal(rsp)
	if err != nil {
		return err
	}
	resp_type := mintcommon.Uint16ToBytes(RESPONSE_SIGNIN)
	suflen := mintcommon.Uint8ToBytes(uint8(len(data)))
	buf_len := mintcommon.Uint16ToBytes(uint16(len(data) + 1))
	if _, err := conn.Write(mintcommon.BytesConcatenate(buf_len, resp_type, suflen, data)); err != nil {
		return err
	}
	return nil
}

// writeSignUpResponse is the internal implementation of serverResponse
// it sends the proto message to client use net.TCPConn.Write
func writeSignUpResponse(conn *net.TCPConn, isError ServerReturnCode) error {
	srvrsp := &RetRegister{Code: isError}
	data, err := proto.Marshal(srvrsp)
	if err != nil {
		return err
	}
	resp_type := mintcommon.Uint16ToBytes(RESPONSE_SIGNUP)
	suflen := mintcommon.Uint8ToBytes(uint8(len(data)))
	buf_len := mintcommon.Uint16ToBytes(uint16(len(data) + 1))
	if _, err := conn.Write(mintcommon.BytesConcatenate(buf_len, resp_type, suflen, data)); err != nil {
		return err
	}
	return nil
}

// writeShakehandResponse is the internal implementation of serverResponse
// it sends the proto message to client use net.TCPConn.Write
func writeShakehandResponse(conn *net.TCPConn) error {
	srvrsp := &Handshake{Token: "Welcome"}
	data, err := proto.Marshal(srvrsp)
	if err != nil {
		return err
	}
	resp_type := mintcommon.Uint16ToBytes(HANDSHAKE)
	suflen := mintcommon.Uint8ToBytes(uint8(len(data)))
	buf_len := mintcommon.Uint16ToBytes(uint16(len(data) + 1))
	if _, err := conn.Write(mintcommon.BytesConcatenate(buf_len, resp_type, suflen, data)); err != nil {
		return err
	}
	return nil
}

// WrapSignUpRequest wrap a AppUser message and a function string into ClientRequest
func WrapSignUpRequest(user *PlayerInfo) *ReqRegister {
	cr := &ReqRegister{}
	cr.PlayerInfo = user
	return cr
}

// WrapSignInRequest wrap a AppUser message and a function string into ClientRequest
func WrapSignInRequest(user *PlayerInfo) *ReqLogin {
	cr := &ReqLogin{}
	cr.Account = user.GetAccount()
	cr.Password = user.GetPassword()
	return cr
}
