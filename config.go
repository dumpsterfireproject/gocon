package hocon

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

type ConfigValueType int

const (
	ConfigObjectType ConfigValueType = iota
	ConfigStringType
	ConfigNumberType
	ConfigBooleanType
	ConfigArrayType
	ConfigNullType
)

type ConfigValue interface {
	Type() ConfigValueType
	String() string
}

type ConfigString struct {
	value string
}

func NewConfigString(value string) ConfigString {
	return ConfigString{value: value}
}

func (c ConfigString) Type() ConfigValueType {
	return ConfigStringType
}
func (c ConfigString) String() string {
	return c.value
}
