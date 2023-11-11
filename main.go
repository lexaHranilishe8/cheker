package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getPrice(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении ответа:", err)
		return ""
	}

	responseString := string(body)
	if strings.Contains(responseString, "lowest_price") {
		startIndex := strings.Index(responseString, "lowest_price") + len("lowest_price\":\"")
		endIndex := strings.Index(responseString[startIndex:], "\"")
		lowestPrice := responseString[startIndex : startIndex+endIndex]
		log.Println("Найден lowest_price:", lowestPrice)

		return lowestPrice
	}

	return ""
}
func main() {
	bot, err := tgbotapi.NewBotAPI("6898487305:AAGTZdutyOhXYI6IvTv3JuZrGUH8s4HSu-o")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "великий бот берт цену предмета с тп стима и кидает ее тебе используй команду /price для капсулы /price1 для кейса")
			bot.Send(msg)
		case "/price":
			url := "https://steamcommunity.com/market/priceoverview/?appid=730&currency=1&market_hash_name=Paris%202023%20Legends%20Sticker%20Capsule"
			price := getPrice(url)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Актуальная цена предмета Капсула с наклейками легенд : "+price)
			bot.Send(msg)
		case "/price1":
			url := "https://steamcommunity.com/market/priceoverview/?appid=730&currency=1&market_hash_name=Prisma%202%20Case"
			price := getPrice(url)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Актуальная цена предмета Prisma Case: "+price)
			bot.Send(msg)
		case "/price2":
			url := "https://steamcommunity.com/market/priceoverview/?appid=730&currency=1&market_hash_name=Dreams%20%26%20Nightmares%20Case"
			price := getPrice(url)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Актуальная цена предмета Dreams & Nightmares Case: "+price)
			bot.Send(msg)
		case "/price3":
			url := "https://steamcommunity.com/market/priceoverview/?appid=730&currency=1&market_hash_name=Fracture%20Case"
			price := getPrice(url)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Актуальная цена предмета Fracture Case: "+price)
			bot.Send(msg)
		case "/price4":
			url := "https://steamcommunity.com/market/priceoverview/?appid=730&currency=1&market_hash_name=Snakebite%20Case"
			price := getPrice(url)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Актуальная цена предмета Snakebite Case: "+price)
			bot.Send(msg)
		case "/price5":
			url := "https://steamcommunity.com/market/priceoverview/?appid=730&currency=1&market_hash_name=Recoil%20Case"
			price := getPrice(url)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Актуальная цена предмета Recoil Case: "+price)
			bot.Send(msg)
		case "/price6":
			url := "https://steamcommunity.com/market/priceoverview/?appid=730&currency=1&market_hash_name=Revolution%20Case"
			price := getPrice(url)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Актуальная цена предмета Revolution Case: "+price)
			bot.Send(msg)
		case "/price7":
			url := "https://steamcommunity.com/market/priceoverview/?appid=730&currency=1&market_hash_name=Paris%202023%20Contenders%20Sticker%20Capsule"
			price := getPrice(url)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Актуальная цена предмета Капсула с наклейками соперников: "+price)
			bot.Send(msg)
		}

	}
}
