package config

type Config struct {
	OutPath    string   `yaml:"out_path"`
	Processors []string `yaml:"processors"`
	Interval   int      `yaml:"interval"`
}

func NewConfig() *Config {
	c := new(Config)
	c.Interval = 2
	c.OutPath = "out.json"
	c.Processors = []string{"smf", "amf"}

	return c
}

func (c *Config) Validate() error { return nil }
