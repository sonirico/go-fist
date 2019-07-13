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

// VersionResponse carries the version issued by the server
type VersionResponse struct {
	version string
}

// IsOk returns whether a valid version string has been obtained
func (vr *VersionResponse) IsOk() bool {
	return len(vr.version) > 0
}

// GetVersion returns the version from the server response
func (vr *VersionResponse) GetVersion() string {
	return vr.version
}

func (vr *VersionResponse) responseElement() {}

// DeleteResponse holds whether keyword deletion has been successful
type DeleteResponse struct {
	ok bool
}

// IsOk returns whether deletion has been rightly performed
func (dr *DeleteResponse) IsOk() bool {
	return dr.ok
}

func (dr *DeleteResponse) responseElement() {}
