package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var gBot *tgbotapi.BotAPI
var gToken string
var gChatID int64

func IsStartMessage(update *tgbotapi.Update) bool  {
	return update.Message != nil && update.Message.Text == "/start"
}

func IsCallBackQuerry(update *tgbotapi.Update) {
	return update.CallbackQuery != nil && update.CallbackQuery.Data != ""
}

func init()  {
	// os.Setenv(TOKEN_NAME_IN_OS, ":TOKEN_NAME_IN_OS")
	_ = os.Setenv(TOKEN_NAME_IN_OS, "6382558385:AAG5Ton5bI2CeF4hH02D4D2JxhKZfDXjKbY")
	
	if gToken = os.Getenv(TOKEN_NAME_IN_OS); gToken == ""{
		panic(fmt.Errorf("failed to load enviroment variable %s", TOKEN_NAME_IN_OS))
	}
	var err error
	gBot, err = tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}
  gBot.Debug = true
}




func delay(second uint8)  {
	time.Sleep(time.Second * time.Duration(second))
}

func SendStringMessage(msg string)  {
 gBot.Send(tgbotapi.NewMessage(gChatID, msg))	
}



func SendMessageWithDelay(delayInSec uint8, message string)  {
	SendStringMessage(message)
     delay(delayInSec)
	}


//TODO: Это что видит пользователь и какие сообщения мы отправляем
func printIntro(update *tgbotapi.Update)  {
	sendMessageWithDelay(2, "Привет, котенок") //add эмоджи с котенком
sendMessageWithDelay // здесь будет сообщние с меню
sendMessageWithDelay // здесь будет сообщние с донатом 
sendMessageWithDelay// здесь будет сообщение какое-то еще

}

// это что он отправляет нам на АПИ

func getKeyboardRow(buttonText, buttonCode string) []tgbotapi.InlineKeyboardButton  {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonText, buttonCode))
}


func askToPrintIntro()  {
	msg := tgbotapi.NewMessage(gChatID, "Чтобы получить удовольствие - ТЫК!") //это сообщение которые будет выдавать бот пользователю
    button := getKeyboardRow("ТЫК", BUTTON_CODE_PRINT_INTRO )

	gBot.Send(msg)
}


// выбор действия пользователя после перехода на бота
func ShowMenu(update *tgbotapi.Update)  {
	msg := tgbotapi.NewMessage(gChatID, "Выбери вариант ответа, сладкий")
msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
	getKeyboardRow(BUTTON_CODE_BALANCE),
	getKeyboardRow(BUTTON_CODE_PRINT_MENU),
	getKeyboardRow(BUTTON_CODE_PRINT_INTRO),
	getKeyboardRow(BUTTON_CODE_DONATE),
)

}

func updateProcessing(update *tgbotapi.Update)  {
	choseCode := update.CallbackQuery.Data
  log.Printf("[%T] %s", time.Now(), choseCode)

switch choseCode {
case BUTTON_CODE_PRINT_INTRO:
   printIntro(update) // выводит сообщение
   showMenu(update) // main menu
case BUTTON_CODE_BALANCE: // будет показывать баланс пользователя
 showMenu(update)
case BUTTON_CODE_PRINT_MENU: // выводит в общее меню
 showMenu(update)
case BUTTON_CODE_DONATE: // донатик кнопочка
showMenu(update)
}

}


func main()  {
	
	

	log.Printf("Authorized on account %s", gBot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := gBot.GetUpdatesChan(updateConfig)
      
	for update := range gBot.GetUpdatesChan(updateConfig) {
		
		if update.CallbackQuery(&update) {
			update.Processing(&update)
		}
		
		if IsStartMessage (&update) 
		{ 
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
            
         gChatID = update.Message.ChatID
		 askToPrintIntro(&update)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			gBot.Send(msg)
		}
	}
}
