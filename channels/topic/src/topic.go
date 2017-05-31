/*
 * Copyright (C) 2017 Meng Shi
 */

package topic

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    TOPIC_MODULE = CHANNEL_MODULE|MAIN_MODULE
    TOPIC_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

type Topic struct {
    *Module_t

    *Publish
     subscribes []*Subscribe

     Filter
}

func NewTopic(name string) *Topic {
    return &Topic{
        Publish: NewPublish(name),
        Filter: filter,
    }
}

func (r *Topic) Create(filter Filter) int {
    r.Filter = filter
    return Ok
}

func (r *Topic) Remove() int {
    return Ok
}

func (r *Topic) New() Channel {
    topic := NewTopic(r.topic)

    if this := r.Filter; this != nil {
        topic.Filter = this
    }

    return topic
}

func (r *Topic) Register(topic string, name string) Channel {
    if r.topic != topic {
        return Ignore
    }

    subscribe := NewSubscribe(name)
    r.subscribes = append(r.subscribes, subscribe)

    return subscribe
}

func (r *Topic) Push(in Event) int {
    r.Publish(in)
    return Ok
}

func (r *Topic) Intercept() int {

    for {
        select {

        case events := <-r.publish:
            r.Washing(events)
            r.Pull(events)

        default:
        }
    }

    return Ok
}

func (r *Topic) Pull(out Event) int {
    for _, v := range r.subscribe {
        v.Push(out)
    }

    return Ok
}

func (r *Topic) Type(name string) int {
    if r.topic != name {
        return Ignore
    }

    return Ok
}

var	topic = String_t{ len("topic"), "topic" }
var topicCommands = []Command_t{

    { topic,
      TOPIC_CONFIG,
      topicBlock,
      0,
      0,
      nil },

    NilCommand,
}

func topicBlock(c *Configure_t, _ *Command_t, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := TOPIC_CONFIG|CONFIG_VALUE
    Block(c, Modules, TOPIC_MODULE, flag)

    return Ok
}

var topicModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    topicCommands,
    CHANNEL_MODULE,
}

func init() {
    Modules = append(Modules, &Topic{Module_t:topicModule})
}
