package messages

const StartMessage = "👋 *Hello, %s!*\n\n" +
	"I will observe 👀 VK groups by their *slugs* and notify 🔔 you about new posts.\n\n" +
	"*Slug* - is the identifier of VK group.\n" +
	"You can get it here: `https://vk.com/<slug>`\n\n" +
	"*Available commands:*\n" +
	"1. /add *<slug>*\n" +
	"2. /delete *<slug>*\n" +
	"3. /list"
const AddSlugSuccessful = "✅ Slug `%s` was added"
const DeleteSlugSuccessful = "✅ Slug `%s` was deleted"
const SlugNotFound = "❌ Slug `%s` not found"
const SlugIsEmpty = "❌ Slug is empty"
const SlugAlreadyExists = "❌ Slug `%s` already exists"
const VkPostMessage = "🆕️ [*%s*](https://vk.com/%s) | 🗓 %s\n\n%s"
const SlugsListMessage = "%d. [%s](https://vk.com/%s) | `%s`\n\n"
const InvalidCommand = "❌ Invalid command"
const CommonError = "❌ Something went wrong"
