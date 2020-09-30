package server

import (
	"fmt"
	mintcommon "mintserver/common"
	"mintserver/config"
	"mintserver/handler"
	mintinterfaces "mintserver/interface"
	"net"
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
		fmt.Sprintf("[info] initiate server %s on %s:%d", s.Name, s.BindAddress, s.Port))
	go func() {
		/* 1. Bind TCP ip:port */
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.BindAddress, s.Port))
		if err != nil {
			if s.enableLog && s.logToConsole {
				mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
					fmt.Sprintf("[error] resolveTCPAddr failed: %s", err))
			}
			return
		}

		/* 2. Listen on server */
		tcplistener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			if s.enableLog && s.logToConsole {
				mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
					fmt.Sprintf("[error] listenTCP failed: %s", err))
			}
		}

		if s.enableLog && s.logToConsole {
			mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
				fmt.Sprintf("[info] start mint %s (%s) server succeeded. Listening...", s.Name, s.IPVersion))
		}

		var connID uint32 = 1

		/* 3. Wait for connections and process requests */
		for {
			conn, err := tcplistener.AcceptTCP()
			if err != nil {
				if s.enableLog && s.logToConsole {
					mintcommon.DebugPrint(s.enableLog, s.logToConsole, s.logPath,
						fmt.Sprintf("[error] AcceptTCP connection failed: %s", err))
				}
				continue
			}

			connHandler := NewConnector(conn, connID, -1, protocol.MainHandler)
			connID++
			go connHandler.Start()
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
func NewMintServerDefault(_name string) mintinterfaces.ServerInterface {
	s := &MintServer{Name: config.GlobalConfiguration.ServerName,
		IPVersion:   "tcp4",
		BindAddress: config.GlobalConfiguration.ServerHost,
		Port:        config.GlobalConfiguration.ServerPort}
	s.bufSize = config.GlobalConfiguration.MaxPackageSize
	s.enableLog = config.GlobalConfiguration.EnableLog
	s.logToConsole = config.GlobalConfiguration.LogToConsole
	s.logPath = config.GlobalConfiguration.LogPath
	return s
}
