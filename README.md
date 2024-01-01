<h1 align="center">Telegram News Bot</h1>

<p align="center"> 
  <img src="https://github.com/isabellymm/NewsBot/assets/96357748/e463b002-3320-4dfc-8d07-5491e1eaf37e" width="300">
</p>


## Overview
The Telegram News Bot is a simple Telegram bot that provides users with the latest news headlines from different countries. Users can choose a specific country to receive news updates.

## Features
- **Country Selection:** Users can choose from a list of supported countries to receive news updates.
- **Top Headlines:** The bot fetches top headlines for the selected country using the News API.
- **Interactive Interface:** The bot provides an interactive interface for users to easily navigate and select their desired country.

## Technologies Used
- [Go](https://golang.org/): The programming language used for bot development.
- [Telegram Bot API](https://core.telegram.org/bots/api): Interact with Telegram using the official Bot API.
- [News API](https://newsapi.org/): Fetch news headlines from various sources worldwide.

## Getting Started
1. **Clone the Repository:**
   ```bash
   git clone <repository-url>
   cd <repository-directory>
2. **Install Dependencies:**
   ```bash
   # Install required Go packages
    go get -u github.com/go-telegram-bot-api/telegram-bot-api
    go get -u gopkg.in/resty.v1

3. **Configure API Tokens:**
    - Obtain a Telegram Bot Token from @BotFather.
    - Obtain a News API Token from News API.
  
4. **Update Configuration:**
    - Replace <TELEGRAM_BOT_TOKEN> and <NEWS_API_TOKEN> with your actual tokens in the source code.

5. **Interact with the Bot:**
    - Start a conversation with your bot on Telegram.
    - Use the /start command to initiate the bot.
    - Use the /news command to choose a country and receive news updates.

## Contributing
Contributions are welcome! Feel free to open issues or pull requests for any improvements or new features.

## License
This project is licensed under the MIT License.
