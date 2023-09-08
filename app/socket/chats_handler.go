package socket

import (
	"Kavka/app/presenters"
	"Kavka/internal/domain/chat"
	"log"
)

func NewChatsHandler(args MessageHandlerArgs) bool {
	event := args.message.Event

	switch event {
	case "get_or_create_chat":
		return GetOrCreateChat(event, args)
	case "create_group":
		return CreateGroup(event, args)
	}

	return false
}

func GetOrCreateChat(event string, args MessageHandlerArgs) bool {
	chatType := args.message.Data["chat_type"]
	username := args.message.Data["username"]

	if chatType != nil && username != nil {
		if chatType == chat.ChatTypeDirect {
			chat, err := args.socketService.chatService.GetOrCreateChat(chatType.(string), username.(string), args.staticID)

			if err != nil {
				log.Println("GetOrCreateChat error in socket:", err)
				return false
			}

			args.conn.WriteJSON(presenters.ChatAsJSON(event, chat))

			return true
		} else if chatType == chat.ChatTypeGroup || chatType == chat.ChatTypeChannel {
			println(username)
			return true
		}
	}

	return false
}

func CreateGroup(event string, args MessageHandlerArgs) bool {
	title := args.message.Data["title"]
	username := args.message.Data["username"]
	description := args.message.Data["description"]

	if title != nil && username != nil && description != nil {
		chat, err := args.socketService.chatService.CreateGroup(args.staticID, title.(string), username.(string), description.(string))
		if err != nil {
			return false
		}

		args.conn.WriteJSON(presenters.ChatAsJSON(event, chat))

		return true
	}

	return false
}
