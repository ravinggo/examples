package env

import (
	"os"
	"strings"

	env "github.com/ravinggo/game/common/base-env"
)

var setDefaultEnvErr = setDefaultEnv()

func init() {
	if setDefaultEnvErr != nil {
		panic(setDefaultEnvErr)
	}
}

func setDefaultEnv() error {
	for k, v := range map[string]string{
		env.ServerType: "login-server",
		env.ConfType:   "consul",
		env.ConfHosts:  "127.0.0.1:8500",
		env.ConfPath:   "/example-game",
		env.ConfFormat: "json",
		env.ServerId:   "0",
		env.PprofAddr:  ":6060",
	} {
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		if strings.TrimSpace(os.Getenv(k)) == "" {
			err := os.Setenv(k, v)
			if err != nil {
				return err
			}
		}
	}
	env.InitConfig()
	return nil
}
