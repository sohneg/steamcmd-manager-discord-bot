package steamcmd_sm

import (
	"os"
	"os/exec"
)

func startServer() {
	cmd := exec.Command(`steam\steamapps\common\VRisingDedicatedServer\VRisingServer.exe`, `steam\steamapps\common\VRisingDedicatedServer\start_server_example.bat`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
}
