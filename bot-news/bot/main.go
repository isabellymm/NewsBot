package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/resty.v1"
)

const (
	telegramToken = "6664933158:AAHma-3hwjT6WzU_dTeLrSHbrbKsbJGphIY"
	newsApiToken  = "c5a0e7df5c8f42b4a16ab238b41a29dd"
)

var supportedCountries = map[string]string{
	"ae": "United Arab Emirates",
	"ar": "Argentina",
	"at": "Austria",
	"au": "Australia",
	"be": "Belgium",
	"bg": "Bulgaria",
	"br": "Brazil",
	"ca": "Canada",
	"ch": "Switzerland",
	"cn": "China",
	"co": "Colombia",
	"cz": "Czech Republic",
	"de": "Germany",
	"eg": "Egypt",
	"es": "Spain",
	"fr": "France",
	"gb": "United Kingdom",
	"gr": "Greece",
	"hk": "Hong Kong",
	"hu": "Hungary",
	"id": "Indonesia",
	"ie": "Ireland",
	"il": "Israel",
	"in": "India",
	"it": "Italy",
	"jp": "Japan",
	"kr": "South Korea",
	"l":  "Luxembourg",
	"lt": "Lithuania",
	"lv": "Latvia",
	"ma": "Morocco",
	"mx": "Mexico",
	"my": "Malaysia",
	"ng": "Nigeria",
	"nl": "Netherlands",
	"no": "Norway",
	"nz": "New Zealand",
	"ph": "Philippines",
	"pl": "Poland",
	"pt": "Portugal",
	"ro": "Romania",
	"rs": "Serbia",
	"ru": "Russia",
	"sa": "Saudi Arabia",
	"se": "Sweden",
	"sg": "Singapore",
	"sk": "Slovakia",
	"th": "Thailand",
	"tr": "Turkey",
	"tw": "Taiwan",
	"ua": "Ukraine",
	"us": "United States",
	"ve": "Venezuela",
	"za": "South Africa",
}

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
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

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text = "Welcome! I am a news bot. Use /news to get the latest news."
			case "news":
				displayCountryOptions(bot, update.Message.Chat.ID)
			case "setcountry":
				if len(update.Message.CommandArguments()) > 0 {
					countryCode := strings.ToLower(update.Message.CommandArguments())
					if countryName, ok := supportedCountries[countryCode]; ok {
						news, err := getNewsByCountry(countryCode)
						if err != nil {
							msg.Text = "Sorry, there was an error fetching the news."
						} else {
							msg.Text = fmt.Sprintf("Notícias de %s:\n\n%s", countryName, news)
						}
					} else {
						msg.Text = "Invalid country code. Use /news to see available country options."
					}
				} else {
					msg.Text = "Please provide a valid country code. Use /news to see available country options."
				}
			default:
				msg.Text = "Unknown command"
			}
			bot.Send(msg)
		}
	}
}

func displayCountryOptions(bot *tgbotapi.BotAPI, chatID int64) {
	var msgText strings.Builder
	msgText.WriteString("Choose a country to receive news:\n\n")
	for code, country := range supportedCountries {
		msgText.WriteString(fmt.Sprintf("/setcountry %s - %s\n", code, country))
	}
	msg := tgbotapi.NewMessage(chatID, msgText.String())
	bot.Send(msg)
}

func createCountryKeyboard() tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	for code, name := range supportedCountries {
		button := tgbotapi.NewInlineKeyboardButtonData(name, fmt.Sprintf("setcountry %s", code))
		buttons = append(buttons, button)
	}

	row := tgbotapi.NewInlineKeyboardRow(buttons...)
	return tgbotapi.NewInlineKeyboardMarkup(row)
}

func getNewsByCountry(countryCode string) (string, error) {
	url := "https://newsapi.org/v2/top-headlines"
	params := map[string]string{
		"country": countryCode,
		"apiKey":  newsApiToken,
	}

	resp, err := resty.R().
		SetQueryParams(params).
		Get(url)

	if err != nil {
		return "", err
	}

	var newsResponse NewsAPIResponse
	err = json.Unmarshal(resp.Body(), &newsResponse)
	if err != nil {
		return "", err
	}

	var newsText strings.Builder
	for i, article := range newsResponse.Articles {
		if i < 4 {
			newsText.WriteString(fmt.Sprintf("📰 %s\n%s\n\n", article.Title, article.URL))
		} else {
			break
		}
	}

	return newsText.String(), nil
}

type NewsAPIResponse struct {
	Articles []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	} `json:"articles"`
}
