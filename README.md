# telegram_bot

Before starting, you need to create a SQLite database.

## Config env

| Parameter            | Description         | Default           |
|:---------------------|---------------------|-------------------|
| LOG_LEVEL            | logging level       | debug             |
| DB_PATH              | path to SQLite      | ./telegram_bot.db |
| AUTO_MIGRATE         | automatic migration | true              |
| MY_AWESOME_BOT_TOKEN | telegram bot token  |                   |


## Run

    go run main.go