/*
 * Copyright (C) 2017 Meng Shi
 */

package file

import . "github.com/rookie-xy/worker/types"

type FileInput struct {
    *Module_t

     group    string
     types    string
     publish  string
     paths    Array_t
     codec    Codec_t

     Input
     Channel
}

func NewFileInput() *FileInput {
    return &FileInput{}
}

