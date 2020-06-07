package timegen

import (
	"time"
)

// TimeGenerator 現在日時取得
type TimeGenerator interface {
	Now() time.Time
}

// Time has generating method.
type Time struct {
}

// Now 現在時刻を取得する
func (*Time) Now() time.Time {
	return time.Now()
}
