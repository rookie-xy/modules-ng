/*
 * Copyright (C) 2017 Meng Shi
 */

package stdout

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
"fmt"
)

type OutputStdout struct {
    *Module_t

     subscribe  Array_t
     cluster    Array_t
}

func NewOutputStdout() *OutputStdout {
    return &OutputStdout{}
}

type OutputStdoutContext struct {
    Name   String_t
    Data   [32]*unsafe.Pointer
}

var stdoutOutput = String_t{ len("stdout_output"), "stdout_output" }
var outputStdoutContext = &OutputStdoutContext{
    Name: stdoutOutput,
}

func (r *OutputStdoutContext) Set() unsafe.Pointer {
    stdout := NewOutputStdout()
    if stdout == nil {
        return nil
    }

    //stdout.subscribe = "zhang yue"

    return unsafe.Pointer(stdout)
}

func (r *OutputStdoutContext) Get() []*unsafe.Pointer {
    return r.Data[:]
}

var (
    stdoutSubscribe = String_t{ len("subscribe"), "subscribe" }
    stdoutCluster = String_t{ len("cluster"), "cluster" }
    outputStdout OutputStdout
)

var outputStdoutCommands = []Command_t{

    { stdoutSubscribe,
      STDOUT_CONFIG|CONFIG_ARRAY,
      SetArray,
      0,
      unsafe.Offsetof(outputStdout.subscribe),
      nil },

    { stdoutCluster,
      STDOUT_CONFIG|CONFIG_ARRAY,
      SetArray,
      0,
      unsafe.Offsetof(outputStdout.cluster),
      nil },

    NilCommand,
}

func (r *OutputStdout) Init(o *Option_t) int {
    context := r.Context.Get()

    for _, v := range context {
        if v != nil {
            this := (*OutputStdout)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            for i := 0; i < this.subscribe.GetLength(); i++ {
                fmt.Println(this.subscribe.GetData(i))
            }

            for i := 0; i < this.cluster.GetLength(); i++ {
                fmt.Println(this.cluster.GetData(i))
            }

            //fmt.Println(i, this.subscribe, this.subscribe)
        } else {
            break
        }
    }

    return Ok
}

func (r *OutputStdout) Main(cfg *Configure_t) int {
    fmt.Println("output main")
    return Ok
}

func (r *OutputStdout) Exit() int {
    fmt.Println("output exit")
    return Ok
}

var outputStdoutModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    outputStdoutContext,
    outputStdoutCommands,
    STDOUT_MODULE,
}

func init() {
    Modules = append(Modules, &OutputStdout{Module_t:outputStdoutModule})
}
