package entity

import (
	"log-collector/app/model"
	"log-collector/app/repo"
)

type EventInterface interface {
	SendEvent(*model.Event) (error, int)
}

type Event struct {
}

func NewEventEntity() EventInterface {
	return &Event{}
}

func (r *Event) SendEvent(event *model.Event) (error, int) {
	var topic string
	switch event.Metric {
	case "pageview":
		topic = model.TopicPageView
	case "click":
		topic = model.TopicClick
	case "order":
		topic = model.TopicOrder
	default:
		topic = ""
	}
	if topic != "" {
		repo.Kafka.SendEventLogAsync(event, topic, model.TopicKey)
	}
	return nil, 200
}
