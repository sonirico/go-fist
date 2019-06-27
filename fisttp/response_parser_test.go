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

func assertExitResponseEqual(t *testing.T, expected *ExitResponse, actual Response) bool {
	if actualResponseType, ok := actual.(*ExitResponse); !ok {
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
