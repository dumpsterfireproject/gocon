package gocon

type Config struct {
	values map[string]ConfigValue
}

func NewConfig() Config {
	return Config{
		values: make(map[string]ConfigValue),
	}
}

// return a new config with all variables resolved
func (c Config) Resolve() (Config, error) {
	return c, nil
}

func (c Config) Get(key string) (ConfigValue, bool) {
	if val, ok := c.values[key]; ok {
		return val, true
	}
	return nil, false
}

func (c Config) addValue(key string, v ConfigValue) {
	// TODO: Want to handle defaults and overrides, so we don't ultimately want this simple map
	c.values[key] = v
}

type ConfigValueType int

const (
	ConfigObjectType ConfigValueType = iota
	ConfigStringType
	ConfigNumberType
	ConfigBooleanType
	ConfigArrayType
	ConfigNullType
	configUnresolvedValueType // holds multiple values until resolved
)

// need to track position, for error purposed and also for resolving
type ConfigValue interface {
	Type() ConfigValueType
	String() string
}

type configUnresolvedValue struct {
	value []string
}

func newConfigUnresolvedValue(value string) configUnresolvedValue {
	return configUnresolvedValue{value: []string{value}}
}

func (c configUnresolvedValue) Type() ConfigValueType {
	return configUnresolvedValueType
}

func (c configUnresolvedValue) String() string {
	if len(c.value) > 0 {
		return c.value[len(c.value)-1]
	}
	return ""
}

func (c configUnresolvedValue) addValue(v string) {
	c.value = append(c.value, v)
}

// type ConfigString struct {
// 	value string
// }

// func (c ConfigString) Type() ConfigValueType {
// 	return ConfigStringType
// }
// func (c ConfigString) String() string {
// 	return c.value
// }
