package helpers

import (
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func FromSqlTime(sqlTime sql.NullTime) time.Time {
	if sqlTime.Valid {
		return sqlTime.Time
	}
	return time.Time{}
}

func ToSqlTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Valid: true,
		Time:  t,
	}
}

func FromSqlString(sqlString sql.NullString) string {
	if sqlString.Valid {
		return sqlString.String
	}
	return ""
}

func ToSqlString(s string) sql.NullString {
	return sql.NullString{
		Valid:  true,
		String: s,
	}
}

func ToProtoTime(t time.Time) timestamppb.Timestamp {
	return timestamppb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
