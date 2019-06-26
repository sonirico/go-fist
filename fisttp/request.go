package fisttp

import "bytes"

type Request interface {
	Type() Verb
	String() string
}

// Exit request
type ExitRequest struct {
}

func (er *ExitRequest) String() string {
	return string(EXIT) + "\n"
}

func (er *ExitRequest) Type() Verb {
	return EXIT
}

func NewExitRequest() Request {
	return &ExitRequest{}
}

// Index request
type IndexRequest struct {
	document string
	payload  string
}

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
	out.WriteString("\n")

	return out.String()
}

func (er *IndexRequest) Type() Verb {
	return INDEX
}

// Search request
type SearchRequest struct {
	payload string
}

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
	out.WriteString("\n")

	return out.String()
}

func (er *SearchRequest) Type() Verb {
	return SEARCH
}
