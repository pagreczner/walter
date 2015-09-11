package main

import (
	"fmt"
	"os"
	"strings"
)

const template string = `
ALTER TYPE ENUM_TYPE RENAME TO __ENUM_TYPE;
CREATE TYPE ENUM_TYPE AS ENUM(ENUM_VALUES);

ALTER TABLE TABLE_NAME RENAME COLUMN COLUMN_NAME to _COLUMN_NAME;
ALTER TABLE TABLE_NAME ADD COLUMN_NAME ENUM_TYPE NOT NULL DEFAULT ENUM_DEFAULT_VALUE;
UPDATE TABLE_NAME SET COLUMN_NAME = _COLUMN_NAME::text::ENUM_TYPE;
ALTER TABLE TABLE_NAME DROP COLUMN _COLUMN_NAME;
DROP TYPE __ENUM_TYPE;
`

func main() {
	if len(os.Args) < 4 {
		fmt.Println(`
			Usage:
			walter <type_name> <table_name> <column_name> [...enum_values]
			`)
		os.Exit(0)
	}

	typeName := os.Args[1]
	tableName := os.Args[2]
	columnName := os.Args[3]

	var enumValues []string
	var defaultEnumValue string
	if len(os.Args) > 4 {
		for _, v := range os.Args[4:] {
			enumValues = append(enumValues, "'"+v+"'")
		}
		defaultEnumValue = os.Args[4]
	}

	enumString := strings.Join(enumValues, ",")

	output := strings.Replace(template, "ENUM_TYPE", typeName, -1)
	output = strings.Replace(output, "TABLE_NAME", tableName, -1)
	output = strings.Replace(output, "COLUMN_NAME", columnName, -1)
	output = strings.Replace(output, "ENUM_TYPE", typeName, -1)
	output = strings.Replace(output, "ENUM_VALUES", enumString, -1)
	output = strings.Replace(output, "ENUM_DEFAULT_VALUE", defaultEnumValue, -1)

	fmt.Println(output)

}
