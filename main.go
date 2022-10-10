package main

import "fmt"

// import (
// 	"os"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// func main() {
// 	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// bot.Debug = true

// 	updateConfig := tgbotapi.NewUpdate(0)

// 	updateConfig.Timeout = 30

// 	updates := bot.GetUpdatesChan(updateConfig)

// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}

// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

// 		msg.ReplyToMessageID = update.Message.MessageID

// 		if _, err := bot.Send(msg); err != nil {
// 			panic(err)
// 		}
// 	}
// }

func main() {
	game := new(ttt_game)
	for Play(game) {
	}
}

func Play(game *ttt_game) bool {
	game.Init()
	game.printField()
	for !game.finished && !game.tie {
		game.UpdateState()
	}
	fmt.Println("Wanna play again?(y/n)")
	var answer string
	_, err := fmt.Scan(&answer)
	for err != nil {
		fmt.Println("Wrong answer=/")
		_, err = fmt.Scan(&answer)
	}
	if answer == "y" {
		return true
	} else {
		fmt.Println("Thank you! Good bye!")
		return false
	}
}
