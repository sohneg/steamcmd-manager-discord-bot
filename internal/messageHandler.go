package steamcmd_sm

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var isRunning = false

func readMessages(session *discordgo.Session, msg *discordgo.MessageCreate){

	if msg.Author.ID == session.State.User.ID {
		return
	}

	lowerContent := strings.ToLower(msg.Content)


	switch lowerContent{
	case HELP: 
		// session.ChannelMessageSend(msg.ChannelID, "To install steamcmd on your pc follow this documentation: https://developer.valvesoftware.com/wiki/SteamCMD")
		session.ChannelMessageSend(msg.ChannelID, COMMANDS);

	case START_SERVER:
		m, _ := session.ChannelMessageSend(msg.ChannelID, "Server is starting, please wait")
		progressText(session, msg, m,"Server is starting, please wait \n[", 1,)
		session.ChannelMessageSend(msg.ChannelID, "Server is now running! Join this IP: " + SERVER_IP_ADDRESS)
		isRunning = true

	case STOP_SERVER:
		isRunning = false

	case INSTALL_STEAMCMD:
		session.ChannelMessageSend(msg.ChannelID, "Installing SteamCMD in root directory...")
		CreateDirectory()

	case CHECK:
		if isRunning {
			session.ChannelMessageSend(msg.ChannelID, "Server is already running. Join this IP: " + SERVER_IP_ADDRESS)
		} else {
			session.ChannelMessageSend(msg.ChannelID, "Server is not running. Start with !start")
		}

	case VERSION:
		session.ChannelMessageSend(msg.ChannelID, BOT_INFO)
	}
	
	if strings.Contains(lowerContent, SHOW_SERVER_LIST){
		SteamcmdServerIds(session, msg, lowerContent)

	}
}



func progressText(session *discordgo.Session, channel *discordgo.MessageCreate , m *discordgo.Message, text string, wait time.Duration){
	
	loading := make(map[int]string)
	loading[1]= "#"
	loading[2]= "##"
	loading[3]= "###"
	loading[4]= "####"
	loading[5]= "#####"
	loading[6]= "######"
	loading[7]= "#######"
	loading[8]= "########"
	loading[9]= "#########"
	loading[10]= "##########"
	
	for i := 0; i <= 10; i++{
		time.Sleep(wait * time.Second)
		session.ChannelMessageEdit(channel.ChannelID, m.ID, text + loading[i] + "]")
	}
}