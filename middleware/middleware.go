package middleware

import (
	"gorm.io/gorm"
)

func RecordCheck(rows *gorm.DB) (string, error) {
	count := rows.RowsAffected
	err := rows.Error
	// fmt.Println(rows)
	// fmt.Println(count)
	if count == 0 {
		return "no-record", err
	} else if err != nil {
		return "query-error", err
	} else {
		return "ok", err
	}
}
