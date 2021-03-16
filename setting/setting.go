package setting

type Config struct {
	*LogConfig   `mapstructure:"log"`
	*MainConfig  `mapstructure:"main"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
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
	Port    int    `mapstructure:"port"`
	Mode    string `mapstucture:"mode"`
	Addr    string `mapstructure:"address"`
	Context string `mapstructure:"context"`
}

type MysqlConfig struct {
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DbName       string `mapstructure:"db"`
	OtherOptions string `mapstructure:"otherOptions"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type RedisConfig struct {
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	PoolSize int    `mapstructure:"poolSize"`
	DB       int    `mapstructure:"db"`
	Cluster  bool   `mapstructure:"cluster"`
}
