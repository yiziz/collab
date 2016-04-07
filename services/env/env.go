package env

import (
	"os"

	"github.com/yiziz/collab/path"
	"github.com/yiziz/collab/services/yml"
)

// SetEnvVars sets the env variables from application.yml
func SetEnvVars(configPath ...string) {
	configMap := yml.ConfigYML(path.AppConfigFilename(configPath...))
	for k, v := range configMap {
		key := k.(string)
		val := v.(string)
		os.Setenv(key, val)
	}
}
