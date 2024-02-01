package models

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalObject(val Object) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalObject(v interface{}) (Object, error) {
	if m, ok := v.(Object); ok {
		return m, nil
	}

	return Object{}, fmt.Errorf("%T is not a map", v)
}
