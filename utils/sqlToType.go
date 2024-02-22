package utils

import "database/sql"

func SqlNullStringToString(sqlNull sql.NullString) *string {
	var result *string
	if sqlNull.Valid {
		result = &sqlNull.String
	}

	return result
}

func StringToSqlNullString(str *string) sql.NullString {
	result := sql.NullString{}

	if str != nil {
		result.String = *str
		result.Valid = true
	}

	return result
}
