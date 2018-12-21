package main

import (
	"fmt"
	"log"
	"os"

	"github.com/twistedogic/dcmd/pkg/docker"
	"github.com/twistedogic/dcmd/pkg/file"
	"github.com/urfave/cli"
)

func run(image string, args []string) error {
	var command string
	hasEntrypoint := docker.HasEntrypoint(image)
	if !hasEntrypoint {
		command, args = args[0], args[1:]
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	workspace := file.CreateWorkspaceName()
	workspaceArg := []string{"-w", workspace}
	volArgs := append(workspaceArg, "-v", fmt.Sprintf("%s:%s", pwd, workspace))
	// for i, a := range args {
	// 	if fp, err := file.GetFilePath(a); err == nil {
	// 		sfp := file.SyntheticPath(fp)
	// 		volMount := fmt.Sprintf("%s:%s", file.DirPath(fp), file.DirPath(sfp))
	// 		volArg := []string{"-v", volMount}
	// 		volArgs = append(volArgs, volArg...)
	// 		args[i] = sfp
	// 	}
	// }
	if !hasEntrypoint {
		args = append([]string{command}, args...)
	}
	cmd := docker.CreateCmd(image, volArgs, args)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	app := cli.NewApp()
	app.Name = "dcmd"
	app.Usage = "run docker container as command with auto file mount"
	app.Action = func(c *cli.Context) error {
		args := c.Args()
		if args.First() == "" {
			cli.ShowAppHelpAndExit(c, 0)
		}
		image, args := args.First(), args.Tail()
		return run(image, args)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
