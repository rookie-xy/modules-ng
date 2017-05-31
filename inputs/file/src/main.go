package file

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
"fmt"
)

func (r *FileInput) Init(o *Option_t) int {

    if configure := o.Configure_t; configure != nil {
        for _, v := range r.Context.Configure() {
            if v != nil {
                this := (*FileInput)(unsafe.Pointer(uintptr(*v)))
                if this == nil {
                    return Error
                }

                this.Input = InitFile()

                //fmt.Println(this.types, this.group, this.publish)
                for _, v := range configure.Channels {
                    // support topic m:n
                    if v.Type(this.publish) == Ok {
                        this.Channel = v.New()
                        break;
                    }
                }

                for i := 0; i < this.paths.GetLength(); i++ {
                    fmt.Println(this.paths.GetData(i))
                }

                //codec := this.codec.New()
                //codec.Encode(nil)

            } else {
                break
            }
        }
    }

    return Ok
}

func (r *FileInput) Main(c *Configure_t) int {
    context := r.Context.Self()
    if context == nil {
        return Error
    }

    var cycle Cycle

    cycle = NewFileCycle(c.Log)
    if cycle == nil {
        return Error
    }

    cycle.Init(nil)

    if configures := r.Context.Configure(); configures != nil {
        for _, configure := range configures {
            cycle.Start(configure, context)
        }
    }

    select {

    case context.Done():
        cycle.Stop()
    }

    fmt.Println("File main")
    return Ok
}

func (r *FileInput) Exit() int {
    if routine := r.Context.Self(); routine != nil {
        routine.Kill()
    }

    fmt.Println("File exit")
    return Ok
}
