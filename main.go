package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/urfave/cli/v2"
)

func execAndReadStdout(cmd *exec.Cmd) string {

	stdout, err := cmd.Output()
	if err != nil {
		log.Print("exec command failed, cmd ", cmd.String(), " error ", err)
	}

	return string(stdout)
}

// 检查环境变量
func checkEnv() {

	whereChoco := execAndReadStdout(exec.Command("where", "choco"))
	log.Printf("whereChoco %s", whereChoco)

	wherePython := execAndReadStdout(exec.Command("where", "python"))
	log.Printf("wherePython %s", wherePython)

	whereFlutter := execAndReadStdout(exec.Command("where", "flutter"))
	log.Printf("whereflutter %s", whereFlutter)

	whereBrew := execAndReadStdout(exec.Command("where", "brew"))
	log.Printf("whereBrew %s", whereBrew)

}

func checkEnvMac() {

	whereChoco := execAndReadStdout(exec.Command("which", "choco"))
	log.Printf("whereChoco %s", whereChoco)

	wherePython := execAndReadStdout(exec.Command("which", "python"))
	log.Printf("wherePython %s", wherePython)

	whereFlutter := execAndReadStdout(exec.Command("which", "flutter"))
	log.Printf("whereflutter %s", whereFlutter)

	whereBrew := execAndReadStdout(exec.Command("which", "brew"))
	log.Printf("whereBrew %s", whereBrew)
}

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")

			if runtime.GOOS == "windows" {
				checkEnv()
			} else if runtime.GOOS == "darwin" {
				checkEnvMac()
			} else {
				log.Fatal("unsupport platform")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
