package steamcmd_sm

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func CreateNewSession() {
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("Error while creating session ", err)
		return
	}

	dg.AddHandler(readMessages)

	dg.Identify.Intents = discordgo.IntentGuildMembers | discordgo.IntentGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("Error while opening connection ", err)
	}

	fmt.Println("Server Manager is now running. Do not close this window. Press CTRL-C to stop the bot.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
