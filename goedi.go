package goedi

import (
	"regexp"
	"strings"
)

func parseSegments(edi string, segmentDelimiter, valueDelimiter string) *map[string][][]string {
	if len(edi) == 0 {
		return nil
	}

	var re = regexp.MustCompile(`[^0-9a-zA-Z]`)

	m := make(map[string][][]string)
	segments := strings.Split(edi, segmentDelimiter)
	for _, s := range segments {
		values := strings.Split(s, valueDelimiter)
		k := re.ReplaceAllString(values[0], ``)

		m[k] = append(m[k], values[1:])
	}

	return &m
}
