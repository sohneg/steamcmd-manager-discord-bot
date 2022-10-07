package steamcmd_sm

import (
	"os"
	"os/exec"
)

func startServer() {
	cmd := exec.Command(`steam\steamapps\common\VRisingDedicatedServer\VRisingServer.exe`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()


	cmd.Run()
}
