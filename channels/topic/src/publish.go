package topic

import (
    . "github.com/rookie-xy/worker/types"
)

type Publish struct {
    topic    string
    publish  chan Event
}

func NewPublish(topic string) *Publish {
    return &Publish{
        topic:   topic,
        publish: make(chan Event, 10000),
    }
}

func (r *Publish) Publish(events Event) int {
    r.publish <- events
    return Ok
}
