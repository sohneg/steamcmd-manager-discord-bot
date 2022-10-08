package steamcmd_sm

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
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

		cmd := exec.Command("cmd", "/C", `cd `+INSTALLATION_PATH+`\steam && steamcmd +login anonymous +app_update 1829350 +quit`)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()
		// fmt.Println("\nServer is now installed! Plese configure your server settings and password.")
		// cmd2 := exec.Command("cmd", "/C", `cd %s\steam\steamapps\common\VRisingDedicatedServer && start notepad "start_server_example.bat"`, INSTALLATION_PATH)
		// cmd3 := exec.Command("cmd", "/C", `cd %s\steam\steamapps\common\VRisingDedicatedServer\VRisingServer_Data\StreamingAssets\Settings && start notepad "ServerHostSettings.json"`, INSTALLATION_PATH)
		// cmd2.Run()
		// cmd3.Run()
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
