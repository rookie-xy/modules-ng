/*
 * Copyright (C) 2017 Meng Shi
 */

package memory

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
"fmt"
)

type ChannelMemory struct {
    *Module_t

     name     string
     size     int
}

func NewChannelMemory() *ChannelMemory {
    return &ChannelMemory{}
}

type ChannelMemoryContext struct {
    Name   String
    Data   [16]*unsafe.Pointer
}

var memoryChannel = String_t{ len("memory_channel"), "memory_channel" }
var channelMemoryContext = &ChannelMemoryContext{
    Name: memoryChannel,
}

func (r *ChannelMemoryContext) Create() unsafe.Pointer {
    memory := NewChannelMemory()
    if memory == nil {
        return nil
    }

    memory.name = "channel name"
    memory.size = 1024

    return unsafe.Pointer(memory)
}

func (r *ChannelMemoryContext) GetDatas() []*unsafe.Pointer {
    return r.Data[:]
}

var (
    name = String_t{ len("name"), "name" }
    size = String_t{ len("size"), "size" }
    channelMemory ChannelMemory
)

var channelMemoryCommands = []Command_t{

    { name,
      MEMORY_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(channelMemory.name),
      nil },

    { size,
      MEMORY_CONFIG|CONFIG_VALUE,
      SetNumber,
      0,
      unsafe.Offsetof(channelMemory.size),
      nil },

    NilCommand,
}


func (r *ChannelMemory) Init(o *Option_t) int {
    context := r.Context.GetDatas()

    for _, v := range context {
        if v != nil {
            this := (*ChannelMemory)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            fmt.Println(this.name, this.size)
        } else {
            break
        }
    }

    return Ok
}

func (r *ChannelMemory) Main(c *Configure_t) int {
    fmt.Println("Memory main")
    return Ok
}

func (r *ChannelMemory) Exit() int {
    fmt.Println("Memory exit")
    return Ok
}

func (r *ChannelMemory) Type() *Module_t {
    return r.Self()
}

var channelMemoryModule = &Module{
    MODULE_V1,
    CONTEXT_V1,
    channelMemoryContext,
    channelMemoryCommands,
    MEMORY_MODULE,
}

func init() {
    Modules = Load(Modules, &ChannelMemory{Module_t:channelMemoryModule})
}