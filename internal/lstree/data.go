package lstree

import (
	"regexp"
	"time"
)

const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
    Gray   = "\033[37m"
    White  = "\033[97m"
)

var (
    ColorFileFiltered = [...]string{Gray, Green, Yellow, Blue, Purple, Cyan, White}
    ColorFolder = Red
)


type fileFilters struct {
    filters []*regexp.Regexp
}

type FileData struct {
    fileName    string
    isDir       bool
    size        int64
    modTime     time.Time
    color       string
    subFileData listFileData
}

type listFileData []FileData