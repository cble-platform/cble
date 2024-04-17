package model

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

var powerStateMap = map[string]pgrpc.PowerState{
	"on":    pgrpc.PowerState_ON,
	"off":   pgrpc.PowerState_OFF,
	"reset": pgrpc.PowerState_RESET,
}

var powerStateStrMap = map[pgrpc.PowerState]string{
	pgrpc.PowerState_ON:    "on",
	pgrpc.PowerState_OFF:   "off",
	pgrpc.PowerState_RESET: "reset",
}

func MarshalPowerState(v pgrpc.PowerState) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(powerStateStrMap[v]))
	})
}

func UnmarshalPowerState(v interface{}) (pgrpc.PowerState, error) {
	switch v := v.(type) {
	case string:
		return powerStateMap[v], nil
	case int:
		return pgrpc.PowerState(v), nil
	default:
		return pgrpc.PowerState_OFF, fmt.Errorf("%T is not a PowerState", v)
	}
}
