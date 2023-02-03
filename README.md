# telegram_bot

This telegram bot is a functional add-on over http://github.com/go-telegram-bot-api/telegram-bot-api.
This service has a configured connection to the database and command menu settings.

Before starting, you need to create a SQLite database.

## Config env

| Parameter            | Description         | Default           |
|:---------------------|---------------------|-------------------|
| LOG_LEVEL            | logging level       | debug             |
| DB_PATH              | path to SQLite      | ./telegram_bot.db |
| AUTO_MIGRATE         | automatic migration | true              |
| MY_AWESOME_BOT_TOKEN | telegram bot token  |                   |
| TRACE_SQL_COMMANDS   | trace sql commands  | false             |
| BOT_DEBUG            | bot debag           | false             |


## Run

    go run main.go