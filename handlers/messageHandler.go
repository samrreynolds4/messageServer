package handlers

import (
	"log"
	"net/http"
)

type httpHandler func(writer http.ResponseWriter, req *http.Request)

type MessageHandler interface {
	PostMessage(writer http.ResponseWriter, req *http.Request)
	GetMessage(writer http.ResponseWriter, req *http.Request)
}

type MessageHandlerImpl struct{}

func NewMessageHandler() *MessageHandlerImpl {
	return &MessageHandlerImpl{}
}

func (handler *MessageHandlerImpl) postMessage(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		log.Printf("posting")
	default:
		log.Printf("anything else")
	}
}

func (handler *MessageHandlerImpl) getMessage(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		log.Printf("get")
	default:
		log.Printf("anything else")
	}
}

func (handler *MessageHandlerImpl) NewPostMessageHandler() (string, httpHandler) {
	return "/message/post/", handler.postMessage
}

func (handler *MessageHandlerImpl) NewGetMessageHandler() (string, httpHandler) {
	return "/message/get/", handler.getMessage
}
