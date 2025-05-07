package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"crabot/crabtalk"
	"crabot/dice"
	"crabot/lastupload"
	"crabot/timestamp"
	"internal/dicecmd"
	"internal/env"

	"github.com/bwmarrin/discordgo"
)

// Bot flags
var (
	BotToken = flag.String("token", "", "Bot access token")
	GuildID  = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")

	AddCommands    = flag.Bool("addcmd", false, "Add all commands on start up (only run on first startup)")
	RemoveCommands = flag.Bool("rmcmd", false, "Remove all commands after shutdowning or not")
)

// Session
var session *discordgo.Session

// Start session
func getCredentials() {
	var err error

	// Check if token is entered.
	if *BotToken == "" {
		// If token isn't entered, get token from env file.
		*BotToken, err = env.GetEnvValue("TOKEN", "")
		if err != nil {
			log.Fatalln("Error occurred while getting the token", err)
		}
	}

	session, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalln("Invaild bot parameters:", err)
	}
}

var (
	// Not needed for only basic command
	/*
		integerOptionMinValue = 1.0
		dmPermission = false
		defaultMemberPermissions int64 = discordgo.PermissionManageServer
	*/

	// Bot commands
	commands = []*discordgo.ApplicationCommand{
		{
			// Basic command for testing/learning
			Name:        "basic-command",
			Description: "A basic test command",
		},
		{
			// Crab talk command
			Name:        "crabtalk",
			Description: "Translate your message to crab!",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "text-to-translate",
					Description: "Message to translate",
					Required:    true,
				},
			},
		},
		{
			// Coin flipping command
			Name:        "coinflip",
			Description: "Flip a coin and optionally call it in the air",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "call",
					Description: "Call the coin in the air",
					Required:    false,
				},
			},
		},
		{
			// Dice rolling command
			Name:        "dice",
			Description: "Roll a given number of dice types",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "type",
					Description: "Type of dice to roll",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "amount",
					Description: "Amount of dice to roll",
					Required:    false,
				},
			},
		},
		{
			// Command to get last upload info
			Name:        "lastupload",
			Description: "Get info about the last time someone uploaded on Youtube",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "Username of the youtuber",
					Required:    true,
				},
			},
		},
		{
			// Help command
			Name:        "help",
			Description: "Get help on usage of commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "command",
					Description: "Command you want help with",
					Required:    false,
				},
			},
		},
		{
			// Command to generate timestamps
			Name:        "timestamp",
			Description: "Generate a timestamp of given time",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "year",
					Description: "Year of timestamp",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "month",
					Description: "Month of timestamp",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "day",
					Description: "Day of timestamp",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "hour",
					Description: "Hour of timestamp",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "minute",
					Description: "Minute of timestamp",
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "timezone",
					Description: "Timezone in (tz timezone format)",
				},
			},
		},
	}

	// Handles what to do when a command is ran
	commandHandlers = map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate){
		// basic-command handler
		"basic-command": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
			session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "This is a basic test command.",
				},
			})
		},
		"coinflip": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
			options := interaction.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			message := "Result: "

			if option, ok := optionMap["call"]; ok {
				option := strings.ToLower(option.StringValue())

				if option != "heads" && option != "tails" {
					log.Panicln("[WARNING]: Unsupported coin side (use 'heads' or 'tails')")
				}

				result, called, err := dice.CoinCall(option)

				if err != nil {
					log.Printf("[ERROR]: CoinCall(%s) failed\nError: %!(error)\n", option, err)
				}

				message += result
				if called {
					message += "\n Coin was called correctly."
				} else {
					message += "\n Coin was not called correctly."
				}
			} else {
				result, _ := dice.CoinFlip()

				message += result
			}

			session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: message,
				},
			})
		},
		"crabtalk": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
			// Access input
			options := interaction.ApplicationCommandData().Options

			// Convert slice (options) into map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			margs := make([]interface{}, 0, len(options))
			msgformat := "\nMessage in Crab: \n"

			if option, ok := optionMap["text-to-translate"]; ok {
				margs = append(margs, option.StringValue())
				margsString := fmt.Sprint(margs)

				translatedText, _ := crabtalk.Get(margsString)
				msgformat += fmt.Sprintf("> %v\n", translatedText)
			}

			session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msgformat,
				},
			})
		},
		"dice": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
			options := interaction.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			diceType := "D6"
			diceAmount := 1

			if option, ok := optionMap["type"]; ok {
				diceType = strings.ToUpper(option.StringValue())
			}

			if option, ok := optionMap["amount"]; ok {
				diceAmount = int(option.IntValue())
			}

			result := dicecmd.DetermineDiceRoll(diceType, diceAmount)

			message := fmt.Sprintf("Rolled %d %s\nResults: %d\n", diceAmount, diceType, result)

			session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: message,
				},
			})
		},
		// basic-command handler
		"help": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
			options := interaction.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			var command string
			if option, ok := optionMap["command"]; ok {
				command = option.StringValue()
			}

			var msgformat string
			switch command {
			case "basic-command":
				msgformat += "> `/basic-command`\n"
				msgformat += "> Runs a basic command to confirm everything is working\n"
			case "crabtalk":
				msgformat += "> `/crabtalk <text-to-translate>`\n"
				msgformat += "> Translates *text-to-translate* to crab\n"
				msgformat += "> `text-to-translate`: Text to translate into crab (string)\n"
			case "coinflip":
				msgformat += "> `/coinflip [guess]`\n"
				msgformat += "> Flips a coin that you can optionally *guess* before it lands.\n"
				msgformat += "> `guess`: heads or tails\n"
			case "dice":
				msgformat += "> `/dice [type] [times]`\n"
				msgformat += "> Rolls the given *type* (defaults to D6) of dice given amount of *times* (defaults to 1)\n"
				msgformat += "> `types`: D2, D4, D6, D8, D10, D12, D20, D100\n"
			case "lastupload":
				msgformat += "> `/lastupload <username>`\n"
				msgformat += "> Fetches the last upload date of a YouTube *username*.\n"
				msgformat += "> `username`: YouTubers's name (string)\n"
			case "help":
				msgformat += "> `/help [command]`\n"
				msgformat += "> Shows help message for *command* or overview of all commands\n"
				msgformat += "> `command`: Commands you want help with\n"
			case "timestamp":
				msgformat += "> `/timestamp [year] [month] [day] [hour] [minute] [timezone]`\n"
				msgformat += "> Generates a timestamp based on given time.\n"
				msgformat += "> Any options omitted default to current time besides timezone which defaults to UTC\n"
			default:
				msgformat += "> `basic-command`: Run a basic test command\n"
				msgformat += "> `crabtalk`: Translates text to crab\n"
				msgformat += "> `coinflip`: Flips a coin optionally allowing you to guess it in the air\n"
				msgformat += "> `dice`: Rolls given type of dice an amount of times\n"
				msgformat += "> `lastupload`: Fetches the last upload date of a YouTuber\n"
				msgformat += "> `help`: This command; optionally pass a command to get more help for it\n"
				msgformat += "> `timestamp`: Generate a Discord timestamp\n"
			}

			session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msgformat,
				},
			})
		},
		"lastupload": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
			options := interaction.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			var username string
			if option, ok := optionMap["username"]; ok {
				username = option.StringValue()
			}

			fetechedCorrectly := true

			YOUTUBE_KEY, err := env.GetEnvValue("YOUTUBE_KEY", "./.env")
			if err != nil {
				fetechedCorrectly = false
				log.Printf("[ERROR]: %v\n", err)
			}

			lastUploadData, err := lastupload.Get(username, YOUTUBE_KEY)
			if err != nil {
				fetechedCorrectly = false
				log.Printf("[ERROR]: %v\n", err)
			}

			uploadDate, err := lastupload.GetDate(lastUploadData)
			if err != nil {
				fetechedCorrectly = false
				log.Printf("[ERROR]: %v\n", err)
			}

			// Get Unix timestamp for formating with Discord's timestamps
			uploadDateUnix := uploadDate.Unix()

			// Get wanted Discord timestamps
			uploadDateRelativeTimestamp := timestamp.UnixRelativeTime(int(uploadDateUnix))
			uploadDateFullTimestamp := timestamp.UnixLongFull(int(uploadDateUnix))

			msgformat := fmt.Sprintf("YouTuber **%v** last uploaded %v ago on %v.", username, uploadDateRelativeTimestamp, uploadDateFullTimestamp)

			// Confirm infomation was fetched correctly.
			if !fetechedCorrectly {
				msgformat = `[ERROR]: Issue fetching data; see log for more info`
			}

			session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msgformat,
				},
			})
		},
		"timestamp": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
			options := interaction.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// Get current time
			now := time.Now()

			// Setup defaults for timestamp options
			year := int(now.Year())
			month := int(now.Month())
			day := int(now.Day())
			hour := int(now.Hour())
			minute := int(now.Minute())

			timezone := "UTC"

			// Get user entered options
			if option, ok := optionMap["year"]; ok {
				year = int(option.IntValue())
			}
			if option, ok := optionMap["month"]; ok {
				month = int(option.IntValue())
			}
			if option, ok := optionMap["day"]; ok {
				day = int(option.IntValue())
			}
			if option, ok := optionMap["hour"]; ok {
				hour = int(option.IntValue())
			}
			if option, ok := optionMap["minute"]; ok {
				minute = int(option.IntValue())
			}
			if option, ok := optionMap["timezone"]; ok {
				timezone = option.StringValue()
			}

			// Message formatting
			var msgformat string

			shortDate, err := timestamp.ShortDate(year, month, day, hour, minute, 00, timezone)
			if err != nil {
				log.Panicf("[ERROR]: %v\n", err)
			}
			msgformat += fmt.Sprintf("> `%v` -> %v\n", shortDate, shortDate)

			longDate, err := timestamp.LongDate(year, month, day, hour, minute, 00, timezone)
			if err != nil {
				log.Panicf("[ERROR]: %v\n", err)
			}
			msgformat += fmt.Sprintf("> `%v` -> %v\n", longDate, longDate)

			shortTime, err := timestamp.ShortTime(year, month, day, hour, minute, 00, timezone)
			if err != nil {
				log.Panicf("[ERROR]: %v\n", err)
			}
			msgformat += fmt.Sprintf("> `%v` -> %v\n", shortTime, shortTime)

			longTime, err := timestamp.LongTime(year, month, day, hour, minute, 00, timezone)
			if err != nil {
				log.Printf("[ERROR]: %v\n", err)
			}
			msgformat += fmt.Sprintf("> `%v` -> %v\n", longTime, longTime)

			shortFull, err := timestamp.ShortFull(year, month, day, hour, minute, 00, timezone)
			if err != nil {
				log.Printf("[ERROR]: %v\n", err)
			}
			msgformat += fmt.Sprintf("> `%v` -> %v\n", shortFull, shortFull)

			longFull, err := timestamp.LongFull(year, month, day, hour, minute, 00, timezone)
			if err != nil {
				log.Printf("[ERROR]: %v\n", err)
			}
			msgformat += fmt.Sprintf("> `%v` -> %v\n", longFull, longFull)

			relativeTime, err := timestamp.RelativeTime(year, month, day, hour, minute, 00, timezone)
			if err != nil {
				log.Printf("[ERROR]: %v\n", err)
			}
			msgformat += fmt.Sprintf("> `%v` -> %v\n", relativeTime, relativeTime)

			session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msgformat,
				},
			})
		},
	}
)

