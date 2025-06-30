package utils

import (
	"regexp"
	"time"
)

func FormatTimestampToDatetime(timestamp int64) string {
	return (time.Unix(timestamp, 0)).Format("02.01.2006 15:04")
}

func FixVKLinks(text string) string {
	re := regexp.MustCompile(`\[(club|public|id)(\d+)\|([^\]]+)\]`)

	return re.ReplaceAllString(text, "[$3](https://vk.com/$1$2)")
}
