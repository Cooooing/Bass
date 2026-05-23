package enum

// Mapping 维护内部 string 枚举与生成的 proto 枚举之间的双向映射。
// 内部包应持久化 string 值，只在 API 边界转换为 proto 值。
type Mapping[E ~string, P ~int32] struct {
	toProto   map[E]P
	fromProto map[P]E
	values    []string
}

// Entry 定义一个内部枚举值对应的 proto 枚举值。
type Entry[E ~string, P ~int32] struct {
	Proto P
}

// NewMapping 根据内部枚举集合构建映射。
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

// ToEnum 将 proto 枚举值转换为内部枚举值。
func (m *Mapping[E, P]) ToEnum(v P) (E, bool) {
	val, ok := m.fromProto[v]
	return val, ok
}

// MustToEnum 将 proto 枚举值转换为内部枚举值；未知值返回零值。
func (m *Mapping[E, P]) MustToEnum(v P) E {
	val, _ := m.fromProto[v]
	return val
}

// ToProto 将内部枚举值转换为 proto 枚举值。
func (m *Mapping[E, P]) ToProto(v E) (P, bool) {
	val, ok := m.toProto[v]
	return val, ok
}

// MustToProto 将内部枚举值转换为 proto 枚举值；未知值返回零值。
func (m *Mapping[E, P]) MustToProto(v E) P {
	return m.toProto[v]
}

// EnumValues 返回所有有效的内部枚举字符串值。
func (m *Mapping[E, P]) EnumValues() []string {
	return m.values
}
