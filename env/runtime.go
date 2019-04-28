package env

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const RUNTIME_CONFIG = "RUNTIME_CONFIG"

type Runtime struct {
	ConfigPath string
}

var Env *Runtime

// 从环境变量对应的目录下加载需要的配置文件
func (r *Runtime) LoadConfig(file string, cfg interface{}) error {
	vip := viper.New()
	vip.SetConfigType("yaml")
	vip.SetConfigFile(filepath.Join(r.ConfigPath, file))
	err := vip.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return vip.Unmarshal(cfg)
}

func init() {
	Env = new(Runtime)
	absPath := os.Getenv(RUNTIME_CONFIG)
	if !filepath.IsAbs(absPath) {
		absPath, err := filepath.Abs(absPath)
		if err != nil {
			panic(absPath + " err:" + err.Error())
		}
		Env.ConfigPath = absPath
	} else {
		Env.ConfigPath = absPath
	}

}
