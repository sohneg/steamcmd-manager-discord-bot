package main

import (
	steamcmd_sm "github.com/sohneg/steamcmd-manager-discord-bot/internal"
)

func main() {
	steamcmd_sm.SetupDiscordToken()
	defer steamcmd_sm.CreateNewSession()
}
