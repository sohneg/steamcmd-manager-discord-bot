package steamcmd_sm

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func CreateDirectory() {
	path := "steam"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Created Folder")
		
		copy("steamcmd/steamcmd.exe", "steam/steamcmd.exe")

		cmd := exec.Command("steam/steamcmd.exe")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Start()
		cmd.Wait()
		
		time.Sleep(30 * time.Second)
		
		cmd = exec.Command("login anonymous")
		cmd.Start()
		time.Sleep(30 * time.Second)
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}