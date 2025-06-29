package utils

import "time"

func FormatTimestampToDatetime(timestamp int64) string {
	return (time.Unix(timestamp, 0)).Format("02.01.2006 15:04")
}
