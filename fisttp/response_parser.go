package fisttp

import (
	"encoding/json"
	"log"
	"strings"
)

func parseIndexResponse(payload []byte) *IndexResponse {
	return &IndexResponse{
		Ok: strings.TrimSpace(string(payload)) == "Text has been indexed",
	}
}

func parseSearchResponse(payload []byte) *SearchResponse {
	var documents []string
	err := json.Unmarshal(payload, &documents)
	if err != nil {
		log.Print("Got malformed response from server")
	}
	return &SearchResponse{Documents: documents}
}

func parseExitResponse(payload []byte) *ExitResponse {
	return &ExitResponse{
		Ok: strings.TrimSpace(string(payload)) == "Bye",
	}
}

func parseVersionResponse(payload []byte) *VersionResponse {
	return &VersionResponse{
		version: strings.TrimSpace(string(payload)),
	}
}

// ParseResponse will return the corresponding type of response
// by given the Verb/Command of the request and resulting body
// payload from server
func ParseResponse(verb Verb, payload []byte) Response {
	switch verb {
	case INDEX:
		return parseIndexResponse(payload)
	case SEARCH:
		return parseSearchResponse(payload)
	case EXIT:
		return parseExitResponse(payload)
	case VERSION:
		return parseVersionResponse(payload)
	}
	return nil
}
