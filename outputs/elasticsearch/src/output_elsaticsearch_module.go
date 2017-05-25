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

func (r *OutputElasticSearchContext) Create() unsafe.Pointer {
    elasticsearch := NewOutputElasticSearch()
    if elasticsearch == nil {
        return nil
    }

    //stdout.subscribe = "zhang yue"

    return unsafe.Pointer(elasticsearch)
}

func (r *OutputElasticSearchContext) GetDatas() []*unsafe.Pointer {
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
    context := r.Context.GetDatas()

    for _, v := range context {
        if v != nil {
            this := (*OutputElasticSearch)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            for i := 0; i < this.subscribe.GetLength(); i++ {
                fmt.Println(this.subscribe.GetData(i))
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

func (r *OutputElasticSearch) Main(cfg *Configure_t) int {
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
    Modules = Load(Modules, &OutputElasticSearch{Module_t:outputElasticSearchModule})
}
