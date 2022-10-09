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

// Discord Output !help
const COMMANDS = `You can use this commands to manage your steamcmd dedicated server on you localhost:

**!start**
Start the server.	

**!stop_admin**
Stop the server. Not functional yet! Server has to be stopped manually + in the Discord Server !stop_admin has to be called. If not, the server can't be started again.

**!check**
Check if the server is already running.

**!install**
Installs V Rising Dedicated Server. More servers comming...

**!list [game]**
Search for game servers. Only server to install right now, is V Rising. But there will be over 500 diffrent servers!

(All commands can be assigned to roles)`

// Bot Info. Please do not change
const BOT_INFO = `SteamCMD Server Manager Bot
Version: 1.1a
Author: Simon S.
Date: 09.10.2022`
