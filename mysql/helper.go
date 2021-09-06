package mysql

import (
	"database/sql"

	"golang.org/x/xerrors"
)

// LastInsertedID extracts the id from the query result.
// If the entity was not created.
func LastInsertedID(res sql.Result) (uint32, error) {
	id, err := res.LastInsertId()
	if err != nil {
		return 0, xerrors.Errorf("failed to LastInsertId: %w", err)
	}

	if id == 0 {
		return 0, xerrors.New("the entity was not created")
	}

	return uint32(id), nil
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// ToNullInt64 returns sql.NullInt64. The int is considered valid if it's not equal 0.
func ToNullInt64(u uint32) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(u), Valid: u != 0}
}

func ToNullIInt64(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: i != 0}
}

// ToUint32 returns uint32 extracted from the sql.NullInt64
func ToUint32(n sql.NullInt64) uint32 {
	return uint32(n.Int64)
}

func joinFields(sls ...[]interface{}) []interface{} {
	res := sls[0]
	for i := 1; i < len(sls); i++ {
		res = append(res, sls[i]...)
	}
	return res
}
