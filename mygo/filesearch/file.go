package filesearch

import (
	"fmt"
)

type File struct {
	fileName string
	firstLine string
	loaded    bool
	rawBody []byte
	body    []string
}

func (f File) String() string {
	var str string
	str = fmt.Sprintf("fileName : %s\n",f.fileName)
	if f.loaded {
		str += "body:\n"
		for _,ln := range f.body {
			str += ln + "\n"
		}
	} else {
		str += "firstLine : " + f.firstLine
	}
	return str
}


type Searcher  interface {
	Search (pattern string)([]File,error)
	Load(*File) error 
}
