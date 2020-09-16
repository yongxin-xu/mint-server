package handler

import (
	"fmt"
	mintcommon "mintserver/common"
	"net"

	"github.com/golang/protobuf/proto"
)

// MainHanlder handles functions
// For example, Login function is received, call LoginCheck
// or Register function is recevied, call Register
// Now only login is implemented here
// TODO: more functions to implement
func MainHandler(conn *net.TCPConn, data []byte, cnt int) error {
	userInfo := &mintcommon.AppUser{}
	if err := proto.Unmarshal(data[:cnt], userInfo); err != nil {
		return err
	}
	result, err := LoginCheck(userInfo.Account, userInfo.Password)
	if err != nil {
		return err
	}
	if result {
		mintcommon.DebugPrint(true, true, "",
			fmt.Sprintf("[info] login succeeded, account: %s", userInfo.Account))
		if _, err := conn.Write([]byte("Login success, welcome yongxin!")); err != nil {
			return err
		}
	} else {
		mintcommon.DebugPrint(true, true, "",
			fmt.Sprintf("[info] login failed, account: %s", userInfo.Account))
		if _, err := conn.Write([]byte("Login failed, try again!")); err != nil {
			return err
		}
	}

	return nil
}
