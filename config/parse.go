package config

type Parser interface {
	Parse(path string)
}

type ConfigParser struct {
}

func (cp *ConfigParser) Parse(path string) {
	
}
