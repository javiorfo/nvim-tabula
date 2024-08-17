package query

import "testing"

func TestQuery(t *testing.T) {
    q := "select something; select other  "
    arr := SplitQueries(q)

    if arr[0] != "select something" {
        t.Fatalf("First row incorrect: %s", arr[0])
    }
    if arr[1] != "select other" {
        t.Fatalf("Second row incorrect: %s", arr[0])
    }

    q = "select one;"
    arr = SplitQueries(q)
    if arr[0] != "select one" {
        t.Fatalf("Select one incorrect: %s", arr[0])
    }
}
