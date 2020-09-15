package server

import (
	"fmt"
	mintcommon "mintserver/common"
	serverinterface "mintserver/interface"
	"net"
	"time"
)

// MintServer implements a server module
type MintServer struct {
	Name        string // Name of the server
	IPVersion   string // IP version
	BindAddress string // Bind IP
	Port        int    // Port to listen

	bufSize      int    // size of the buffer
	enableLog    bool   // whether log is enabled
	logToConsole bool   // output log to console
	logPath      string // log path
	// TODO: log into file
}

// Init the server
func (s *MintServer) Init() {
	mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
		fmt.Sprintf("Initiate server %s on %s:%d", s.Name, s.BindAddress, s.Port))
	go func() {
		/* 1. Bind TCP ip:port */
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.BindAddress, s.Port))
		if err != nil {
			if s.enableLog && s.logToConsole {
				mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
					fmt.Sprintf("ResolveTCPAddr failed: %s", err))
			}
			return
		}

		/* 2. Listen on server */
		tcplistener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			if s.enableLog && s.logToConsole {
				mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
					fmt.Sprintf("ListenTCP failed: %s", err))
			}
		}

		if s.enableLog && s.logToConsole {
			mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
				fmt.Sprintf("Start mint %s (%s) server succeeded. Listening...", s.Name, s.IPVersion))
		}

		/* 3. Wait for connections and process requests */
		for {
			conn, err := tcplistener.AcceptTCP()
			if err != nil {
				if s.enableLog && s.logToConsole {
					mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
						fmt.Sprintf("Accept connection failed: %s", err))
				}
				continue
			}

			/* Simple function, return welcome */
			go func() {
				retry := false
				retryTimes := 0
				buf := make([]byte, s.bufSize)
				for {
					/* retry 10 times with 1 second each time */
					if retryTimes == 10 {
						mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
							fmt.Sprintf("Close a connection due to timeout"))
						break
					}
					if retry {
						retryTimes++
					}
					if _, err := conn.Read(buf); err != nil {
						retry = true
						time.Sleep(time.Second * 1)
						continue
					}

					/* Respond to client */
					if _, err := conn.Write([]byte("Welcome!")); err != nil {
						if s.enableLog && s.logToConsole {
							mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
								fmt.Sprintf("Write buffer failed: %s", err))
						}
						retry = false
						continue
					}
				}
			}()
		}
	}()
}

// Stop the server
func (s *MintServer) Stop() {

}

// Run the server
func (s *MintServer) Run() {
	s.Init()
	select {}
}

// NewMintServerDefault returns a new server interface
func NewMintServerDefault(_name string) serverinterface.ServerInterface {
	if _name == "" {
		_name = "MintDefault"
	}
	s := &MintServer{Name: _name, IPVersion: "tcp4", BindAddress: "127.0.0.1", Port: 30000}
	s.bufSize = 1024
	s.enableLog = true
	s.logToConsole = true
	s.logPath = ""
	return s
}
