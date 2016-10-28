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
	repo.Kafka.SendEventLogAsync(event, model.EventTopic, model.EventKey)
	return nil, 200
}
