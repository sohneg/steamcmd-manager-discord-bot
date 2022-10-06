package steamcmd_sm

var SERVER_IP_ADDRESS = CheckPublicIP() + ":9876"

const TOKEN = "OTc3NDg3ODU4ODYyNTQyOTM4.GOqRsS.CbR5141yjbSqX3IoDM6YzPHxBQ4OBa3-z0F7ZI"

const START_SERVER = "!start"
const STOP_SERVER = "!stop_admin"
const HELP = "!help"
const CHECK = "!check"
const SHOW_SERVER_LIST = "!list"
const VERSION = "!version"

const INSTALL_STEAMCMD = "!install"

const COMMANDS = `You can use this commands to manage your steamcmd dedicated server on you localhost:

**!start**
Start the server.	

**!stop**
Stop the server.

**!check**
Check if the server is already running.

**!list [game]**
Search if server exists for game.

(All commands can be assigned to roles)`

const BOT_INFO = `SteamCMD Server Manager Bot
Version: 1.0a
Author: Simon S.`
