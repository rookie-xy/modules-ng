/*
 * Copyright (C) 2017 Meng Shi
 */

package elasticsearch

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
"fmt"
)

type ElasticSearchOutput struct {
    *Module_t

     subscribe  Array_t
     cluster    Array_t

     Output
}

func NewElasticSearchOutput() *ElasticSearchOutput {
    return &ElasticSearchOutput{}
}
