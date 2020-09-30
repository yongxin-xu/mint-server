package server

import (
	"fmt"
	mintcommon "mintserver/common"
	"mintserver/config"
	mintinterfaces "mintserver/interface"
	"net"
)

// Connector module
type Connector struct {
	Conn        *net.TCPConn               // TCP connection
	ConnID      uint32                     // Connect ID
	UserID      int						   // User ID
	isClosed    bool                       // Whether the connection is closed
	funcHanlder mintinterfaces.MintHandler // function handler
	ExitChan    chan bool                  // chan has exited
}

// NewConnector receives a TCP connection and its callback function
// and returns a new connector
func NewConnector(conn *net.TCPConn, connID uint32, userid int, callback mintinterfaces.MintHandler) *Connector {
	c := &Connector{Conn: conn, ConnID: connID, UserID: userid, funcHanlder: callback,
		isClosed: false, ExitChan: make(chan bool, 1)}
	return c
}

// Connection start to work
func (c *Connector) Start() {
	mintcommon.DebugPrint(true, true, "",
		fmt.Sprintf("[info] start connection, id: %d", c.ConnID))
	go func() {
		defer mintcommon.DebugPrint(true, true, "",
			fmt.Sprintf("[info] connection exit, id: %s", c.Conn.RemoteAddr()))
		defer c.Stop()

		for {
			buf := make([]byte, config.GlobalConfiguration.MaxPackageSize)
			cnt, err := c.Conn.Read(buf)
			if err != nil {
				break
			}
			/* call handler */
			if err := c.funcHanlder(c.Conn, &c.UserID, buf, cnt); err != nil {
				mintcommon.DebugPrint(true, true, "",
					fmt.Sprintf("[error] handler error: %s", err))
			}
			if c.isClosed {
				break
			}
		}
	}()
}

// Stop the connection
func (c *Connector) Stop() {
	mintcommon.DebugPrint(true, true, "",
		fmt.Sprintf("[info] stop connection, id: %d", c.ConnID))

	if c.isClosed {
		mintcommon.DebugPrint(true, true, "",
			fmt.Sprintf("[warning] connection already closed, id: %d", c.ConnID))
	} else {
		if err := c.Conn.Close(); err != nil {
			mintcommon.DebugPrint(true, true, "",
				fmt.Sprintf("[warning] close connection failed , id: %d", c.ConnID))
		}
		c.isClosed = true
		close(c.ExitChan)
	}
}

// Get TCP connection
func (c *Connector) GetClientConnection() *net.TCPConn {
	return c.Conn
}

// Get connection ID
func (c *Connector) GetClientConnID() uint32 {
	return c.ConnID
}

// Get IP:Port
func (c *Connector) GetClientAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send data to client
func (c *Connector) Send(data []byte, cnt int) error {
	return nil
}
