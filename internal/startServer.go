package steamcmd_sm

import (
	"os"
	"os/exec"
)

func startServer() {
	cmd := exec.Command("cmd", "/C", `cd %s\steam\steamapps\common\VRisingDedicatedServer\ && start start_server_example.bat`, INSTALLATION_PATH)
	// cmd := exec.Command(`steam\steamapps\common\VRisingDedicatedServer\VRisingServer.exe`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
