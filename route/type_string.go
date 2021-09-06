// Code generated by "enumer -type=Type -transform=snake -output=type_string.go"; DO NOT EDIT.

package route

import (
	"fmt"
)

const _TypeName = "boulderlead"

var _TypeIndex = [...]uint8{0, 7, 11}

const _TypeLowerName = "boulderlead"

func (i Type) String() string {
	if i >= Type(len(_TypeIndex)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _TypeName[_TypeIndex[i]:_TypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TypeNoOp() {
	var x [1]struct{}
	_ = x[Boulder-(0)]
	_ = x[Lead-(1)]
}

var _TypeValues = []Type{Boulder, Lead}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:7]:       Boulder,
	_TypeLowerName[0:7]:  Boulder,
	_TypeName[7:11]:      Lead,
	_TypeLowerName[7:11]: Lead,
}

var _TypeNames = []string{
	_TypeName[0:7],
	_TypeName[7:11],
}

// TypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TypeString(s string) (Type, error) {
	if val, ok := _TypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

// TypeValues returns all values of the enum
func TypeValues() []Type {
	return _TypeValues
}

// TypeStrings returns a slice of all String values of the enum
func TypeStrings() []string {
	strs := make([]string, len(_TypeNames))
	copy(strs, _TypeNames)
	return strs
}

// IsAType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Type) IsAType() bool {
	for _, v := range _TypeValues {
		if i == v {
			return true
		}
	}
	return false
}