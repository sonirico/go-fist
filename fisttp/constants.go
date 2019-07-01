package fisttp

// Verb or command issued by the client
type Verb string

// This is a list of currently implemented commands
const (
	INDEX  Verb = "INDEX"  // Request the server to index a document
	SEARCH      = "SEARCH" // Query the server to find a matching document for a custom payload
	EXIT        = "EXIT"   // Request to server to terminate the session gracefully
)
