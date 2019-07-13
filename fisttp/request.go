package fisttp

import "bytes"

// Request provides an interface that all types of request should match
type Request interface {
	// Get the type of the request, which is determined by the command/verb
	Type() Verb
	// Get the string that will be sent to the server over the wire
	String() string
}

// ExitRequest represents an attempt to terminate a connection
type ExitRequest struct{}

func (er *ExitRequest) String() string {
	return string(EXIT) + REOL
}

// Type gets the type of the request. EXIT will be issued
func (er *ExitRequest) Type() Verb {
	return EXIT
}

// NewExitRequest returns a pointer to a newly allocated ExitRequest. No context required
func NewExitRequest() Request {
	return &ExitRequest{}
}

// IndexRequest represents the action of indexing a document
type IndexRequest struct {
	document string // The target document
	payload  string // The actual content to be indexed
}

// NewIndexRequest returns a pointer to a newly allocated IndexRequest struct, by receiving
// the target document and content to be indexed
func NewIndexRequest(doc string, payload string) *IndexRequest {
	return &IndexRequest{
		document: doc,
		payload:  payload,
	}
}

func (er *IndexRequest) String() string {
	var out bytes.Buffer

	out.WriteString(string(er.Type()))
	out.WriteString(" ")
	out.WriteString(er.document)
	out.WriteString(" ")
	out.WriteString(er.payload)
	out.WriteString(REOL)

	return out.String()
}

// Type gets the type of the request. INDEX will be issued
func (er *IndexRequest) Type() Verb {
	return INDEX
}

// SearchRequest represents a query to the server in order to find the
// matching documents for a given payload
type SearchRequest struct {
	payload string
}

// NewSearchRequest returns a newly allocated SearchRequest by given the
// content whose documents are to be found as context
func NewSearchRequest(payload string) *SearchRequest {
	return &SearchRequest{
		payload: payload,
	}
}

func (er *SearchRequest) String() string {
	var out bytes.Buffer

	out.WriteString(string(er.Type()))
	out.WriteString(" ")
	out.WriteString(er.payload)
	out.WriteString(REOL)

	return out.String()
}

// Type gets the type of the request. SEARCH will be issued
func (er *SearchRequest) Type() Verb {
	return SEARCH
}

// DeleteRequest requests the action of keyword removal. Will
// apply to all documents
type DeleteRequest struct {
	keywords string
}

// NewDeleteRequest returns a newly allocated DeleteRequest which
// will carry the keywords to be deleted
func NewDeleteRequest(payload string) *DeleteRequest {
	return &DeleteRequest{keywords: payload}
}

// Type gets the type of the request. DELETE will be issued
func (dr *DeleteRequest) Type() Verb {
	return DELETE
}

func (dr *DeleteRequest) String() string {
	var out bytes.Buffer

	out.WriteString(string(dr.Type()))
	out.WriteString(" ")
	out.WriteString(dr.keywords)
	out.WriteString(REOL)

	return out.String()
}

// VersionRequest represents a query to the server in order to
// find out the version of it
type VersionRequest struct{}

// NewVersionRequest returns a newly allocated VersionRequest
func NewVersionRequest() *VersionRequest {
	return &VersionRequest{}
}

// Type returns the type of the request. VERSION will be issued
func (vr *VersionRequest) Type() Verb {
	return VERSION
}

func (vr *VersionRequest) String() string {
	return string(VERSION) + REOL
}
