package model

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sirupsen/logrus"
)

type StrMap map[string]string

func MarshalStrMap(val map[string]string) graphql.Marshaler {
	logrus.Printf("marshalling: %+v", val)
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalStrMap(v interface{}) (map[string]string, error) {
	logrus.Printf("unmarshalling: %+v", v)
	if m, ok := v.(map[string]string); ok {
		return m, nil
	}

	return nil, fmt.Errorf("%T is not a map", v)
}
