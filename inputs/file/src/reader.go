package file

import . "github.com/rookie-xy/worker/types"

type Reader struct {

}

func NewReader() *Reader {
    return &Reader{}
}

func (r *Reader) Next() int {
    return Ok
}
