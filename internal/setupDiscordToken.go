package steamcmd_sm

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var TOKEN string
var INSTALLATION_PATH string

func SetupDiscordToken() {
	settingsDir := "settings"
	err := os.MkdirAll(settingsDir, os.ModePerm)
	if err != nil {
		fmt.Println("Could not create settings directory. Error:", err)
		return
	}

	path := filepath.Join(settingsDir, "discord_token.txt")
	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		newFile, err := os.Create(path)
		if err != nil {
			fmt.Println("Could not create discord_token.txt. Error:", err)
			return
		}
		defer newFile.Close()

		fmt.Print("Thank you for using SteamCMD Dedicated Discord Bot. Please enter your bot token, which you get from https://discord.com/developers/applications: \n")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		TOKEN = input.Text()

		_, writeErr := newFile.WriteString(input.Text())
		if writeErr != nil {
			fmt.Println("Could not write to discord_token.txt. Error:", err)
		}

	} else {
		content, _ := ioutil.ReadFile(path)
		myString := string(content)
		TOKEN = myString
	}

	setupInstallationPath()
	time.Sleep(1 * time.Second)
}

func setupInstallationPath() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	var path = "settings/installation_path.txt"
	INSTALLATION_PATH = exPath

	fmt.Printf("Bot installed in %s\n", INSTALLATION_PATH)

	newFile, err := os.Create(path)
	if err != nil {
		fmt.Println("Could not create installation_path.txt. Error: ", err)
	}
	defer newFile.Close()

	_, writeErr := newFile.WriteString(exPath)
	if writeErr != nil {
		fmt.Println("Could not write to installation_path.txt. Error: ", err)
	}
}
