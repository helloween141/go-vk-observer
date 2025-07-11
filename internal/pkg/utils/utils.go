package utils

import (
	"regexp"
	"strings"
	"time"
)

func ExtractVkSlug(input string) string {
	input = strings.TrimSpace(input)

	if idx := strings.IndexAny(input, "?#"); idx != -1 {
		input = input[:idx]
	}

	input = strings.TrimPrefix(input, "https://")
	input = strings.TrimPrefix(input, "http://")
	input = strings.TrimPrefix(input, "www.")
	input = strings.TrimPrefix(input, "vk.com/")
	input = strings.TrimPrefix(input, "vk.com")

	if idx := strings.Index(input, "/"); idx != -1 {
		input = input[:idx]
	}

	return input
}

func FormatTimestampToDatetime(timestamp int64) string {
	return (time.Unix(timestamp, 0)).Format("02.01.2006 15:04")
}

func FormatPostLinks(text string) string {
	// [club12345|Name]
	reVK := regexp.MustCompile(`\[(club|public|id)(\d+)\|([^\]]+)\]`)
	text = reVK.ReplaceAllString(text, `<a href="https://vk.com/$1$2">$3</a>`)

	// [https://vk.com/whatever|Name]
	reFullURL := regexp.MustCompile(`\[(https?://[^\|\]]+)\|([^\]]+)\]`)
	text = reFullURL.ReplaceAllString(text, `<a href="$1">$2</a>`)

	return text
}

func FormatRepostLinks(text string) string {
	re := regexp.MustCompile(`\[[^\|\[\]]+\|([^\[\]]+)\]`)

	return re.ReplaceAllString(text, "$1")
}
