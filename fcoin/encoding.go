package fcoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// encode json marshal a struct into io.Reader
func encode(args interface{}) (*bytes.Buffer, error) {
	// According to: https://golang.org/pkg/encoding/json/#Marshal
	// "The map keys are sorted and used as JSON object keys"
	// means json maintains fields alphabetically
	// ref: https://github.com/golang/go/commit/181000896e381f07e8f105eef2667d566729f6eb
	j, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %v", args)
	}
	return bytes.NewBuffer(j), nil
}

// Convert struct to map[string]string for signature, key being json field tag
// url.Values.Encode() keeps keys alphabetically ordered
// Credits: https://gist.github.com/tonyhb/5819315
func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(typ.Field(i).Tag.Get("json"), v)
	}
	return
}
