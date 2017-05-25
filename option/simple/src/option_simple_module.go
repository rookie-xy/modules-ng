/*
 * Copyright (C) 2017 Meng Shi
 */

package simple

import (
    . "github.com/rookie-xy/worker/types"
)

type SimpleOption struct {
    *Module_t
}

func (r *SimpleOption) Init(o *Option_t) int {
    argv := o.GetArgv()

    for i := 1; i < o.GetArgc(); i++ {

        if argv[i][0] != '-' {
            return Error
        }

        switch argv[i][1] {

        case 'c':
	           if argv[i + 1] == "" {
                return Error
            }

            /* file://resource= */
            o.SetItem("configure", argv[i + 1])
            i++
            break

        case 'r':
	           if argv[i + 1] == "" {
	               return Error
	           }

            /* zookeeper://resource= */
            o.SetItem("resource", argv[i + 1])
            i++
            break

        case 'f':
	           if argv[i + 1] == "" {
	               return Error
	           }

            /* yaml, json, xml ... */
            o.SetItem("format", argv[i + 1])
            i++
            break

        case 't':
            o.SetItem("test", true)
	           break

        default:
            o.SetItem("invaild", "")
            o.Info("not found any option")
            break
        }
    }

    configure := NewConfigure(o.Log_t)

    if item := o.GetItem("format"); item != nil {
        name := item.(string)

        for _, codec := range Codecs {
            if codec.Type(name) == Ignore {
                continue
            }

            codec.Init(nil)

            this := NewCodec(codec)
            //code.SetName(item.(string))
            configure.Codec_t = &this
        }
    } else {
        return Error
    }

    o.Configure_t = configure

    return Ok
}

var SimpleOptionModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    nil,
    SYSTEM_MODULE,
}

func init() {
    Modules = Load(Modules, &SimpleOption{SimpleOptionModule})
}
