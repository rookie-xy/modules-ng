/*
 * Copyright (C) 2017 Meng Shi
 */

package elasticsearch

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)


type ElasticSearch struct{
    *Module_t
}

func NewElasticSearch() *ElasticSearch {
    return &ElasticSearch{}
}

func (r *ElasticSearch) Writer(events Event) int {
    return Ok
}
