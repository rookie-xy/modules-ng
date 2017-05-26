/*
 * Copyright (C) 2017 Meng Shi
 */

package topic

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
//    . "github.com/rookie-xy/plugins/channels/topic"
"fmt"
)

type ChannelTopic struct {
    *Module_t

     name     string
}

func NewChannelTopic() *ChannelTopic {
    return &ChannelTopic{}
}

type ChannelTopicContext struct {
    Name   String_t
    Data   [16]*unsafe.Pointer
}

var topicChannel = String_t{ len("topic_channel"), "topic_channel" }
var channelTopicContext = &ChannelTopicContext{
    Name: topicChannel,
}

func (r *ChannelTopicContext) Set() unsafe.Pointer {
    topic := NewChannelTopic()
    if topic == nil {
        return nil
    }

    topic.name = "topic"

    return unsafe.Pointer(topic)
}

func (r *ChannelTopicContext) Get() []*unsafe.Pointer {
    return r.Data[:]
}

var (
    name = String_t{ len("name"), "name" }
    channelTopic ChannelTopic
)

var channelTopicCommands = []Command_t{

    { name,
      TOPIC_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(channelTopic.name),
      nil },

    NilCommand,
}

func (r *ChannelTopic) Init(o *Option_t) int {
    configure := o.Configure_t
    if configure == nil {
        return Error
    }

    context := r.Context.Get()
    for _, v := range context {
        if v != nil {
            this := (*ChannelTopic)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            //topic := NewTopic(this.name)
            //topic.Init(nil)

            fmt.Println(this.name)
            //configure.Channel = append(configure.Channel, topic)
        } else {
            break
        }
    }

    return Ok
}

func (r *ChannelTopic) Main(cfg *Configure_t) int {
    fmt.Println("Topic main")
    return Ok
}

func (r *ChannelTopic) Exit() int {
    fmt.Println("Topic exit")
    return Ok
}

var channelTopicModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    channelTopicContext,
    channelTopicCommands,
    TOPIC_MODULE,
}

func init() {
    Modules = append(Modules, &ChannelTopic{Module_t:channelTopicModule})
}
