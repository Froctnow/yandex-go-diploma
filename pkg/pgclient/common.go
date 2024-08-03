package pgclient

import "github.com/jmoiron/sqlx"

func ListValuesFromRows[T any](rows *sqlx.Rows) ([]T, error) {
	result := make([]T, 0)
	for rows.Next() {
		var val T
		err := rows.StructScan(&val)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return result, nil
}

func StructValueFromRows[T any](rows *sqlx.Rows) (T, error) {
	var val T
	if rows.Next() {
		err := rows.StructScan(&val)
		if err != nil {
			return val, err
		}
	}
	err := rows.Err()
	if err != nil {
		return val, err
	}
	if err = rows.Close(); err != nil {
		return val, err
	}
	return val, nil
}

func ValueFromRows[T any](rows *sqlx.Rows) (T, error) {
	var val T
	if rows.Next() {
		err := rows.Scan(&val)
		if err != nil {
			return val, err
		}
	}
	err := rows.Err()
	if err != nil {
		return val, err
	}
	if err = rows.Close(); err != nil {
		return val, err
	}
	return val, nil
}
