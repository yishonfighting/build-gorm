package model

import (
	"fmt"
	"strings"
)

var typeForMysqlToGo = map[string]string{
	"int":                "int64",
	"integer":            "int64",
	"tinyint":            "int",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int64",
	"int unsigned":       "int64",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int64",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time.Time or string
	"datetime":           "time.Time", // time.Time or string
	"timestamp":          "time.Time", // time.Time or string
	"time":               "time.Time", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

var str string

func Sql2Struct(t string, s []string) string {
	str = fmt.Sprintf("type %s struct{ \n", formatString(t))
	for _, field := range s {
		sc := strings.Split(field, " ")

		n := ""
		t := ""

		for i := 0; i < len(sc); i++ {
			if i == 0 {
				n = strings.Replace(sc[i], "`", "", 2)
			}
			if i == 1 {
				t = typeForMysqlToGo[strings.Split(sc[i], "(")[0]]
			}
		}

		str += fmt.Sprintf("\t%s \t\t%s  `gorm:\"column:%s\" json:\"%s\" form:\"%s\"`\n", formatString(n), t, n, n, n)
	}
	str += fmt.Sprintf("\tPostfix \t\tstring  `gorm:\"column:postfix\" json:\"postfix\" form:\"postfix\"`\n}\n")
	return str
}
