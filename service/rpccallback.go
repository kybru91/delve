package service

// RPCCallback is used by RPC methods to return their result asynchronously.
type RPCCallback interface {
	Return(out interface{}, err error)

	// SetupDoneChan returns a channel that should be closed to signal that the
	// asynchronous method has completed setup and the server is ready to
	// receive other requests.
	SetupDoneChan() chan struct{}

	// DisconnectChan returns a channel that should be closed to signal that
	// the client that initially issued the command has been disconnected.
	DisconnectChan() chan struct{}
}
