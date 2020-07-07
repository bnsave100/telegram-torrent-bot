# telegram-torrent-bot

A very basic bot for telegram that connets to qBittorent and adds torrent links.

## Usage

### Flags

`key` - Telegram api key, which you can get from the BotFather. All the required info can be found at https://core.telegram.org/bots.

`user` - Telegram user id. If this key is provided, the bot will only accept messages from that user. If the key is not proved, the bot will **accept messages from any user**.

`qbl` - qBittorent web ui username.

`qbp` - qBittorent web ui password.

`qbu` - qBittorent web ui base url, defaults to `http://localhost:8080`.
