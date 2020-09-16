package mintinterfaces

// ServerInterface is the interface of the server
type ServerInterface interface {
	// Init the server
	Init()

	// Stop the server
	Stop()

	// Server run
	Run()
}
