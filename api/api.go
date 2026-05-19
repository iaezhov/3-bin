package api

import (
	"fmt"
	"hw/3/config"
)

func OutputConfig(configObject *config.Config) {
	fmt.Println(configObject.Key)
}
