package steamcmd_sm

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var TOKEN string

func SetupDiscordToken() {
	path := `settings\discord_token.txt`
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {

		newFile, err := os.Create(path)
		if err != nil {
			fmt.Println("Could not create discord_token.txt. Error: ", err)
		}
		defer newFile.Close()

		fmt.Print("Thank you for using SteamCMD Dedicated Discord Bot. Please enter your bot token, which you get from https://discord.com/developers/applications: \n")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		TOKEN = input.Text()

		_, writeErr := newFile.WriteString(input.Text())
		if writeErr != nil {
			fmt.Println("Could not write to discord_token.txt. Error: ", err)
		}

	} else {
		content, _ := ioutil.ReadFile(path)
		myString := string(content)
		TOKEN = string(myString)
	}

	time.Sleep(1 * time.Second)
}

//OTc3NDg3ODU4ODYyNTQyOTM4.G4NQG8.cVdADO8aib9VYYnGl-k715boax1hNC_3lncJJw
