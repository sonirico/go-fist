package fisttp

import (
	"testing"
)

type ResponseParserTest struct {
	input    string
	expected Response
}

func itemInList(payload []string, target string) bool {
	for _, item := range payload {
		if target == item {
			return true
		}
	}

	return false
}

func assertIndexResponseEqual(t *testing.T, expected *IndexResponse, actual Response) bool {
	if actualResponseType, ok := actual.(*IndexResponse); !ok {
		t.Errorf("wrong response type. want '%T', have '%T'",
			expected, actualResponseType)
		return false
	}
	return true
}

func assertSearchResponseEqual(t *testing.T, expected *SearchResponse, actual Response) bool {
	searchResponse, ok := actual.(*SearchResponse)
	if !ok {
		t.Errorf("wrong response type. want '%T', have '%T'",
			expected, searchResponse)
		return false
	}
	if len(expected.Documents) != len(searchResponse.Documents) {
		t.Errorf("wrong number of matching documents. want '%d', have '%d'",
			len(expected.Documents), len(searchResponse.Documents))
		return false
	}
	for _, actualDocument := range searchResponse.Documents {
		if !itemInList(expected.Documents, actualDocument) {
			t.Errorf("item %s not in list", actualDocument)
			return false
		}
	}
	return true
}

func assertVersionResponseEqual(t *testing.T, expected *VersionResponse, actual Response) bool {
	actualResponse, ok := actual.(*VersionResponse)
	if !ok {
		t.Errorf("wrong response type. want '%T', have '%T'",
			expected, actualResponse)
		return false
	}
	if expected.GetVersion() != actualResponse.GetVersion() {
		t.Errorf("wrong version. want '%s', have '%s'",
			expected.GetVersion(), actualResponse.GetVersion())
		return false
	}
	return true
}

func assertExitResponseEqual(t *testing.T, expected *ExitResponse, actual Response) bool {
	if actualResponseType, ok := actual.(*ExitResponse); !ok {
		t.Errorf("wrong response type. want '%T', have '%T'",
			expected, actualResponseType)
		return false
	}
	return true
}

func assertDeleteResponseEqual(t *testing.T, expected *DeleteResponse, actual Response) bool {
	if actualResponseType, ok := actual.(*DeleteResponse); !ok {
		t.Errorf("wrong response type. want '%T', have '%T'",
			expected, actualResponseType)
		return false
	}
	return true
}

func assertResponseEqual(t *testing.T, expected Response, actual Response) bool {
	t.Helper()

	if expected.IsOk() != actual.IsOk() {
		t.Errorf("expected Response to be ok. want 'OK', have 'KO'")
		return false
	}

	switch expectedResponse := expected.(type) {
	case *IndexResponse:
		return assertIndexResponseEqual(t, expectedResponse, actual)
	case *SearchResponse:
		return assertSearchResponseEqual(t, expectedResponse, actual)
	case *ExitResponse:
		return assertExitResponseEqual(t, expectedResponse, actual)
	case *VersionResponse:
		return assertVersionResponseEqual(t, expectedResponse, actual)
	case *DeleteResponse:
		return assertDeleteResponseEqual(t, expectedResponse, actual)
	}
	return true
}

func runResponseParserTests(t *testing.T, verb Verb, tests []ResponseParserTest) {
	t.Helper()

	for _, test := range tests {
		actualResponse := ParseResponse(verb, []byte(test.input))
		ok := assertResponseEqual(t, test.expected, actualResponse)
		if !ok {
			t.Fatalf("request payload made tests to fail: %s", test.input)
		}
	}
}

func runIndexResponseParserTests(t *testing.T, tests []ResponseParserTest) {
	t.Helper()

	runResponseParserTests(t, INDEX, tests)
}
func runSearchResponseParserTests(t *testing.T, tests []ResponseParserTest) {
	t.Helper()

	runResponseParserTests(t, SEARCH, tests)
}
func runExitResponseParserTests(t *testing.T, tests []ResponseParserTest) {
	t.Helper()

	runResponseParserTests(t, EXIT, tests)
}

func runVersionResponseParserTests(t *testing.T, tests []ResponseParserTest) {
	t.Helper()

	runResponseParserTests(t, VERSION, tests)
}

func runDeleteResponseParserTests(t *testing.T, tests []ResponseParserTest) {
	t.Helper()

	runResponseParserTests(t, DELETE, tests)
}

func TestParseIndexResponse(t *testing.T) {
	tests := []ResponseParserTest{
		{
			"Text has been indexed",
			&IndexResponse{
				Ok: true,
			},
		},
		{
			"Text has NOT been indexed",
			&IndexResponse{
				Ok: false,
			},
		},
		{
			"",
			&IndexResponse{
				Ok: false,
			},
		},
		{
			"  \n",
			&IndexResponse{
				Ok: false,
			},
		},
	}

	runIndexResponseParserTests(t, tests)
}

func TestParseSearchResponse(t *testing.T) {
	tests := []ResponseParserTest{
		{
			`[]`,
			&SearchResponse{
				Documents: []string{},
			},
		},
		{
			`["one item"]`,
			&SearchResponse{
				Documents: []string{
					"one item",
				},
			},
		},
		{
			`["one item", "another item"]`,
			&SearchResponse{
				Documents: []string{
					"one item",
					"another item",
				},
			},
		},
	}

	runSearchResponseParserTests(t, tests)
}

func TestParseExitResponse(t *testing.T) {
	tests := []ResponseParserTest{
		{
			"Bye",
			&ExitResponse{
				Ok: true,
			},
		},
		{
			"BYE",
			&ExitResponse{
				Ok: false,
			},
		},
		{
			" ",
			&ExitResponse{
				Ok: false,
			},
		},
	}

	runExitResponseParserTests(t, tests)
}

func TestParseVersionResponse(t *testing.T) {
	tests := []ResponseParserTest{
		{
			"0.0.0",
			&VersionResponse{
				version: "0.0.0",
			},
		},
		{
			" 0.0.0 ",
			&VersionResponse{
				version: "0.0.0",
			},
		},
	}

	runVersionResponseParserTests(t, tests)
}

func TestParseDeleteResponse(t *testing.T) {
	tests := []ResponseParserTest{
		{
			"Key Removed",
			&DeleteResponse{
				ok: true,
			},
		},
		{
			" Key Removed ",
			&DeleteResponse{
				ok: true,
			},
		},
		{
			" whatever nosense",
			&DeleteResponse{
				ok: false,
			},
		},
	}

	runDeleteResponseParserTests(t, tests)
}
