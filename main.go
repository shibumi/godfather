package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("usage: ./god 'name:os:arch:cgo'")
	}
	instructions := strings.Split(os.Args[1], ",")
	for _, instruction := range instructions {
		build := strings.Split(instruction, ":")
		cmd := exec.Command("go", "build", "-o", build[0]+"-"+build[1]+"-"+build[2])
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GOOS="+build[1])
		cmd.Env = append(cmd.Env, "GOARCH="+build[2])
		cmd.Env = append(cmd.Env, "CGO_ENABLED="+build[3])
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
	}
}
