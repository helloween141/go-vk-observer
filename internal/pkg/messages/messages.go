package messages

const StartMessage = "👋 Привет, *%s!*\n\n" +
	"Я отслеживаю 👀 новые посты в сообществах ВКонтакте.\n\n" +
	"📌 Сначала добавь идентификаторы интересных тебе сообществ.\n\n" +
	"⏰ Раз в час я буду проверять: появились ли новые посты?\n\n" +
	"📩 Если появились, то отправлю тебе сообщения с их содержанием.\n\n" +
	"*Где взять идентификатор?*:\nНапример: `https://vk.com/vk`, где `vk` - это идентификатор. \n\n" +
	"*Доступные команды:*\n" +
	"/add *<id>* - добавить идентификатор\n" +
	"/delete *<id>* - удалить идентификатор\n" +
	"/list - список добавленных идентификатор"
const AddSlugSuccessful = "✅ Идентификатор `%s` добавлен"
const DeleteSlugSuccessful = "✅ Идентификатор `%s` удален"
const SlugNotFound = "❌ Идентификатор `%s` не найден"
const SlugIsEmpty = "❌ Идентификатор не задан"
const SlugAlreadyExists = "❌ Идентификатор `%s` уже существует"
const SlugsListIsEmpty = "ℹ️ Список идентификаторов пустой!"
const InvalidCommand = "❌ Неверная команда"
const CommonError = "❌ Что-то пошло не так"
const VkPostMessage = "🆕️ [**%s**](https://vk.com/%s) /// 🗓 %s\n\n%s```Репост\n%s```"
const SlugsListMessage = "%d. [%s](https://vk.com/%s) /// `%s`\n\n"
