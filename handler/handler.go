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
	SETPROGRESS functionType = "SetProgress"
	GETPROGRESS functionType = "GetProgress"
)

const (
	HANDSHAKE = 0
	REQUSET_SIGNIN = 1
	REQUEST_SIGNUP = 2
	RESPONSE_SIGNIN = 3
	RESPONSE_SIGNUP = 4
	REQUEST_SET_PROGR = 5
	RESPONSE_PROGRESS = 6
	REQUEST_GET_PROGR = 7
)

// MainHanlder handles functions
// For example, Login function is received, call LoginCheck
// or Register function is received, call Register
// Now only login is implemented here
func MainHandler(conn *net.TCPConn, CID *int, data []byte, cnt int) error {
	for cnt > 0 {
		fmt.Println(data)
		if len(data) <= 4 {
			if err := serverResponse(conn, nil, UNKNOWN, ServerReturnCode_UNKNOWN_FUNC); err != nil {
				return err
			}
			return nil
		}
		proto_len := mintcommon.BytesToUint16(data[0:2])
		fn_type := mintcommon.BytesToUint16(data[2:4])
		suflen := mintcommon.BytesToUint8(data[4:5])

		cnt = cnt - 5 - int(suflen)

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
		case REQUEST_SET_PROGR:
			fn = SETPROGRESS
		case REQUEST_GET_PROGR:
			fn = GETPROGRESS
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
			if err := proto.Unmarshal(data[5:5+suflen], rl); err != nil {
				return err
			}
			au.Account = rl.GetAccount()
			au.Password = rl.GetPassword()
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
				config.GlobalConfiguration.LogToConsole,
				config.GlobalConfiguration.LogPath,
				rl.String())
			__id, result := signIn(au)
			*CID = __id
			if err2 := serverResponse(conn, au, SIGNIN, result); err2 != nil {
				return err2
			}
		case SIGNUP:
			if err := proto.Unmarshal(data[5:5+suflen], rr); err != nil {
				return err
			}
			au = rr.GetPlayerInfo()
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
				config.GlobalConfiguration.LogToConsole,
				config.GlobalConfiguration.LogPath,
				rr.String())
			result := signUp(au)
			err2 := serverResponse(conn, au, SIGNUP, result);
			if err2 != nil {
				return err2
			}
		case GETPROGRESS:
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
				config.GlobalConfiguration.LogToConsole,
				config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] user get progress, uid %d", *CID))
			rp, err := getProgress(*CID)
			if err != nil && rp == nil {
				rp = &RetProgress{Chapter: -1, Section: -1}
			}
			if err3 := writeProgressResponse(conn, *CID, rp, false); err3 != nil {
				return err3
			}
		case SETPROGRESS:
			rp := &ReqProgress{}
			if err := proto.Unmarshal(data[5:5+suflen], rp); err != nil {
				return err
			}
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
				config.GlobalConfiguration.LogToConsole,
				config.GlobalConfiguration.LogPath,
				rr.String())
			rspp, err := setProgress(*CID, rp)
			if err != nil {
				return err
			}
			if err2 := writeProgressResponse(conn, *CID, rspp, true); err2 != nil {
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
		data = data[5+suflen:]
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
			return writeSignUpResponse(conn, ServerReturnCode_ACC_EXISTED)
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

// writeProgressResponse is the internal implementation of serverResponse
// it sends the proto message to client use net.TCPConn.Write
func writeProgressResponse(conn *net.TCPConn, id int, rsp *RetProgress, output bool) error {
	if output == true {
		if rsp.GetChapter() != -1 && rsp.GetSection() != -1 {
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
				config.GlobalConfiguration.LogToConsole,
				config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[info] set progress for, id: %d, chapter: %d, section: %d succeeded",
					id, rsp.GetChapter(), rsp.GetSection()))
		} else {
			mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
				config.GlobalConfiguration.LogToConsole,
				config.GlobalConfiguration.LogPath,
				fmt.Sprintf("[error] set progress for, id: %d failed",
					id, rsp.GetChapter(), rsp.GetSection()))
		}
	}
	mintcommon.DebugPrint(config.GlobalConfiguration.EnableLog,
		config.GlobalConfiguration.LogToConsole,
		config.GlobalConfiguration.LogPath,
		fmt.Sprintf("[info] return progress to, id: %d, chapter: %d, section: %d",
		id, rsp.GetChapter(), rsp.GetSection()))
	data, err := proto.Marshal(rsp)
	if err != nil {
		return err
	}
	resp_type := mintcommon.Uint16ToBytes(RESPONSE_PROGRESS)
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

// WrapSignUpRequest wrap a PlayerInfo message and a function string into ClientRequest
func WrapSignUpRequest(user *PlayerInfo) *ReqRegister {
	cr := &ReqRegister{}
	cr.PlayerInfo = user
	return cr
}

// WrapSignInRequest wrap a PlayerInfo message and a function string into ClientRequest
func WrapSignInRequest(user *PlayerInfo) *ReqLogin {
	cr := &ReqLogin{}
	cr.Account = user.GetAccount()
	cr.Password = user.GetPassword()
	return cr
}

// WrapSetProgressRequest wrap a PlayerInfo message and a function string into ClientRequest
func WrapSetProgressRequest(chap int, sec int) *RetProgress {
	cr := &RetProgress{}
	cr.Chapter = int32(chap)
	cr.Section = int32(sec)
	return cr
}