package main


import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)


type Post struct { //#A
	XMLName xml.Name `xml:"post"`
	Id int    `xml:"id,attr"`
	Content string `xml:"content"`
	Author string `xml:"author"`
	Xml  string `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}


type Comment struct {
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
}


type Author struct {
	Id string `xml:"id,attr"`
	Name string `xml:,chardata"`
}


func main() {
	fileName := "posts.xml"
	fl,err := os.Open(fileName)
	if err != nil {
		panic(fmt.Errorf("Error opening file %s",fileName))
	}
	defer fl.Close()
		
	decoder := xml.NewDecoder(fl)
	for {
		t,err := decoder.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error decoding xml")
			return
		}
		
		switch se := t.(type) {
			case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment,&se)
			}
		}
	}
}
