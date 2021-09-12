package mintinterfaces

import "net"

// Interface Connector
type ConnectorInterface interface {
	// Connection start to work
	Start()

	// Stop the connection
	Stop()

	// Get TCP connection
	GetClientConnection() *net.TCPConn

	// Get connection ID
	GetClientConnID() uint32

	// Get IP:Port
	GetClientAddr() net.Addr

	// Send data to client
	Send(data []byte, cnt int) error
}

type MintHandler func(*net.TCPConn, *int, chan []byte, []byte, int) error
