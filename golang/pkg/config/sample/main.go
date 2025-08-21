package main

import (
	"cfg/ce"
	"cfg/de"
	"cfg/vp"
	"flag"
)

func main() {
	var (
		cleanEnvPath string
		dotEnvPath   string
		viperPath    string
	)

	flag.StringVar(&cleanEnvPath, "cleanenv", "./ce/config.yaml", "path to clean env config file")
	flag.StringVar(&dotEnvPath, "dotenv", "./de/config.yaml", "path to dot env config file")
	flag.StringVar(&viperPath, "viper", "./vp/config.yaml", "path to viper config file")
	flag.Parse()

	ce.LoadConfigCleanEnv(cleanEnvPath)
	de.LoadConfigDotEnv(dotEnvPath)
	vp.LoadConfigViper(viperPath)
}
