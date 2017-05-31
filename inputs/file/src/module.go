package file

import . "github.com/rookie-xy/worker/types"

var fileModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    fileCommands,
    INPUT_MODULE,
}

var fileInputModule = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    fileInputContext,
    fileInputCommands,
    FILE_MODULE,
}

func init() {
    Modules = append(Modules, &File{Module_t:fileModule},
                              &FileInput{Module_t:fileInputModule})
}