func init() {
	// Parse flags
	flag.Parse()
}

func main() {
	// Get credentials
	getCredentials()

	// Add handlers
	session.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		// Confirm handlers are added correctly
		if h, ok := commandHandlers[interaction.ApplicationCommandData().Name]; ok {
			h(session, interaction)
		}
	})

	// Logged in notification
	session.AddHandler(func(session *discordgo.Session, ready *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v\n", session.State.User.Username, session.State.User.Discriminator)
	})

	// Throw error if session can't be opened
	err := session.Open()
	if err != nil {
		log.Fatalln("Can't open the session:", err)
	}

	// Create commands var
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	// Check if we want to add commands
	if *AddCommands {
		log.Println("Adding commands...")

		// Loop for adding each interaction/command
		for i, v := range commands {
			cmd, err := session.ApplicationCommandCreate(session.State.User.ID, *GuildID, v)
			if err != nil {
				log.Panicf("Can't create '%v' command: %v\n", v.Name, err)
			}
			registeredCommands[i] = cmd
		}
	}

	defer session.Close()

	// Handles ending bot session
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Bot is running... (Press CTRL-C to exit)")
	<-stop

	// Check if we want to remove commands
	if *RemoveCommands {
		log.Println("Removing commands...")

		// Loop to delete commands
		for _, v := range registeredCommands {
			err := session.ApplicationCommandDelete(session.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Can't delete '%v' command: %v\n", v.Name, err)
			}
		}
	}

	// Shut down log
	log.Println("Gracefully shutting down.")
}
