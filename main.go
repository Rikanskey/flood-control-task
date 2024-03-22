package main

import "task/internal/runner"

const configDir = "./config/"

func main() {
	runner.Start(configDir)
}
