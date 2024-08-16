package query

import "strings"

func IsSelectQuery(query string) bool {
	query = strings.TrimSpace(query)

	if len(query) == 0 {
		return false
	}

	return strings.HasPrefix(strings.ToUpper(query), "SELECT")
}
