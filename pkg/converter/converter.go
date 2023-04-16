package converter

import (
	"errors"
	"fmt"
	"strings"
)

func ToCamelCase(text string) string {
	var camelCase string

	for i, val := range strings.Split(text, "_") {
		if len(val) == 0 {
			continue
		}

		if i == 0 {
			camelCase += strings.ToLower(val[:1]) + val[1:]
			continue
		}

		camelCase += strings.ToUpper(val[:1]) + val[1:]
	}

	return camelCase
}

func ToPascalCase(text string) string {
	camelCase := ToCamelCase(text)

	if len(camelCase) > 0 {
		return strings.ToUpper(camelCase[:1]) + camelCase[1:]
	}

	return ""
}

const (
	IntType    = "int"
	FloatType  = "float64"
	StringType = "string"
	BoolType   = "bool"
	StructType = "struct"
)

func PgToGoType(typ string) (string, error) {
	switch typ {
	case "smallint", "integer", "bigint",
		"smallserial", "serial", "bigserial":
		return IntType, nil

	case "decimal", "numeric", "real", "double precision":
		return FloatType, nil

	case "character varying", "character", "text", "uuid":
		return StringType, nil

	case "boolean":
		return BoolType, nil

	default:
		return "", errors.New(fmt.Sprintf("тип %s не поддерживается", typ))
	}
}
