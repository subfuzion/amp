package agentcore

import ()

var functionMap map[string]interface{}

func initFunctionMap() {
	functionMap = make(map[string]interface{})
	functionMap["ping"] = ping
	functionMap["info"] = getInfo
	functionMap["setLogLevel"] = setLogLevel
}

func ping(g *SwarmAgent, name string) {
	logf.debug("execute ping from: %s\n", name)
}

func getInfo(g *SwarmAgent) int {
	return 0
}

func setLogLevel(g *SwarmAgent, level string) {
	logf.setLevel(level)
	logf.printf("Set log level: " + logf.levelString())
}
