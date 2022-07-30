package factory

import (
	"fmt"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/config"
	"github.com/spf13/viper"
)

var MonitorCfg *config.Config

func InitConfigFactory() error {
	viper.SetConfigName("cfg")             // name of config file (without extension)
	viper.SetConfigType("yaml")            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../../../config") // path to look for the config file in
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig()            // Find and read the config file
	if err != nil {                        // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println(viper.GetString("out_path"))
	fmt.Println(viper.GetInt("Interval"))
	fmt.Println(viper.GetStringSlice("processors"))

	MonitorCfg.Processors = viper.GetStringSlice("processors")
	MonitorCfg.Interval = viper.GetInt("Interval")
	MonitorCfg.OutPath = viper.GetString("out_path")

	return nil
}

func init() {
	MonitorCfg = config.NewConfig()
}
