package mnubo

import (
	"os"
	"testing"
)

type SelectOperation struct {
	Count string `json:"count"`
}

type SimpleQuery struct {
	From   string            `json:"from"`
	Select []SelectOperation `json:"select"`
}

type SearchResultsColumn struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

type SearchResults struct {
	Columns []SearchResultsColumn `json:"columns"`
	Rows    [][]interface{}       `json:"rows"`
}

func TestSearchBasic(t *testing.T) {
	m := NewClient(os.Getenv("MNUBO_CLIENT_ID"), os.Getenv("MNUBO_CLIENT_SECRET"), os.Getenv("MNUBO_HOST"))
	m.getAccessTokenWithScopeAll()

	var results = [3]SearchResults{}
	cases := []struct {
		Error error
	}{
		{
			Error: m.createBasicQuery(SimpleQuery{
				From: "event",
				Select: []SelectOperation{
					{
						Count: "*",
					},
				},
			}, &results[0]),
		},
		{
			Error: m.createBasicQueryWithString(`
				{
				    "from": "event",
				    "select": [
				        { "count": "*" }
				    ]
				}
			`, &results[1]),
		},
	}

	for i, c := range cases {
		if c.Error != nil {
			t.Errorf("%d, could not create basic query: %t", i, c.Error)
		}

		if len(results[i].Rows) != 1 || len(results[i].Rows[0]) != 1 {
			t.Errorf("%d, expecting results to have a count in firt row and cell", i)
		}

		t.Logf("Got Count %d", results[i].Rows[0][0])
	}
}
