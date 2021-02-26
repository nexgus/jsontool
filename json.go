package jsontool

import (
	"encoding/json"
	"io/ioutil"
)

// JSONDumpB dump a JSON structure to a slice of byte.
func JSONDumpB(j interface{}, indent ...string) ([]byte, error) {
	var (
		b   []byte
		err error
	)
	if indent == nil {
		b, err = json.Marshal(j)
	} else {
		b, err = json.MarshalIndent(j, "", indent[0])
	}
	return b, err
}

// JSONDumpF dump a JSON structrue to a file.
func JSONDumpF(filepath string, j interface{}, indent ...string) error {
	var (
		b   []byte
		err error
	)
	if indent == nil {
		b, err = json.Marshal(j)
	} else {
		b, err = json.MarshalIndent(j, "", indent[0])
	}
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath, b, 0664)
	return err
}

// JSONDumpS dump a JSON structrue to a string.
func JSONDumpS(j interface{}, indent ...string) string {
	var (
		b   []byte
		err error
	)
	if indent == nil {
		b, err = json.Marshal(j)
	} else {
		b, err = json.MarshalIndent(j, "", indent[0])
	}
	if err != nil {
		return ""
	}
	return string(b)
}

// JSONLoadB loads JSON from a slice of byte.
func JSONLoadB(b []byte) (interface{}, error) {
	var j interface{}
	err := json.Unmarshal(b, &j)
	return j, err
}

// JSONLoadF load JSON from a file.
func JSONLoadF(filepath string) (interface{}, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	j, err := JSONLoadB(b)
	return j, err
}

// JSONLoadS loads JSON from a string.
func JSONLoadS(s string) (interface{}, error) {
	j, err := JSONLoadB([]byte(s))
	return j, err
}
