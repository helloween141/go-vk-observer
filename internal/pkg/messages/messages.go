package messages

const StartMessage = "👋 Привет, *%s!*\n\n" +
	"Я отслеживаю 👀 новые посты в сообществах ВКонтакте.\n\n" +
	"📌 Сначала добавь идентификаторы интересных тебе сообществ.\n\n" +
	"⏰ Раз в час я буду проверять: появились ли новые посты?\n\n" +
	"📩 Если появились, то отправлю тебе сообщения с их содержанием.\n\n"
const MenuMessage = "👨‍💻 *Выбери действие...*\n\nЧтобы добавить или удалить сообщество введи URL или идентификатор.\n\n*Пример:* `https://vk.com/vk` или `vk`"
const AddSlugSuccessful = "✅ Сообщество `%s` добавлено!"
const DeleteSlugSuccessful = "✅ Сообщество `%s` удалено!"
const SlugNotFound = "❌ Сообщество `%s` не найдено!"
const SlugAlreadyExists = "❌ Сообщество `%s` уже существует!"
const SlugsListIsEmpty = "ℹ️ Список сообществ пустой!"
const SlugsListTitle = "📝 *Список сообществ:*\n\n"
const SlugsListMessage = "%d. [%s](https://vk.com/%s) /// `%s`\n\n"
const AddSlugTooltip = "Введи ссылку или идентификатор для добавления сообщества"
const DeleteSlugTooltip = "Введи ссылку или идентификатор для удаления сообщества"
