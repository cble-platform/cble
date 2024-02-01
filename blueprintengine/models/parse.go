package models

var (
	blueprintVariableTypeMap = map[string]BlueprintVariableType{
		"STRING": BlueprintVariableType_STRING,
		"INT":    BlueprintVariableType_INT,
	}
)

func ParseBlueprintVariableType(value string) (BlueprintVariableType, bool) {
	t, ok := blueprintVariableTypeMap[value]
	return t, ok
}
