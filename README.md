# personal-website

My personal portfolio website built with Go backend and modern web technologies.

![Status](https://img.shields.io/badge/Status-Active_Development-orange)

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/Belixk/personal-website.git
cd personal-website

# Install dependencies
go mod download

# Run the server
go run main.go

# Open in browser
# http://localhost:8080
```

## ⚙️ Configuration

### Telegram Notifications
1. Create a bot via @BotFather in Telegram
2. Get your bot token
3. Find your Chat ID (send a message to the bot)
4. Create `.env` file in the root:
```
TELEGRAM_BOT_TOKEN=your_bot_token
TELEGRAM_CHAT_ID=your_chat_id
```

## 🛠️ Tech Stack

- **Backend:** Go, Gin Framework
- **Frontend:** HTML5, CSS3, JavaScript
- **Notifications:** Telegram Bot API
- **Validation:** Gin binding + custom validation

## 📁 Project Structure

```
personal-website/
├── main.go                
├── handlers/            
│   ├── contact.go        
│   ├── skills.go       
│   └── pages.go         
├── models/              
│   └── contact.go        
├── services/               
│   ├── telegram.go        
│   ├── validation.go      
│   └── storage.go         
├── templates/              
│   └── index.html
└── static/                 
    ├── css/
    ├── js/
    └── images/
```

## 🗺️ Roadmap

- [x] Basic contact form
- [x] Telegram integration  
- [x] Input validation
- [ ] Email notifications
- [ ] Admin panel for messages
- [ ] Docker containerization
- [ ] Deployment to production

## 📡 API Endpoints

- `GET /` - Home page
- `GET /api/skills` - Skills list (JSON)
- `POST /contact` - Contact form submission

## 👨‍💻 Author

**Maxim A.**
- GitHub: [@Belixk](https://github.com/Belixk)
