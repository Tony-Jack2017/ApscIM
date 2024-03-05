package common

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type BaseTime struct {
	CreatedAt LocalTime `json:"createAt"`
	UpdatedAt LocalTime `json:"updatedAt"`
	DeletedAt LocalTime `json:"deletedAt"`
}

type LocalTime time.Time

// override MarshalJSON

func (t LocalTime) MarshalJSON() ([]byte, error) {
	localTime := time.Time(t)
	return []byte(fmt.Sprintf("\"%v\"", localTime.Format("2006-01-02 15:04:05"))), nil
}

// when storage the timestamp convert time string to sql timestamp

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	localTime := time.Time(t)
	if localTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return localTime, nil
}

// convert sql timestamp to string which is yyyy-mm-dd HH:MM:SS format

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can't convert %v to timestamp", v)
}

func (t LocalTime) ParseTime(format string, isUnix bool) string {
	localTime := time.Time(t)
	if isUnix {
		return fmt.Sprintf("%s", localTime.UnixNano())
	} else {
		return localTime.Format(judgeTimeFormat(format))
	}
}

func judgeTimeFormat(format string) string {
	return "2006-01-02 15:04:05"
}
