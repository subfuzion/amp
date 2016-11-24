package main

import (
	"fmt"
	"os"
	"strconv"
)

//AgentConfig configuration parameters
type ClientConfig struct {
	serverAddr string
	serverPort string
}

//update conf instance with default value and environment variables
func (c *ClientConfig) init(version string, build string) {
	c.setDefault()
	c.loadConfigUsingEnvVariable()
	//cfg.displayConfig(version, build)
}

//Set default value of configuration
func (c *ClientConfig) setDefault() {
	c.serverAddr = "127.0.0.1"
	c.serverPort = "31315"
}

//Update config with env variables
func (c *ClientConfig) loadConfigUsingEnvVariable() {
	c.serverAddr = c.getStringParameter("SERVER_ADDRESS", c.serverAddr)
	c.serverPort = c.getStringParameter("SERVER_PORT", c.serverPort)
}

//display amp-pilot configuration
func (c *ClientConfig) displayConfig(version string, build string) {
	fmt.Printf("agrid version: %v build: %s\n", version, build)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("Configuration:")
	fmt.Printf("agrid address: %s:%d\n", c.serverAddr, c.serverPort)
}

//return env variable value, if empty return default value
func (c *ClientConfig) getStringParameter(envVariableName string, def string) string {
	value := os.Getenv(envVariableName)
	if value == "" {
		return def
	}
	return value
}

//return env variable value convert to int, if empty return default value
func (c *ClientConfig) getIntParameter(envVariableName string, def int) int {
	value := os.Getenv(envVariableName)
	if value != "" {
		ivalue, err := strconv.Atoi(value)
		if err != nil {
			return def
		}
		return ivalue
	}
	return def
}
