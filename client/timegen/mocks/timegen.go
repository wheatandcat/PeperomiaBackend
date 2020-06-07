// Code generated  DO NOT EDIT.
package mock_timegen

import (
	"time"
)

// Time has generating method.
type Time struct {
}

// Now 現在時刻を取得する
func (*Time) Now() time.Time {
	return time.Date(2020, 1, 1, 00, 00, 00, 0, time.Local)
}
