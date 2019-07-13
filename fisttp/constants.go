package fisttp

// Verb or command issued by the client
type Verb string

// This is a list of currently implemented commands
const (
	INDEX   Verb = "INDEX"   // Request the server to index a document
	SEARCH       = "SEARCH"  // Query the server to find a matching document for a custom payload
	EXIT         = "EXIT"    // Request to server to terminate the session gracefully
	VERSION      = "VERSION" // Request the server to get the server version
	DELETE		 = "DELETE"  // Request the server to remove a set of keywords from all documents
)

// REOL marks the end of a request
const REOL = "\r\n"