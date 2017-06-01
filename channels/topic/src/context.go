package topic

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

type ChannelTopicContext struct {
    *Context_t

     Name   String_t
     Data   [16]*unsafe.Pointer
}

var topicChannel = String_t{ len("topic_channel"), "topic_channel" }
var channelTopicContext = &ChannelTopicContext{
    Name:      topicChannel,
    Context_t: NewContext(),
}

func (r *ChannelTopicContext) Set() unsafe.Pointer {
    topic := NewChannelTopic()
    if topic == nil {
        return nil
    }

    topic.name = "topic"

    if this := CreateBackgroundContext(); this != nil {
        if r.WithCancel(this) != Ok {
            return nil
        }
    }

    return unsafe.Pointer(topic)
}

func (r *ChannelTopicContext) Get() []*unsafe.Pointer {
    return r.Data[:]
}
