package main

import (
	"fmt"
	"log"
	"planner2/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Инициализация базы данных
	config.InitDB()

	// Настройка Telegram-бота
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panicf("Ошибка подключения к Telegram API: %v", err)
	}

	bot.Debug = true
	log.Printf("Бот авторизован под именем %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	u.AllowedUpdates = []string{"message", "callback_query"} // Добавляем обработку CallbackQuery
	updates := bot.GetUpdatesChan(u)

	// Все вопросы по порядку
	questions := []string{
		// Позитивные вопросы
		"Поговорим про Финансовую подушку: чтобы чувствовать себя уверенно на 🌊 волнах экономических кризисов и быть защищенным от 🛑 потери работы – надо иметь финансовую подушку безопасности. Оптимально иметь подушку безопасности на 6 месяцев. Если у тебя такая подушка уже есть – это 👍 отлично! Если нет – напиши, сколько тебе не хватает денег, чтобы чувствовать себя спокойно. Подушка безопасности должна ровняться твоим суммарным расходам за месяц ✖️ 6.",
		"Прекрасный инструмент для создания финансовой свободы – это акции. Если ты готов их покупать – напиши на какую сумму в год ты хотел бы купить акции.",
		"На данном этапе экономики государства облигации тоже отличный инструмент для создания финансовой свободы. Если ты готов их покупать – напиши на какую сумму в год ты хотел бы купить облигации.",
		"Криптовалюта – это отличный инструмент создания капитала на перспективу. Но он высокорискованный не только из-за взлетов и падения цен на крипту, но и из-за большого количества возникающих криптовалют, которые ничем не подкреплены. Но если ты решился, то на какую сумму в год ты хотел бы купить крипту.",
		" Давай помечтаем о Больших проектах. Если ты готов идти в большие проекты, то ты должен четко отдавать себе отчет, что это стоит твоих усилий, времени и финансовых вложений. Сколько стоит твой час работы? Оцифруй своё время, которое планируешь потратить на проект. Добавь деньги, которые нужно вложить физически. Какая сумма необходима для твоего проекта в год?",
		"Планируешь покупку квартиры/дома? Если нет, то ставь «0». А если да, то поздравляю! Если ты что-то продал или накопил или у тебя просто есть деньги, которые пойдут на покупку недвижимости – ты молодец! Напиши ниже, сколько тебе не хватает, чтобы купить жильё мечты? (Если у тебя есть вся сумма на приобретение жилья, то ниже поставь 0).",
		"А если ты планируешь ремонт?! Ремонт – это всегда обновление и новая энергия, но и затраты… Прекрасно, если у тебя есть накопления. Напиши сумму, которой тебе пока не хватает для того, чтобы начать/завершить ремонт.",
		"Планируешь немного поучиться? Инвестиции в себя – это лучшая инвестиция и всегда рост. Сколько тебе нужно денег, чтобы потом кратно вырасти? Напиши сумму, которую планируешь потратить на свое обучение в год.",
		"Обучение детей – это святое. Напиши, сколько надо потратить на их образование.",
		"Может ты задумал купить новую машину? Ей-ей-ей!!! Будет от чего получить удовольствие на дороге! Если ты планируешь продать старую машину или поднакопил уже на новую, это сократит твои вложения! Напиши ниже, сколько тебе не хватает, чтобы купить комфортное перемещение? (Если у тебя есть вся сумма на приобретение машины, то ниже поставь 0).",
		"Мечтаешь о помощниках по дому? Разгрузить себя от быта – отличная идея! Сколько стоит в год услуги помощников по дому? (Обратите внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Няня – это почти необходимость, когда есть дети. 😊 Сколько стоит, чтобы почувствовать себя человеком? (Обратите внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Не стоит забывать про специалистов по здоровью. Здоровье – это мега важно в любом возрасте! Если случаются проблемы со здоровьем, то все остальные сферы автоматически проседают, так как мы вынуждены отвлекаться на себя. Поэтому лучше профилактика, чем лечение! Посчитайте, сколько будут стоить услуги врачей, нутрициолога, иглотерапевта или других нужных вам специалистов в год? (Обратите внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Красивая улыбка требуется? Красивая улыбка – это тоже важно! Как ни странно, но от правильного прикуса зависит даже красивый овал лица. Сколько в этом году будет стоить улыбка мечты?",
		"Массаж, баню, спа. Отдыхом пренебрегать никак нельзя, иначе депрессии, выгорание и еще много умных психологических терминов 😊 Разве тебе это надо? Сколько тебе нужно денег в год, чтобы успевать отдохнуть и расслабиться? (Обратите внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Движение – это жизнь! 💪 Ты как? Готов тратиться на спортзал и/или бассейн? Напиши стоимость годового абонемента.",
		"Мечтаешь об отдыхе на море? Море – это прекрасно! Сколько будет стоить поездка на всю семью?",
		"Планируешь еще путешествия? Если ты планируешь путешествовать куда-то еще, кроме моря, то напиши общий бюджет, который хочешь выделить на поездки.",
		"Обязательно запланируй затраты на свободное время! Свободное время – это в первую очередь переключение, а значит психологическая разгрузка и энергия на новые и текущие дела и проекты. Идеально разгружаться один раз в две недели, или хотя бы один раз в месяц. Напиши, сколько стоит твоё хобби? (Обратите внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Наше окружение – это отображение нас самих и нашего внутреннего состояния. Чтобы поменять своё состояние – надо найти новых друзей, которые мыслят иначе. Сделать это можно на выездах на различные семинары, конференции, мастер-классы по твоей теме, ретриты и т.д. Кем хочешь окружить себя ты, чтобы выйти на новый уровень? Сколько тебе нужно денег, чтобы «пробраться» в это окружение? (Обратите внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Если мечтаешь о кардинальных переменах – запланируй себе шоппинг со стилистом 😉 А ниже просто напиши, сколько это будет тебе стоить.",
		"Дарить подарки родным, близким и друзьям всегда приятно. Но иногда оказывается, что родных и друзей больше, чем свободных денег 😊 Подумай, сколько ты хотел бы тратить на подарки, чтобы чувствовать себя хорошо? (Обратите внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		" Вкусная полезная еда – это также важно, как и забота о здоровье. По сути, это оно и есть плюс удовольствие. Но обычно такая еда дороже, чем макароны… Сколько мечтаешь тратить, чтобы питаться полезно, вкусно и разнообразно? (Обратите внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		" Вдруг ты собрался переезжать?! Переезд затратен, что тут говорить… Но это новый виток жизни и это здорово! Сколько будет стоить твой переезд?",
		//негативные

		" Хочешь оставить в прошлом тревогу о финансах? Если она есть, то сама собой не уйдет ☹ С ней нужно что-то делать. Хорошо бы найти корень проблемы и разобраться с тревогой раз и навсегда. В этом помогают психологи, курсы, специальные книги. Если ты хочешь с этим разобраться, то напиши сумму, которую готов потратить в год. (Обрати внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Долги могут быть и хорошими. Но в целом – долги и есть долги – их надо отдавать. Какова твоя долговая нагрузка, от которой хочешь избавиться в этом году? (Обрати внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Хочешь забыть про некоторые домашние обязанности? Освобождение от чего-то часто влечет за собой финансовые затраты. Если ты решил больше не мыть посуду руками, то нужно купить посудомойку – это затраты. Или не мыть окна самостоятельно – это тоже затраты на клининг. Напиши, сколько ты готов заплатить, чтобы освободить себе время. (Если ты писал сумму в разделе «помощники по дому», то тут повторять эту сумму не нужно).",
		"Может оставить в прошлом самоощущение неудачника? Внутреннее ощущение неудачника не даёт двигаться вперед, зарабатывать столько, сколько ты хочешь и жить ту жизнь, о которой мечтаешь. Это как огромный камень, который привязан к ноге и тянет на дно реки. Но это можно изменить. В этом тоже помогут психологи, курсы, специальные книги. Если ты хочешь с этим разобраться, то напиши сумму, которую готов потратить в год.",
		"Хочешь снизить операционку? Чтобы развязать себе руки и начать придумывать что-то, надо освободиться от ежедневных однообразных действий. А это делегировать и платить за работу другому человеку. Сколько ты готов заплатить, чтобы новые интересные мысли смогли прийти к тебе в голову? (Обрати внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Если ты готов уже к другому уровню жизни, то можно отказаться от общественного транспорта и пересесть на такси, например. Но никогда не делай этого, если твои расходы равны или даже больше твоих доходов. НИКОГДА! Это приведет к финансовому кризису. А если уже готов к новому уровню, то сколько будет стоить твоё такси? (Обрати внимание, что сумма должна быть за год, то есть ежемесячную сумму надо умножить на 12).",
		"Экономить на всём не нужно, но рациональная экономия – это хорошо и «по-взрослому». Подумай, какие ненужные траты, которые, как тебе кажется, поддерживают тебя на определенном уровне жизни, можно убрать и вообще не заметить этого? Указывай примерную цифру в год",
		"И последнее… Посчитай, сколько у тебя уходит денег в месяц на жизнь – коммунальные расходы, повседневные траты и питание. И умножь эту цифру нв 12. Запиши ниже годовой жизненный расход.",
	}

	// Обработка входящих сообщений
	for update := range updates {
		log.Printf("Full Update: %+v", update)
		handleMessage(bot, update, questions)
	}

}

// Обработка сообщений пользователя
func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, questions []string) {
	if update.Message != nil && update.Message.IsCommand() && update.Message.Command() == "start" {
		handleStartCommand(bot, update)
		return
	}

	if update.CallbackQuery != nil {
		handleCallbackQuery(bot, update, questions)
		return
	}

	if update.Message != nil {
		handleTextResponse(bot, update, questions)
	}
}

