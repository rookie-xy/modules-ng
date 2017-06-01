package elasticsearch

import (
      "fmt"
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

func (r *ElasticSearchOutput) Init(o *Option_t) int {
    configure := o.Configure_t
    if configure == nil {
        return Error
    }

    elasticsearch := elasticSearchOutput.Data.(string)

    context := r.Context.Configure()
    for _, v := range context {
        if v != nil {
            this := (*ElasticSearchOutput)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            this.Output = NewElasticSearch()

            for i := 0; i < this.subscribe.GetLength(); i++ {
                topic := this.subscribe.GetData(i)

                for _, v := range configure.Channels {
                    if topic := v.Register(topic, elasticsearch); topic != nil {
                        configure.Channels = append(configure.Channels, topic)
                    }
                }
            }

            for i := 0; i < this.cluster.GetLength(); i++ {
                fmt.Println(this.cluster.GetData(i))
            }

        } else {
            break
        }
    }

    return Ok
}

func (r *ElasticSearchOutput) Main(c *Configure_t) int {
    var channel []Channel
    elasticsearch := elasticSearchOutput.Data.(string)

    for _, topic := range c.Channels {
        if topic.Type(elasticsearch) != Ok {
            continue
        }

        channel = append(channel, topic)
        //go r.Output(topic, r.Output)
    }

    for {
        for _, topic := range channel {
            events := topic.Pull(nil)
            if events != nil {
                if r.Writer(events) != Ok {
                    continue
                }

                topic.Commit()
            }
        }

        // TODO p->done()
    }

    fmt.Println("output main")
    return Ok
}

func (r *ElasticSearchOutput) Exit() int {
    fmt.Println("output exit")
    return Ok
}
