package model

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

type StrMap map[string]string

func MarshalStrMap(val map[string]string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalStrMap(v interface{}) (map[string]string, error) {
	if m, ok := v.(map[string]interface{}); ok {
		strMap := make(map[string]string)
		for k, v := range m {
			value, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("key %s contains non-string value", k)
			}
			strMap[k] = value
		}
		return strMap, nil
	}

	return nil, fmt.Errorf("%T is not a map", v)
}
