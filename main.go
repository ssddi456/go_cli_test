package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

func execAndReadStderr(cmd *exec.Cmd) (string, error) {
	var errbuf bytes.Buffer
	cmd.Stderr = &errbuf

	err := cmd.Run()
	if err != nil {
		log.Print("exec command failed, cmd ", cmd.String(), " error ", err)
	}
	return errbuf.String(), err
}

func execAndReadStdout(cmd *exec.Cmd) (string, error) {
	stdout, err := cmd.Output()
	if err != nil {
		log.Print("exec command failed, cmd ", cmd.String(), " error ", err)
		// return "", err
	}

	return string(stdout), err
}

func downloadFlutterw() {

}

type checkRes struct {
	WhereChoco    string `json:"whereChoco"`
	ChocoVersion  string `json:"chocoVersion"`
	WherePython   string `json:"wherePython"`
	PythonVersion string `json:"pythonVersion"`
	SystemPath    string `json:"systemPath"`
}

// 检查环境变量
func checkEnv() checkRes {
	whereChoco, _ := execAndReadStdout(exec.Command("where", "choco"))
	chocoVersion, _ := execAndReadStdout(exec.Command("choco", "-v"))
	// log.Printf("whereChoco %s", whereChoco)

	wherePython, _ := execAndReadStdout(exec.Command("where", "python"))
	// log.Printf("wherePython %s", wherePython)
	wherePython = strings.TrimSpace(wherePython)

	pythonVersion, _ := execAndReadStderr(exec.Command("python", "--version"))
	pythonVersion = strings.SplitN(strings.TrimSpace(pythonVersion), " ", 2)[1]

	// log.Printf("pythonVersion %s", pythonVersion)
	systemPath := os.Getenv("PATH")

	res := checkRes{
		whereChoco,
		chocoVersion,
		wherePython,
		pythonVersion,
		systemPath,
	}
	return res
}

func doctor() {
	res := checkEnv()
	jsonBytes, _ := json.MarshalIndent(res, "", " ")

	fmt.Println("please post this doct info to maintainers")
	fmt.Println(string(jsonBytes))
}

func checkEnvMac() {

	whereChoco, _ := execAndReadStdout(exec.Command("which", "choco"))
	log.Printf("whereChoco %s", whereChoco)

	wherePython, _ := execAndReadStdout(exec.Command("which", "python"))
	log.Printf("wherePython %s", wherePython)

	whereFlutter, _ := execAndReadStdout(exec.Command("which", "flutter"))
	log.Printf("whereflutter %s", whereFlutter)

	whereBrew, _ := execAndReadStdout(exec.Command("which", "brew"))
	log.Printf("whereBrew %s", whereBrew)
}

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")

			if runtime.GOOS == "windows" {
				jsonBytes, _ := json.MarshalIndent(checkEnv(), "", " ")

				fmt.Println(string(jsonBytes))
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
