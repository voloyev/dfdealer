package commands

import (
	"os"
	"path"

	"log"

	"github.com/urfave/cli/v2"
)

func Check(ctx *cli.Context) error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Cant read home directory")
		return err
	}
	symlinkPath := path.Join(dirname, ".config", ctx.Args().Get(0))
	log.Printf("Path: %q", symlinkPath)

	lpath, err := os.Readlink(symlinkPath)
	if err != nil {
		log.Fatal("Cant find symbolic link to the config")
		return err
	}

	log.Printf("Checking %q", lpath)
	return nil
}
