package topic

import (
      "fmt"
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

func (r *ChannelTopic) Init(o *Option_t) int {
    configure := o.Configure_t
    if configure == nil {
        return Error
    }

    context := r.Context.Configure()
    for _, v := range context {
        if v != nil {
            this := (*ChannelTopic)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            topic := NewTopic(this.name)
            if topic.Create(this.filter) == Error {
                return Error
            }

            configure.Channels = append(configure.Channels, topic)
        } else {
            break
        }
    }

    return Ok
}

func (r *ChannelTopic) Main(c *Configure_t) int {
    for _, v := range c.Channels {
        go v.Intercept()
    }

    fmt.Println("Topic main")
    return Ok
}

func (r *ChannelTopic) Exit() int {
    fmt.Println("Topic exit")
    return Ok
}
