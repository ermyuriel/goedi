package goedi

import (
	"bytes"
	"regexp"
	"strings"
)

func ParseSegments(edi []byte, segmentDelimiter, valueDelimiter string) map[string][][]string {
	if len(edi) == 0 {
		return nil
	}

	var re = regexp.MustCompile(`[^0-9a-zA-Z]`)

	m := make(map[string][][]string)
	segments := bytes.Split(edi, []byte(segmentDelimiter))
	for _, s := range segments {
		values := strings.Split(string(s), valueDelimiter)
		k := re.ReplaceAllString(values[0], ``)

		m[k] = append(m[k], values[1:])

	}

	return m

}
