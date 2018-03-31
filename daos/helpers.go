package daos

import (
	"fmt"
	"net/url"
	"reflect"
)

func queryStringToWhereClause(queryString url.Values) string {
	var where string

	if len(queryString) > 0 {
		for k, v := range queryString {
			if len(where) > 0 {
				where += " AND"
			}
			t := reflect.ValueOf(v[0]).Kind()
			fmt.Println(t)
			switch t {
			case reflect.String:
				where += fmt.Sprintf(" %s='%s'", k, v[0])
			case reflect.Int:
				where += fmt.Sprintf(" %s=%s", k, v[0])
			}
		}
	}

	return "WHERE" + where
}
