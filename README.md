# discord-pipe-logger

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Hexiro/discord-pipe-logger)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/2af2ae9526ae4b0d9a81635db8f983ab)](https://www.codacy.com/gh/Hexiro/discord-pipe-logger/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Hexiro/discord-pipe-logger&amp;utm_campaign=Badge_Grade)
![license MIT](https://img.shields.io/github/license/Hexiro/discord-pipe-logger)

Feed pipe input into a Discord server via webhook.<br>

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