// Обработка команды /start
func handleStartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userName := update.Message.From.FirstName

	// Сброс данных пользователя
	config.DB.Delete(&config.UserAnswer{}, "user_id = ?", update.Message.From.ID)

	// Первое сообщение с первой кнопкой
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(
		"Привет, <b>%s</b>! 👋\n\nДавай начнем планирование 2025 года! 📅\n\nЯ буду задавать по очереди вопросы, а ты пиши нужные тебе суммы или «0», если тебе эта статья не подходит. 💬\n\nДоговорились? Нажми «Да», если согласен, и мы приступим! 😊",
		userName,
	))
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Да", "start_step1"),
		),
	)
	bot.Send(msg)
}

func handleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update, questions []string) {
	chatID := update.CallbackQuery.Message.Chat.ID
	userID := update.CallbackQuery.From.ID
	callbackData := update.CallbackQuery.Data

	log.Printf("CallbackQuery received: %s", callbackData)

	switch callbackData {
	case "start_step1":
		sendStep1(bot, chatID)
	case "start_step2":
		startQuestions(bot, chatID, userID, questions)
	case "start_negative":
		askQuestion(bot, chatID, userID, questions, 25) // Переходим к первому негативному вопросу
	}

	// Подтверждение обработки
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Ошибка при подтверждении CallbackQuery: %v", err)
	}
}

