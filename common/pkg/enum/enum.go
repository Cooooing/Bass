package enum

// Mapping 内部枚举(string) ↔ proto枚举(int32) 双向映射
type Mapping[E ~string, P ~int32] struct {
	toProto   map[E]P
	fromProto map[P]E
	values    []string
}

type Entry[E ~string, P ~int32] struct {
	Proto P
}

// NewMapping 创建映射，传入内部枚举集
func NewMapping[E ~string, P ~int32](entries map[E]Entry[E, P]) *Mapping[E, P] {
	m := &Mapping[E, P]{
		toProto:   make(map[E]P, len(entries)),
		fromProto: make(map[P]E, len(entries)),
		values:    make([]string, 0, len(entries)),
	}
	for val, entry := range entries {
		protoVal := entry.Proto
		m.toProto[val] = protoVal
		m.fromProto[protoVal] = val
		m.values = append(m.values, string(val))
	}
	return m
}

// ToEnum proto 枚举值 → 内部枚举
func (m *Mapping[E, P]) ToEnum(v P) (E, bool) {
	val, ok := m.fromProto[v]
	return val, ok
}

// MustToEnum proto 枚举值 → 内部枚举，非法时返回零值
func (m *Mapping[E, P]) MustToEnum(v P) E {
	val, _ := m.fromProto[v]
	return val
}

// ToProto 内部枚举 → proto 枚举值
func (m *Mapping[E, P]) ToProto(v E) (P, bool) {
	val, ok := m.toProto[v]
	return val, ok
}

// MustToProto 内部枚举 → proto 枚举值，非法时返回零值
func (m *Mapping[E, P]) MustToProto(v E) P {
	return m.toProto[v]
}

// EnumValues 返回所有合法值
func (m *Mapping[E, P]) EnumValues() []string {
	return m.values
}
