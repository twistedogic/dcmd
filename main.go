package main

import (
	"fmt"
	"log"
	"os"

	"github.com/twistedogic/dcmd/pkg/docker"
	"github.com/twistedogic/dcmd/pkg/file"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("missing image")
	}
	var command string
	image, args := args[0], args[1:]
	hasEntrypoint := docker.HasEntrypoint(image)
	if !hasEntrypoint {
		command, args = args[0], args[1:]
	}
	volArgs := []string{}
	for i, a := range args {
		if fp, err := file.GetFilePath(a); err == nil {
			sfp := file.SyntheticPath(fp)
			volMount := fmt.Sprintf("%s:%s", file.DirPath(fp), file.DirPath(sfp))
			volArg := []string{"-v", volMount}
			volArgs = append(volArgs, volArg...)
			args[i] = sfp
		}
	}
	if !hasEntrypoint {
		args = append([]string{command}, args...)
	}
	cmd := docker.CreateCmd(image, volArgs, args)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(cmd.Args)
		log.Fatal(err)
	}
}