func sendStep1(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID,
		"Для начала я хочу, чтобы ты определился, какими «приятностями» ты наполнишь свой 2025 год. 🎉\n\nНажми «Да», если готов мечтать и оцифровывать! 🌟",
	)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Да", "start_step2"),
		),
	)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка отправки сообщения на step1: %v", err)
	}
}

func startQuestions(bot *tgbotapi.BotAPI, chatID int64, userID int64, questions []string) {
	msg := tgbotapi.NewMessage(chatID, "Отлично, начнем!")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка отправки сообщения для запуска вопросов: %v", err)
	}

	askQuestion(bot, chatID, userID, questions, 0)
}

// Обработка текстовых ответов
func handleTextResponse(bot *tgbotapi.BotAPI, update tgbotapi.Update, questions []string) {
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID
	userText := update.Message.Text

	// Игнорируем ввод "Да" как текст (это часть интерфейса, а не ответ)
	if userText == "Да" {
		return
	}
	step := getCurrentStep(userID)
	if step >= len(questions) {
		msg := tgbotapi.NewMessage(chatID, "Ты уже ответил на все вопросы. Спасибо за участие!")
		msg.ParseMode = "HTML"
		bot.Send(msg)
		return
	}

	// Сохранение ответа
	if err := config.CreateOrUpdateUserAnswer(userID, questions[step], update.Message.Text, step, step >= len(questions)/2); err != nil {
		log.Printf("Ошибка сохранения ответа: %v", err)
		bot.Send(tgbotapi.NewMessage(chatID, "Произошла ошибка при сохранении ответа. Попробуй снова."))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatID, "Ответ сохранён!"))
	askQuestion(bot, chatID, userID, questions, step+1)
}

// Получение текущего шага из базы данных
func getCurrentStep(userID int64) int {
	var userAnswer config.UserAnswer
	err := config.DB.Where("user_id = ?", userID).Order("step DESC").First(&userAnswer).Error
	if err != nil {
		return 0
	}
	return userAnswer.Step + 1
}

func askQuestion(bot *tgbotapi.BotAPI, chatID int64, userID int64, questions []string, step int) {
	const firstNegativeQuestionIndex = 24 // Индекс первого негативного вопроса

	if step == firstNegativeQuestionIndex {
		msg := tgbotapi.NewMessage(chatID, "Ты уже наполнил свой год приятными моментами. Но есть то, что ты точно хотел бы оставить в прошлом и забыть.")
		bot.Send(msg)

		step = firstNegativeQuestionIndex
	}

	if step < len(questions) {
		msg := tgbotapi.NewMessage(chatID, questions[step])
		bot.Send(msg)
		step++ // Увеличиваем шаг для следующего вопроса
	} else {
		totalAnnual, avgMonthly := config.CalculateFinalSummary(userID)
		msg := tgbotapi.NewMessage(chatID, config.FinalSummary(totalAnnual, avgMonthly))
		msg.ParseMode = "HTML"
		bot.Send(msg)
	}
}
func calculatePositiveSummary(userID int64, midPoint int) float64 {
	var answers []config.UserAnswer
	config.DB.Where("user_id = ? AND step < ?", userID, midPoint).Find(&answers)

	var total float64
	for _, answer := range answers {
		total += answer.Answer
	}

	return total
}
