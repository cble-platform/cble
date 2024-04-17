package model

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cble-platform/cble-backend/engine/models"
)

type VarTypeMap map[string]models.BlueprintVariableType

func MarshalVarTypeMap(val map[string]models.BlueprintVariableType) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalVarTypeMap(v interface{}) (map[string]models.BlueprintVariableType, error) {
	if m, ok := v.(map[string]interface{}); ok {
		varTypes := make(map[string]models.BlueprintVariableType)
		for k, v := range m {
			value, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("variable %s type is not string", k)
			}
			varTypes[k], ok = models.ParseBlueprintVariableType(value)
			if !ok {
				return nil, fmt.Errorf("variable %s has invalid type", k)
			}
		}
		return varTypes, nil
	}

	return nil, fmt.Errorf("%T is not a map", v)
}
