/*
 * Copyright (C) 2017 Meng Shi
 */

package stdout

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
"fmt"
)

type OutputElasticSearch struct {
    *Module_t

     subscribe  Array_t
     cluster    Array_t

     Output
}

func NewOutputElasticSearch() *OutputElasticSearch {
    return &OutputElasticSearch{}
}

type OutputElasticSearchContext struct {
    Name   String_t
    Data   [32]*unsafe.Pointer
}

var elasticSearchOutput = String_t{ len("elasticsearch_output"), "elasticsearch_output" }
var outputElasticSearchContext = &OutputElasticSearchContext{
    Name: elasticSearchOutput,
}

func (r *OutputElasticSearchContext) Set() unsafe.Pointer {
    elasticsearch := NewOutputElasticSearch()
    if elasticsearch == nil {
        return nil
    }

    //stdout.subscribe = "zhang yue"

    return unsafe.Pointer(elasticsearch)
}

func (r *OutputElasticSearchContext) Get() []*unsafe.Pointer {
    return r.Data[:]
}

var (
    elasticSearchSubscribe = String_t{ len("subscribe"), "subscribe" }
    cluster = String_t{ len("cluster"), "cluster" }
    outputElasticSearch OutputElasticSearch
)

var outputElasticSearchCommands = []Command_t{

    { elasticSearchSubscribe,
      ELASTICSEARCH_CONFIG|CONFIG_ARRAY,
      SetArray,
      0,
      unsafe.Offsetof(outputElasticSearch.subscribe),
      nil },


    { cluster,
      ELASTICSEARCH_CONFIG|CONFIG_ARRAY,
      SetArray,
      0,
      unsafe.Offsetof(outputElasticSearch.cluster),
      nil },

    NilCommand,
}

func (r *OutputElasticSearch) Init(o *Option_t) int {
    configure := o.Configure_t
    if configure == nil {
        return Error
    }

    elasticsearch := elasticSearchOutput.Data.(string)

    context := r.Context.Get()
    for _, v := range context {
        if v != nil {
            this := (*OutputElasticSearch)(unsafe.Pointer(uintptr(*v)))
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

func (r *OutputElasticSearch) Main(c *Configure_t) int {
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
                r.Writer(events)
            }
        }

        // TODO p->done()
    }

    fmt.Println("output main")
    return Ok
}


func (r *OutputElasticSearch) Exit() int {
    fmt.Println("output exit")
    return Ok
}

var outputElasticSearchModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    outputElasticSearchContext,
    outputElasticSearchCommands,
    ELASTICSEARCH_MODULE,
}

func init() {
    Modules = append(Modules, &OutputElasticSearch{Module_t:outputElasticSearchModule})
}
