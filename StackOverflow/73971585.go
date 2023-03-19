package __

import (
	"fmt"
	"strings"
)

func SQLCreateReturn(tablename string, params map[string]interface{}, returnparam string) {
	var insertQuery string
	var keys []string
	var values []string

	insertQuery = "INSERT INTO " + tablename

	for key, value := range params {
		keys = append(keys, key)
		values = append(values, "'"+value.(fmt.Stringer).String()+"'")
	}
	keylist := " (" + strings.Join(keys, ", ") + ")"
	valuelist := "(" + strings.Join(values, ", ") + ")"

	query := insertQuery + keylist + " VALUES " + valuelist + " RETURNING " + returnparam
	_ = query
}
