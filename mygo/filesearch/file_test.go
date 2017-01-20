package filesearch

import (
	"testing"
)

func TestFileStringNotLoaded(t *testing.T) {
	fl := File{fileName : "test.txt", firstLine: "Hello,World", loaded : false,}
	str := fl.String()
	var testStr = "fileName : test.txt\nfirstLine : Hello,World" 
	if str != testStr  {
		t.Errorf("Expected  %s : Actual %s",testStr,str)
	}
}

func TestFileStringLoaded (t* testing.T) {
	fl := File{fileName : "test.txt", firstLine: "Hello", loaded : true,
		body : [] string {"Hello", "World",},}
	str := fl.String()
	var testStr = "fileName : test.txt\nbody:\nHello\nWorld\n"
	if str != testStr  {
		t.Errorf("Expected  %s : Actual %s",testStr,str)
	}
}
