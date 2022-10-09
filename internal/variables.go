package steamcmd_sm

// IP address will be searched auto.
var SERVER_IP_ADDRESS = CheckPublicIP() + ":9876"

// Usable commands
const START_SERVER = "!start"
const STOP_SERVER = "!stop_admin"
const HELP = "!help"
const CHECK = "!check"
const SHOW_SERVER_LIST = "!list"
const VERSION = "!version"
const INSTALL_STEAMCMD = "!install"

var IS_INSTALLED bool = false

// Discord Output !help
const COMMANDS = `You can use this commands to manage your steamcmd dedicated server on you localhost:

**!start**
Start the server.	

**!stop**
Stop the server.

**!check**
Check if the server is already running.

**!install**
Installs V Rising Dedicated Server. More comming.

**!list [game]**
Search for game servers. Only server to install right now, is V Rising. But there will be over 500 diffrent servers!

(All commands can be assigned to roles)`

// Bot Info. Please do not change
const BOT_INFO = `SteamCMD Server Manager Bot
Version: 1.0a
Author: Simon S.`
