package gocon

type configUnresolvedValue struct {
	value []string
}

func newConfigUnresolvedValue(value string) *configUnresolvedValue {
	return &configUnresolvedValue{value: []string{value}}
}

func (c *configUnresolvedValue) Type() ConfigValueType {
	return configUnresolvedValueType
}

func (c *configUnresolvedValue) String() string {
	if len(c.value) > 0 {
		return c.value[len(c.value)-1]
	}
	return ""
}

func (c *configUnresolvedValue) addValue(v string) {
	c.value = append(c.value, v)
}
