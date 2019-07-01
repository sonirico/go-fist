package fisttp

// Response represents the main interface exposed by all kinds of responses
type Response interface {
	IsOk() bool       // Whether the operation completed successfully
	responseElement() // Hacky custom method besides `IsOk` so as to have inheritors to adhere to something
}

// IndexResponse holds the feedback from server after an INDEX request. It only stores
// a binary flag, that is whether the operation completed victoriously
type IndexResponse struct {
	Ok bool
}

func (ir *IndexResponse) responseElement() {}

// IsOk returns whether the action in the server can be considered to have completed ok
func (ir *IndexResponse) IsOk() bool {
	return ir.Ok
}

// SearchResponse hols the matching documents the server has found occurrences
type SearchResponse struct {
	Documents []string
}

func (sr *SearchResponse) responseElement() {}

// IsOk returns whether the server has found matches in documents for the given search payload
func (sr *SearchResponse) IsOk() bool {
	// Empty responses (meaning no matches) are not errors...
	return len(sr.Documents) > 0
}

// ExitResponse carries whether the server has terminated in a way we can say everything went
// as expected
type ExitResponse struct {
	Ok bool
}

func (er *ExitResponse) responseElement() {}

// IsOk returns whether the action in the server can be considered to have completed ok
func (er *ExitResponse) IsOk() bool {
	return er.Ok
}
