# telegram-torrent-bot

A very basic bot for Telegram that connects to qBittorent and adds torrent links.

## Usage

The most basic usage is just sending `.torrent` files to the bot.

### Commands

`/add` - send torrent links to the bot.

`/list` - list all current torrents with their names and progress.
 

### Flags

`key` - Telegram api key, which you can get from the BotFather. All the required info can be found at https://core.telegram.org/bots.

`user` - Telegram user id. If this key is provided, the bot will only accept messages from that user. If the key is not proved, the bot will **accept messages from any user**.

`qbl` - qBittorent web ui username.

`qbp` - qBittorent web ui password.

`qbu` - qBittorent web ui base url, defaults to `http://localhost:8080`.
