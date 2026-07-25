package str

import "testing"

func TestSonyflake(
	t *testing.T,
) {
	sf, err := NewSonyflake()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		id, _ := sf.NextID()
		t.Log(id)
	}
}
