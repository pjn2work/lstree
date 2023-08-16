package lstree

import (
	"fmt"
	"os"
	"regexp"
)

// setup filename filters from argsys
func (ff *fileFilters) initByArgs() {
	if len(os.Args) <= 1 {
		ff.filters = nil
	} else {
		for _, regex := range os.Args[1:] {
			r, err := regexp.Compile(regex)
			if err == nil {
				ff.filters = append(ff.filters, r)
			} else {
				fmt.Printf("Filter %s not allowed because: %s\n", regex, err)
			}
		}
	}
}

// check if filename is valid by the argsys filters
// returns int with the filter number that had match
//
//	-1 no match
//	 0 no filters, every file is valid
//	 1..n filter position on args that had a match
func (ff *fileFilters) isValid(name string) int {
	// if no filter defined then all is valid
	if ff.filters == nil {
		return 0
	}

	// check for any valid filter
	for i, r := range ff.filters {
		if r.MatchString(name) {
			return i + 1
		}
	}

	// name didn't meet any of the required filters
	return -1
}
