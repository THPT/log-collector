package model

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	TopicPageView = "page_view_logs"
	TopicClick    = "click_logs"
	TopicOrder    = "order_logs"
	TopicKey      = ""
)

type Event struct {
	Ip         string
	CreatedAt  int64
	Agent      string
	Uuid       string
	Referrer   string
	Url        string
	Metric     string
	ProductId  string
	VideoId    string
	OrderId    int
	CustomerId int

	// Kafka encoded
	encoded []byte
	err     error
}

func ParseEvent(c *gin.Context) Event {
	var e Event
	e.Uuid = c.Query("uuid")
	e.Referrer = c.Query("referrer")
	e.Url = c.Query("url")
	e.Metric = c.Query("metric")
	e.ProductId = c.Query("product")
	e.VideoId = c.Query("video")
	orderStr := c.DefaultQuery("order", "0")
	e.OrderId, _ = strconv.Atoi(orderStr)
	CustomerIdStr := c.DefaultQuery("customer_id", "0")
	e.CustomerId, _ = strconv.Atoi(CustomerIdStr)
	e.CreatedAt = time.Now().Unix()
	e.Agent = c.Request.Header.Get("User-Agent")
	e.Ip = strings.Split(c.Request.RemoteAddr, ":")[0]
	return e
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
