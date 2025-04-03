package config

import (
	"flag"
	"os"
	"strings"
)

type Config struct {
	prefix   string
	keyMap   map[string]string // key: name, value: cli flag name
	valueMap map[string]string // key: name, value: value
}

func InitConfig(prefix string) *Config {
	return &Config{
		prefix: prefix,
	}
}

func (c Config) prefixEnvVarName(name string) string {
	envVar := strings.Join([]string{c.prefix, name}, "-")
	return envVar
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
	for name, cliFlag := range c.keyMap {
		// check cli flag first
		flagValue := flag.Lookup(cliFlag)
		if flagValue != nil && flagValue.Value.String() != "" {
			c.valueMap[name] = flagValue.Value.String()
			return
		}

		// check env var value
		envValue := os.Getenv(c.prefixEnvVarName(name))
		if envValue != "" {
			c.valueMap[name] = envValue
		}
	}
}
