package jsontool

import (
	"encoding/json"
	"io/ioutil"
)

// DumpB dump a JSON structure to a slice of byte.
func DumpB(j interface{}, indent ...string) ([]byte, error) {
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

// DumpF dump a JSON structrue to a file.
func DumpF(filepath string, j interface{}, indent ...string) error {
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

// DumpS dump a JSON structrue to a string.
func DumpS(j interface{}, indent ...string) string {
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

// LoadB loads JSON from a slice of byte.
func LoadB(b []byte) (interface{}, error) {
	var j interface{}
	err := json.Unmarshal(b, &j)
	return j, err
}

// LoadF load JSON from a file.
func LoadF(filepath string) (interface{}, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	j, err := LoadB(b)
	return j, err
}

// LoadS loads JSON from a string.
func LoadS(s string) (interface{}, error) {
	j, err := LoadB([]byte(s))
	return j, err
}
