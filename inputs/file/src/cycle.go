package file

import (
    . "github.com/rookie-xy/worker/types"
    "unsafe"
)

type FileCycle struct {
    Log
}

func NewFileCycle(log Log) *FileCycle {
    object := &FileCycle{
        Log: log,
    }

    return object
}

func (r *FileCycle) Init(configure interface{}) int {
    return Ok
}

func (r *FileCycle) Start(p *unsafe.Pointer) int {
    return Ok
}

func (r *FileCycle) Stop() int {
    return Ok
}
