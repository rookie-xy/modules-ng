package elasticsearch

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

type ElasticSearchOutputContext struct {
    *Context_t

     Name   String_t
     Data   [32]*unsafe.Pointer
}

var elasticSearchOutput = String_t{ len("elasticsearch_output"), "elasticsearch_output" }
var outputElasticSearchContext = &ElasticSearchOutputContext{
    Name: elasticSearchOutput,
}

func (r *ElasticSearchOutputContext) Set() unsafe.Pointer {
    elasticsearch := NewElasticSearchOutput()
    if elasticsearch == nil {
        return nil
    }

    //stdout.subscribe = "zhang yue"

    return unsafe.Pointer(elasticsearch)
}

func (r *ElasticSearchOutputContext) Get() []*unsafe.Pointer {
    return r.Data[:]
}
