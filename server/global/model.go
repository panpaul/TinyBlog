package global

type ConfigWebServer struct {
	Address string `yaml:"Address"`
	Port    int    `yaml:"Port"`
}

type ConfigJwt struct {
	JwtSecret     string `yaml:"JwtSecret"`
	JwtExpireHour int    `yaml:"JwtExpireHour"`
}

type ConfigRedis struct {
	Db       int    `yaml:"Db"`
	Address  string `yaml:"Address"`
	Password string `yaml:"Password"`
}

type ConfigDatabase struct {
	Address  string `yaml:"Address"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
	Prefix   string `yaml:"Prefix"`
}

type Config struct {
	WebServer   ConfigWebServer `yaml:"WebServer"`
	Jwt         ConfigJwt       `yaml:"Jwt"`
	Redis       ConfigRedis     `yaml:"Redis"`
	Database    ConfigDatabase  `yaml:"Database"`
	Development bool            `yaml:"Development"`
	FirstRun    bool            `yaml:"-"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Body interface{} `json:"body"`
}
