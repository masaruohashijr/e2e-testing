package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func WriteActionXML(xmlName, strXML string) {
	err := ioutil.WriteFile("xml/"+xmlName+".xml", []byte(strXML), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func AppendActionXML(actionUrl, appendix string) {
	slice := strings.Split(strings.ToLower(actionUrl), "/")
	xmlName := slice[len(slice)-1]
	bXML, err2 := ioutil.ReadFile("xml/" + xmlName + ".xml")
	if err2 != nil {
		log.Fatal(err2)
	}
	strXML := string(bXML)
	index := strings.Index(strXML, "</Response>")
	strXML = strXML[:index] + appendix + strXML[index:]
	fmt.Println(strXML)
	WriteActionXML(xmlName, strXML)
}
