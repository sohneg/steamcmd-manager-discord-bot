package steamcmd_sm

// Enter your token
const TOKEN = ""

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

// Discord Output !help
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



// Bot Info. Please do not change
const BOT_INFO = `SteamCMD Server Manager Bot
Version: 1.0a
Author: Simon S.`
