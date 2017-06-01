package topic

import (
    . "github.com/rookie-xy/worker/types"
    "unsafe"
)

type TopicCycle struct {
    Log
}

func NewTopicCycle(log Log) *TopicCycle {
    object := &TopicCycle{
        Log: log,
    }

    return object
}

func (r *TopicCycle) Init(configure interface{}) int {
    return Ok
}

func (r *TopicCycle) Start(p *unsafe.Pointer) int {
    return Ok
}

func (r *TopicCycle) Stop() int {
    return Ok
}
