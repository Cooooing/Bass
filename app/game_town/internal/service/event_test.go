package service

import "testing"

func TestNewEventPayloadNormalizesTypedSlices(
	t *testing.T,
) {
	payload, err := newEventPayload(map[string]any{
		"character_traits": []string{"scavenger", "mecha-whisperer"},
		"suggested_actions": []any{
			map[string]any{
				"label": "Ask more",
				"targets": []any{
					map[string]any{"type": "npc", "id": int64(12)},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	traits := payload.GetFields()["character_traits"].GetListValue().GetValues()
	if len(traits) != 2 || traits[0].GetStringValue() != "scavenger" {
		t.Fatalf("unexpected traits: %#v", traits)
	}
}

func TestInt64ValueSupportsPayloadNumbers(
	t *testing.T,
) {
	tests := []struct {
		name  string
		value any
		want  int64
	}{
		{name: "int64", value: int64(7), want: 7},
		{name: "int", value: int(8), want: 8},
		{name: "float64", value: float64(9), want: 9},
		{name: "string", value: "10", want: 10},
		{name: "float string", value: "11.0", want: 11},
		{name: "bad string", value: "bad", want: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := int64Value(test.value); got != test.want {
				t.Fatalf("int64Value() = %d, want %d", got, test.want)
			}
		})
	}
}
