package TelegramBot

import (
	"math/rand"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

//InitTBot Telega
func InitTBot() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	log := logger.Sugar()

	bot, err := tgbotapi.NewBotAPI("1344007633:AAHQM-4JAAF3ktfO-Jbr5xsB-oCGuz-r2dI")
	if err != nil {
		log.Error(err)
	}

	bot.Debug = false

	log.Info("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		switch update.Message.Text {
		case "/start":
			//Send message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я - Бот статистики по профилю партнёра Почты России.\n 'Отчет' - получить статистику за сегодня.\n 'Отчет Н' - получить статистику за Неделю.\n 'Отчет М' - получить статистику за Месяц.\n 'Отчет К' - получить статистику за Квартал.")
			bot.Send(msg)
		case "Отчет":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Продаж за Сегодня: "+strconv.Itoa(rand.Intn(100))+"\nНа сумму: "+strconv.Itoa(rand.Intn(9000)))
			bot.Send(msg)
		case "Отчет Н":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Продаж за Неделю: "+strconv.Itoa(rand.Intn(1000))+"\nНа сумму: "+strconv.Itoa(rand.Intn(90000)))
			bot.Send(msg)
		case "Отчет М":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Продаж за Месяц: "+strconv.Itoa(rand.Intn(10000))+"\nНа сумму: "+strconv.Itoa(rand.Intn(900000)))
			bot.Send(msg)
		case "Отчет К":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Продаж за Квартал: "+strconv.Itoa(rand.Intn(100000))+"\nНа сумму: "+strconv.Itoa(rand.Intn(9000000)))
			bot.Send(msg)
		default:
			//Send message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я - Бот статистики по профилю партнёра Почты России.\n 'Отчет' - получить статистику за сегодня.\n 'Отчет Н' - получить статистику за Неделю.\n 'Отчет М' - получить статистику за Месяц.\n 'Отчет К' - получить статистику за Квартал.")
			bot.Send(msg)
		}

	}
}
