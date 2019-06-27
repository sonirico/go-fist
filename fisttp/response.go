package fisttp

type Response interface {
	IsOk() bool
	responseElement()
}

type IndexResponse struct {
	Ok bool
}

func (ir *IndexResponse) responseElement() {}
func (ir *IndexResponse) IsOk() bool {
	return ir.Ok
}

type SearchResponse struct {
	Documents []string
}

func (sr *SearchResponse) responseElement() {}
func (sr *SearchResponse) IsOk() bool {
	// Empty responses (meaning no matches) are not errors...
	return len(sr.Documents) > 0
}

type ExitResponse struct {
	Ok bool
}

func (er *ExitResponse) responseElement() {}
func (er *ExitResponse) IsOk() bool {
	return er.Ok
}
