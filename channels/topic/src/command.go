package topic

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

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

var (
    name   = String_t{ len("name"), "name" }
    filter = String_t{ len("filter"), "filter" }
    channelTopic ChannelTopic
)

var channelTopicCommands = []Command_t{

    { name,
      TOPIC_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(channelTopic.name),
      nil },

    { filter,
      TOPIC_CONFIG|CONFIG_VALUE,
      SetArray,
      0,
      unsafe.Offsetof(channelTopic.filter),
      nil },

    NilCommand,
}
