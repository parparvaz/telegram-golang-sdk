package common

import "bytes"

func ToJSONList(v []byte) []byte {
	if len(v) > 0 && v[0] == '{' {
		var b bytes.Buffer
		b.Write([]byte("["))
		b.Write(v)
		b.Write([]byte("]"))
		return b.Bytes()
	}
	return v
}
