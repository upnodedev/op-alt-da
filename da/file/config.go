package file

import (
	"github.com/spf13/viper"
	"os"
	"path"
)

const (
	PlasmaFilePath = "FILE_PATH"
)

type Config struct {
	Directory string
}

func DefaultFileStoreCfg() Config {
	homedir, _ := os.UserHomeDir()
	cfg := Config{
		Directory: path.Join(homedir, ".plasma-da/data/filestore"),
	}
	if plasmaFilePath := viper.GetString(PlasmaFilePath); plasmaFilePath != "" {
		cfg.Directory = plasmaFilePath
	}

	return cfg
}
