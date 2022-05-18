package lstree

import (
	"fmt"
	"io/ioutil"
	"sort"
)

const ScreenFileSpacing = 110

func (fdr *FileData) searchForFiles(baseFolder string, ff *fileFilters) listFileData {
    files, _ := ioutil.ReadDir(baseFolder)
    for _, f := range files {

        fd := FileData {
            fileName:    f.Name(),
            isDir:       f.IsDir(),
            size:        f.Size(),
            modTime:     f.ModTime(),
            color:       ColorFolder,
            subFileData: nil,
        }

        if f.IsDir() {
            fd.subFileData = fd.searchForFiles(baseFolder+"/"+f.Name(), ff)
            if fd.subFileData != nil {
                fdr.subFileData = append(fdr.subFileData, fd)
            }
        } else {
            filterMatchPos := ff.isValid(f.Name())
            if filterMatchPos >= 0 {
                fd.color = ColorFileFiltered[filterMatchPos]    // TODO % ColorFileFiltered
                fdr.subFileData = append(fdr.subFileData, fd)
            }
        }
    }

    // sort by name
    sort.Slice(fdr.subFileData, func(i, j int) bool {
        return fdr.subFileData[i].fileName < fdr.subFileData[j].fileName
    })

    return fdr.subFileData
}

func (fdr *FileData) printListDir(prev string) {
    total := len(fdr.subFileData) - 1
    next, curr := prev+"│   ", "├── "
    for i, fd := range fdr.subFileData {
        if i == total {
            next = prev + "    "
            curr = "└── "
        }

        if fd.isDir {
            fmt.Printf("%s%s%s%s%s\n", prev, curr, fd.color, fd.fileName, Reset)
            fd.printListDir(next)
        } else {
            output := fmt.Sprintf("%s%s%s%s%s", prev, curr, fd.color, fd.fileName, Reset)
            fillN := ScreenFileSpacing - getStringLen(output)
            fmt.Printf("%s %s | %s | %9d\n", output, fillWith("-", fillN), fd.modTime, fd.size)
        }
    }
}

func ListDirWithFilters(baseFolder string) {
    // setup file filters
    ff := fileFilters{}
    ff.initByArgs()

    // search for files and subfolders
    root := FileData{
        fileName: baseFolder,
        isDir: true,
        color: ColorFolder,        
    }
    root.subFileData = root.searchForFiles(baseFolder, &ff)

    // print tree results
    root.printListDir("")
}
