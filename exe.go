package exe

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"

	"github.com/gen2brain/beeep"
	"github.com/rodolfoag/gow32"
)

func execjava(jar string) {
	cmdInstance := exec.Command("java", "-jar", jar)
	cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmdOutput, err := cmdInstance.Output()

	fmt.Println(cmdOutput)

	if err != nil {
		log.Fatal(err)
	}
}

func open(id string, title string, jar string) {
	_, err := gow32.CreateMutex(id)
	if err != nil {
		err := beeep.Notify(title, "App já iniciado, verifique.", "assets/information.png")
		if err != nil {
			panic(err)
		}
	} else {
		execjava(jar)
	}
}
