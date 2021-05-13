package apiserver

type Config struct {
	PortAddr string `toml:"port_addr"` // variable for definitions port address
}

func NewConfig() *Config {
	return &Config{
		PortAddr: ":8080", // Default value for port
	}
}
