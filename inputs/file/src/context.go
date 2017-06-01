package file

import (
      "unsafe"
    . "github.com/rookie-xy/worker/types"
)

type FileInputContext struct {
    *Context_t

     Name     String_t
     Data     [32]*unsafe.Pointer
}

var fileInputContext = &FileInputContext{
    Name:      String_t{ len("file_input"), "file_input" },
    Context_t: NewContext(),
}

func (r *FileInputContext) Build() unsafe.Pointer {
    file := NewFileInput()
    if file == nil {
        return nil
    }

    file.group   = "worker"
    file.types   = "file"
    file.publish = "topic"

    if this := CreateBackgroundContext(); this != nil {
        if r.WithCancel(this) != Ok {
            return nil
        }
    }

    return unsafe.Pointer(file)
}

func (r *FileInputContext) Configure() []*unsafe.Pointer {
    return r.Data[:]
}
