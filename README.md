# go-tg-pinger

A go telegram pinger sends a customized message or a ping to a specific telegram chat

The Bot API and the ChatID of the user can be injected on build. They can also be overridden by environment variable on run

You can pipe an output of a command into it through stdin or run it simply after a command is done

```bash
echo "Hello" | ./pinger
# Output: Host: peusdo-hostname, Msg: hello
echo "hello" && ./pinger
# Output: Host: peusdo-hostname, Msg: Ping
```
By default, it sends a "Ping" message when the program is executed if no stdin is found

It will also add the hostname of the machine before the message if the hostname is set

## Override BotAPI and ChatID on run
```bash
BOT_API=<BOT-API>  CHAT_ID=<ChatID> ./pinger
```

##  Inject API Key and ChatID on compilation
```bash
go build -ldflags "-X main.DefaultBotAPI=<Bot-API-Key> -X main.DefaultChatID=<ChatID>" -o pinger
```

##  Or build without sensitive data
```bash
go build -o pinger
```

#### TODOs:
- [ ] Add default helper message when there is no chatID or BotAPIKey
- [ ] Use flag to parse customized message that could override default Ping
- [ ] Add makefile to cross compile for all platforms 
