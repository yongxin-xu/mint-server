package main

import "mint-server/server"

func main() {
	// 1. init a server handler based on mint server
	s := server.NewMintServerDefault("[Mint Test]")
	// 2. run the server
	s.Run()
}