package elasticsearch

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

var elasticsearch = String_t{ len("elasticsearch"), "elasticsearch" }
var elasticSearchCommands = []Command_t{

    { elasticsearch,
      ELASTICSEARCH_CONFIG,
      elasticSearchBlock,
      0,
      0,
      nil },

    NilCommand,
}

func elasticSearchBlock(c *Configure_t, _ *Command_t, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := ELASTICSEARCH_CONFIG|CONFIG_VALUE
    Block(c, Modules, ELASTICSEARCH_MODULE, flag)

    return Ok
}

var (
    elasticSearchSubscribe = String_t{ len("subscribe"), "subscribe" }
    cluster = String_t{ len("cluster"), "cluster" }
    outputElasticSearch ElasticSearchOutput
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
