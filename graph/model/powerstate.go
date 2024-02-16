package model

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	providerGRPC "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

var powerStateMap = map[string]providerGRPC.PowerState{
	"on":    providerGRPC.PowerState_ON,
	"off":   providerGRPC.PowerState_OFF,
	"reset": providerGRPC.PowerState_RESET,
}

var powerStateStrMap = map[providerGRPC.PowerState]string{
	providerGRPC.PowerState_ON:    "on",
	providerGRPC.PowerState_OFF:   "off",
	providerGRPC.PowerState_RESET: "reset",
}

func MarshalPowerState(v providerGRPC.PowerState) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(powerStateStrMap[v]))
	})
}

func UnmarshalPowerState(v interface{}) (providerGRPC.PowerState, error) {
	switch v := v.(type) {
	case string:
		return powerStateMap[v], nil
	case int:
		return providerGRPC.PowerState(v), nil
	default:
		return providerGRPC.PowerState_OFF, fmt.Errorf("%T is not a PowerState", v)
	}
}
