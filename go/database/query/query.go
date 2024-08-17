package query

import "strings"

func IsSelectQuery(query string) bool {
	query = strings.TrimSpace(query)

	if len(query) == 0 {
		return false
	}

	return strings.HasPrefix(strings.ToUpper(query), "SELECT")
}

func IsInsertUpdateOrDelete(query string) bool {
	query = strings.TrimSpace(query)

	if len(query) == 0 {
		return false
	}

	if strings.HasPrefix(strings.ToUpper(query), "INSERT") {
		return true
	}
	if strings.HasPrefix(strings.ToUpper(query), "UPDATE") {
		return true
	}
	if strings.HasPrefix(strings.ToUpper(query), "DELETE") {
		return true
	}

	return false
}

func SplitQueries(queries string) []string {
	queryArray := strings.Split(queries, ";")
	for i := range queryArray {
		queryArray[i] = strings.TrimSpace(queryArray[i])
	}

	var result []string
	for _, query := range queryArray {
		if query != "" {
			result = append(result, query)
		}
	}

	return result
}

func ContainsSemicolonInMiddle(s string) bool {
	s = strings.TrimSpace(s)
	if strings.Contains(s, ";") {
		index := strings.Index(s, ";")
		return index != -1 && index < len(s)-1
	}
	return false
}
