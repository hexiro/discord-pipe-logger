# discord-pipe-logger
Feed pipe input into a Discord server via webhook.<br>


![GitHub](https://img.shields.io/github/license/Hexiro/discord-pipe-logger)

## Installation
```console
go get github.com/Hexiro/discord-pipe-logger
```

## Usage
Only ID and Token are required.<br/>
All types of Discord webhook URLs should work
```console
$ {command} | discord-pipe-logger "ID/Token"
$ {command} | discord-pipe-logger "https://discord.com/api/webhooks/ID/Token"
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
