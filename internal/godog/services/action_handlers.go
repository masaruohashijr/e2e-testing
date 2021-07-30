package services

import (
	"log"
	"os"
)

func WriteActionXML(xmlName, strXML string) {
	f, err := os.Create("xml/" + xmlName + ".xml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(strXML)
	if err2 != nil {
		log.Fatal(err2)
	}
}
