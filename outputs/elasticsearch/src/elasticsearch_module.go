/*
 * Copyright (C) 2017 Meng Shi
 */

package stdout

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

const (
    ELASTICSEARCH_MODULE = OUTPUT_MODULE|MAIN_MODULE
    ELASTICSEARCH_CONFIG = USER_CONFIG|CONFIG_ARRAY
)

type ElasticSearch struct{
    *Module_t
}

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

var elasticSearchModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    elasticSearchCommands,
    OUTPUT_MODULE,
}

func init() {
    Modules = append(Modules, &ElasticSearch{elasticSearchModule})
}