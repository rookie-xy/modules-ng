package topic

import (
    . "github.com/rookie-xy/worker/types"
)

type Subscribe struct {
    name       string
    subscribe  chan Event
}

func NewSubscribe(name string) *Subscribe {
    return &Subscribe{
        name:      name,
        subscribe: make(chan Event, 10000),
    }
}

func (r *Subscribe) Push(in Event) int {
    return Ok
}

func (r *Subscribe) Pull(out Event) int {
    return Ok
}
