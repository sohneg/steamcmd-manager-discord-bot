package steamcmd_sm

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Server struct {
	AppID               int    `json:"appid"`
	Subscriptionlinux   string `json:"subscriptionlinux"`
	Linux               bool   `json:"linux"`
	Subscriptionwindows string `json:"subscriptionwindows"`
	Windows             bool   `json:"windows"`
	Name                string `json:"name"`
}

func readJSONFromUrl(url string) ([]Server, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var countryList []Server
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &countryList); err != nil {
		return nil, err
	}

	return countryList, nil
}

func SteamcmdServerIds(session *discordgo.Session, channel *discordgo.MessageCreate, search string) {
	url := "https://raw.githubusercontent.com/dgibbs64/SteamCMD-AppID-List-Servers/master/steamcmd_appid_servers.json"
	serverList, err := readJSONFromUrl(url)
	if err != nil {
		panic(err)
	}

	res1 := strings.Split(search, "!list ")
	a := strings.Join(res1[1:], " ")

	session.ChannelMessageSend(channel.ChannelID, "Searching for Game...")
	time.Sleep(2 * time.Second) // just for usage feedback

	gameFound := false
	for _, row := range serverList {
		if strings.Contains(strings.ToLower(row.Name), a) {
			session.ChannelMessageSend(channel.ChannelID, "AppID: "+strconv.Itoa(row.AppID)+": "+row.Name)
			gameFound = true
		}
	}

	if !gameFound {
		session.ChannelMessageSend(channel.ChannelID, "No existing Server or typo.")
	}
}
