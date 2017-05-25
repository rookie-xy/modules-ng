/*
 * Copyright (C) 2017 Meng Shi
 */

package file

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
"fmt"
)

type InputFile struct {
    *Module_t

     group    string
     types    string
     publish  string
     paths    Array_t
     codec    Codec_t

     Channel
}

func NewInputFile() *InputFile {
    return &InputFile{}
}

type InputFileContext struct {
    Name   String_t
    Data   [32]*unsafe.Pointer
}

var fileInput = String_t{ len("file_input"), "file_input" }
var inputFileContext = &InputFileContext{
    Name: fileInput,
}

func (r *InputFileContext) Create() unsafe.Pointer {
    file := NewInputFile()
    if file == nil {
        return nil
    }

    file.group   = "worker"
    file.types   = "file"
    file.publish = "topic"
    //file.paths   = nil
    //file.codec   = nil

    return unsafe.Pointer(file)
}

func (r *InputFileContext) GetDatas() []*unsafe.Pointer {
    return r.Data[:]
}

var (
    group   = String_t{ len("group"), "group" }
    types   = String_t{ len("type"), "type" }
    paths   = String_t{ len("paths"), "paths" }
    publish = String_t{ len("publish"), "publish" }
    codec   = String_t{ len("codec"), "codec" }

    inputFile InputFile
)

var inputFileCommands = []Command_t{

    { group,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(inputFile.group),
      nil },

    { types,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(inputFile.types),
      nil },

    { paths,
      FILE_CONFIG|CONFIG_ARRAY,
      SetArray,
      0,
      unsafe.Offsetof(inputFile.paths),
      nil },

    { publish,
      FILE_CONFIG|CONFIG_VALUE,
      SetString,
      0,
      unsafe.Offsetof(inputFile.publish),
      nil },

    { codec,
      FILE_CONFIG|CONFIG_BLOCK,
      SetCodec,
      0,
      unsafe.Offsetof(inputFile.codec),
      nil },

    NilCommand,
}

func (r *InputFile) Init(o *Option_t) int {
    configure := o.Configure_t
    if configure == nil {
        return Error
    }

    context := r.Context.GetDatas()
    for _, v := range context {
        if v != nil {
            this := (*InputFile)(unsafe.Pointer(uintptr(*v)))
            if this == nil {
                return Error
            }

            fmt.Println(this.types, this.group, this.publish)
            for _, v := range configure.Channels {
                // support topic m:n
                if v.Init(this.publish) == Ok {
                    this.Channel = v.New()
                    break;
                }
            }

            /*
            for i := 0; i < this.paths.GetLength(); i++ {
                fmt.Println(this.paths.GetData(i))
            }

            codec := this.codec.New()
            codec.Encode(nil)
            */
        } else {
            break
        }
    }

    return Ok
}

func (r *InputFile) Main(cfg *Configure_t) int {
    fmt.Println("File main")
    return Ok
}

func (r *InputFile) Exit() int {
    fmt.Println("File exit")
    return Ok
}

var inputFileModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    inputFileContext,
    inputFileCommands,
    FILE_MODULE,
}

func init() {
    Modules = Load(Modules, &InputFile{Module_t:inputFileModule})
}
