package setting

type Config struct {
	*LogConfig  `mapstructure:"log"`
	*MainConfig `mapstructure:"main"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"fileName"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
	LogType    string `mapstructure:"logType"`
}

type MainConfig struct {
	Port int `mapstructure:"port"`
}
