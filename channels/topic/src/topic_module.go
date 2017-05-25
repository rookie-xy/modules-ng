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
    Modules = Load(Modules, &Topic{Module_t:topicModule})
}
