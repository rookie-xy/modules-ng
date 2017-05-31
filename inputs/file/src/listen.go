package file

import . "github.com/rookie-xy/worker/types"

type Listen struct{

}

func NewListen() *Listen {
    return &Listen{}
}

func (r *Listen) Accept() int {
    return Ok
}
