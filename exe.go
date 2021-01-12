package exe

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/gen2brain/beeep"
	"github.com/rodolfoag/gow32"
)

func cmd(command string, args ...string) (string, error) {
	fmt.Print("Executando '" + command + "' -> ")
	fmt.Println(args)

	cmdInstance := exec.Command(command, args...)
	cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmdOutput, err := cmdInstance.Output()

	if err != nil {
		return "", err
	}

	return string(cmdOutput), nil
}

func execjava(jar string) {
	output, err := cmd("java", "-jar", jar)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}

func fileexists(path string) bool {
	var status bool
	status = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		status = false
	}
	return status
}

func notify(title string, message string) {
	err := beeep.Notify(title, message, "assets/information.png")
	if err != nil {
		panic(err)
	}
}

// JavaIsInstalled verifica se o java está instalado
func JavaIsInstalled() bool {
	var status bool
	status = true

	output, err := cmd("java", "-version")

	if err != nil {
		status = false
	}

	fmt.Println(output)

	return status
}

// Open Abre o arquivo jar usando o ID informado.
func Open(id string, title string, jar string) {
	if !JavaIsInstalled() {
		notify(title, "Java não está instalado, verifique.")
	}

	if !fileexists(jar) {
		notify(title, "Arquivo .jar não encontrado, verifique.")
	}

	_, err := gow32.CreateMutex(id)
	if err != nil {
		notify(title, "App já iniciado, verifique.")
	} else {
		execjava(jar)
	}
}

func main() {
	Open("app", "app", "app.jar")
}
