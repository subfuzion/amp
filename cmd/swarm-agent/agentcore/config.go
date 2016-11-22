package agentcore

import (
	"fmt"
	"os"
	"strconv"
)

//AgentConfig configuration parameters
type AgentConfig struct {
	dockerEngine string
	apiPort      string
	agentID      string
	serverPort   string
	serverAddr   string
}

var conf AgentConfig

//update conf instance with default value and environment variables
func (c *AgentConfig) init(version string, build string) {
	c.setDefault()
	c.loadConfigUsingEnvVariable()
	c.displayConfig(version, build)
}

//Set default value of configuration
func (c *AgentConfig) setDefault() {
	c.dockerEngine = "unix:///var/run/docker.sock"
	c.apiPort = "3000"
	c.agentID = os.Getenv("HOSTNAME")
	c.serverPort = "3010"
}

//Update config with env variables
func (c *AgentConfig) loadConfigUsingEnvVariable() {
	c.dockerEngine = c.getStringParameter("DOCKER", c.dockerEngine)
	c.apiPort = c.getStringParameter("API_PORT", c.apiPort)
	c.serverPort = c.getStringParameter("SERVER_PORT", c.serverPort)
	c.serverAddr = c.getStringParameter("SERVER", c.serverAddr)
}

//display amp-pilot configuration
func (c *AgentConfig) displayConfig(version string, build string) {
	fmt.Printf("amp-swarm version: %v build: %s\n", version, build)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("Configuration:")
	fmt.Printf("Docker-engine: %s\n", c.dockerEngine)
	fmt.Printf("AgentId: %s\n", c.agentID)
	fmt.Printf("ServerPort: %s\n", c.serverPort)
	fmt.Println("----------------------------------------------------------------------------")
}

//return env variable value, if empty return default value
func (c *AgentConfig) getStringParameter(envVariableName string, def string) string {
	value := os.Getenv(envVariableName)
	if value == "" {
		return def
	}
	return value
}

//return env variable value convert to int, if empty return default value
func (c *AgentConfig) getIntParameter(envVariableName string, def int) int {
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
