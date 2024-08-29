package mongo

import "testing"

func TestStringParsers(t *testing.T) {
    result, err := getArrayParsed(`[{"value": "value1"}, {"value": "value2"}]`)
    if err != nil {
        t.Fatal(err)
    }
    if len(result) != 2 {
        t.Fatal("Incorrect length")
    }
}
