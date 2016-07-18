package csvutils

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"strings"
	"unicode"
)

func Convert(r io.Reader) ([]map[string]string, error) {
	data, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, err
	}

	keys := replaceAll(data[0])
	var out []map[string]string

	for _, item := range data[1:] {
		m := make(map[string]string)
		for i, k := range keys {
			m[k] = item[i]
		}
		out = append(out, m)
	}
	return out, nil
}

func Unmarshal(data []byte, v interface{}) error {
	r := bytes.NewReader(data)
	m, err := Convert(r)
	if err != nil {
		return err
	}

	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func replaceAll(list []string) []string {
	for i := range list {
		list[i] = replaceName(list[i])
	}
	return list
}

func replaceName(name string) string {
	newString := ""
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			newString += string(r)
		} else {
			newString += "_"
		}
	}
	return strings.ToLower(newString)
}
