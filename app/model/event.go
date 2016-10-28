package model

import (
	"encoding/json"
	"time"
)

const (
	EventTopic = "event_logs"
	EventKey   = ""
)

type Event struct {
	// User info
	UserId     int
	CountryId  int
	CityId     int
	DistrictId int

	// Timestamp
	begin time.Time
	end   time.Time

	// Client info
	Platform        int
	Version         string
	Carrier         string
	CarrierStrength string
	IpAddr          string
	Client          string
	Device          string

	// Object info
	Event      string
	ObjectType string
	ObjectId   string

	// Kafka encoded
	encoded []byte
	err     error
}

func (e *Event) ensureEncoded() {
	if e.encoded == nil && e.err == nil {
		e.encoded, e.err = json.Marshal(e)
	}
}

func (e *Event) Length() int {
	e.ensureEncoded()
	return len(e.encoded)
}

func (e *Event) Encode() ([]byte, error) {
	e.ensureEncoded()
	return e.encoded, e.err
}
