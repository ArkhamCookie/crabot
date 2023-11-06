package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"crabot/crabtalk"
	"formatting/markdown"
	"internal/env"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

// Session
var session *discordgo.Session

// Parse flags
func init() { flag.Parse() }

// Start session
func init() {
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
			Name: "crabtalk",
			Description: "Translate your message to crab!",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type: discordgo.ApplicationCommandOptionString,
					Name: "text-to-translate",
					Description: "Message to translate",
					Required: true,
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
				msgformat += markdown.Blockquote(translatedText)
				msgformat += "\n"
			}

			session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msgformat,
				},
			})
		},
	}
)

// Add handlers
func init() {
	session.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		// Confirm handlers are added correctly
		if h, ok := commandHandlers[interaction.ApplicationCommandData().Name]; ok {
			h(session, interaction)
		}
	})
}

func main() {
	// Logged in notification
	session.AddHandler(func(session *discordgo.Session, ready *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v\n", session.State.User.Username, session.State.User.Discriminator)
	})

	// Throw error if session can't be opened
	err := session.Open()
	if err != nil {
		log.Fatalln("Can't open the session:", err)
	}

	// Create commands
	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	// Loop for adding each interaction
	for i, v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Can't create '%v' command: %v\n", v.Name, err)
		}
		registeredCommands[i] = cmd
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
	}

	// Loop to delete commands
	for _, v := range registeredCommands {
		err := session.ApplicationCommandDelete(session.State.User.ID, *GuildID, v.ID)
		if err != nil {
			log.Panicf("Can't delete '%v' command: %v\n", v.Name, err)
		}
	}

	// Shut down log
	log.Println("Gracefully shutting down.")
}
