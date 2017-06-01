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
     filter   Array_t
}

func NewChannelTopic() *ChannelTopic {
    return &ChannelTopic{}
}



