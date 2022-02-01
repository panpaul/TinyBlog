package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
	"log"
	"server/utils"
)

func Setup(configFile string) {
	CONF = new(Config)
	CONF.SetupConfig(configFile)
	setupZap()
}

func (c *Config) SetupConfig(configFile string) {
	// Zap not available Now
	data, err := utils.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err.Error())
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		log.Fatalf("Couldn't unmarshal the config file: %s\n", err.Error())
	}
	c.FirstRun = !utils.FileExist("./resource/data/FIRST_RUN")
}

func (c *Config) WriteConfig(configFile string) {
	out, err := yaml.Marshal(c)
	if err != nil {
		LOG.Warn("failed to marshal config", zap.Error(err))
		return
	}
	_ = utils.WriteFile(configFile, out)
}

func setupZap() {
	cfg := zap.Config{
		Level: zap.NewAtomicLevelAt(
			utils.If(CONF.Development, zapcore.DebugLevel, zapcore.InfoLevel).(zapcore.Level),
		),
		Development:      CONF.Development,
		Encoding:         "console", // or json
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	LOG, err = cfg.Build()
	if err != nil {
		log.Fatalf("Failed to setup logger (Zap): %s\n", err.Error())
	}
}
