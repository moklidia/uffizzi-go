package settings

type Settings struct {
	DefaultHost string             `yaml:"default_host"`
	ConfigFile  ConfigFileSettings `yaml:"config_file"`
}

type ConfigFileSettings struct {
	Dir  string `yaml:"dir"`
	Name string `yaml:"name"`
	Ext  string `yaml:"ext"`
}
