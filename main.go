package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.SetConfigName("config")
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Retrieve the slice of script IDs
	scriptIdsY := viper.Get("scriptIds")

	// Assert that the returned value is a slice
	ids, ok := scriptIdsY.([]interface{})
	if !ok {
		fmt.Println("scriptIds is not a slice")
		return
	}

	// Loop through the slice and print each script ID
	for i, scriptId := range ids {
		process(scriptId.(string), i)
		fmt.Println("Finish: ", i+1, scriptId.(string))
		fmt.Println()
	}
}
