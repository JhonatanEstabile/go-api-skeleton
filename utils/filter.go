package utils

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

type QueryParam struct {
	Field    string
	Operator string
	Filter   string
}

func ParseQueryParams(queryString string) (string, []interface{}) {
	var result []QueryParam
	queryParams := strings.Split(queryString, "&")
	for _, queryParam := range queryParams {
		parts := strings.Split(queryParam, ",")
		if len(parts) == 2 {
			result = append(result, QueryParam{
				Field:    parts[0],
				Operator: parts[1],
				Filter:   "",
			})
		} else if len(parts) == 3 {
			result = append(result, QueryParam{
				Field:    parts[0],
				Operator: parts[1],
				Filter:   parts[2],
			})
		}
	}

	var whereClause []string
	var args []interface{}
	for _, qp := range result {
		whereClause = append(whereClause, fmt.Sprintf("%s %s ?", qp.Field, qp.Operator))
		args = append(args, decodeQueryParam(qp.Filter))
	}
	where := ""
	if len(whereClause) > 0 {
		where = "WHERE " + strings.Join(whereClause, " AND ")
	}
	return where, args
}

func decodeQueryParam(param string) string {
	decoded, err := url.QueryUnescape(param)
	if err != nil {
		log.Fatal(err)
	}
	return decoded
}
