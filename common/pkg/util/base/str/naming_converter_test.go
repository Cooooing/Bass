package str

import (
	"reflect"
	"testing"
)

func TestNaming(t *testing.T) {
	tests := []struct {
		name      string
		converter func(string) string
		input     []string
		expected  []string
	}{
		{"UpperCamelCase", ToUpperCamelCase, []string{"myVariableName", "MyVariableName", "my_variable_name"}, []string{"MyVariableName", "MyVariableName", "MyVariableName"}},
		{"LowerCamelCase", ToLowerCamelCase, []string{"myVariableName", "MyVariableName", "my_variable_name"}, []string{"myVariableName", "myVariableName", "myVariableName"}},
		{"SnakeCase", ToSnakeCase, []string{"myVariableName", "MyVariableName", "my_variable_name"}, []string{"my_variable_name", "my_variable_name", "my_variable_name"}},
		{"KebabCase", ToKebabCase, []string{"myVariableName", "MyVariableName", "my_variable_name"}, []string{"my-variable-name", "my-variable-name", "my-variable-name"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, tc := range tt.input {
				result := tt.converter(tc)
				if !reflect.DeepEqual(result, tt.expected[i]) {
					t.Errorf("Expected %s, but got %s", tt.expected[i], result)
				}
			}
		})
	}

}
