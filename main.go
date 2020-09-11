package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("usage: ./god  <project name> <'os:arch:cgo'>")
	}
	instructions := strings.Split(os.Args[2], ",")
	for _, instruction := range instructions {
		build := strings.Split(instruction, ":")
		cmd := exec.Command("go", "build", "-o", os.Args[1]+"-"+build[0]+"-"+build[1])
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GOOS="+build[0])
		cmd.Env = append(cmd.Env, "GOARCH="+build[1])
		cmd.Env = append(cmd.Env, "CGO_ENABLED="+build[2])
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
	}
}
