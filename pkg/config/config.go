package config

import (
	"flag"
	"os"
)

type Config struct {
	keyMap   map[string]string // key: name, value: cli flag name
	valueMap map[string]string // key: name, value: value
}

func InitConfig() *Config {
	return &Config{
		keyMap:   make(map[string]string),
		valueMap: make(map[string]string),
	}
}

func (c *Config) Add(name, cliFlag string, defaultValue string, description string) {
	// register CLI flag
	flag.String(cliFlag, defaultValue, description)
	// update keyMap
	c.keyMap[name] = cliFlag
}

func (c *Config) Get(name string) (string, bool) {
	v, ok := c.valueMap[name]
	return v, ok
}

func (c *Config) Parse() {
	// parse CLI flags
	flag.Parse()

	// update value map
	for name, cliFlag := range c.keyMap {
		// check cli flag first
		flagValue := flag.Lookup(cliFlag)
		if flagValue != nil && flagValue.Value.String() != "" && flagValue.Value.String() != flagValue.DefValue {
			c.valueMap[name] = flagValue.Value.String()
			return
		}

		// check env var value
		envValue := os.Getenv(name)
		if envValue != "" {
			c.valueMap[name] = envValue
			return
		}

		// if both cli flag and env var not exist, use default value
		c.valueMap[name] = flagValue.DefValue
	}
}

func (c *Config) Configuration() []string {
	var values []string
	for name, value := range c.valueMap {
		values = append(values, name+"="+value)
	}
	return values
}
