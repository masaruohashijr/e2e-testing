package general

import (
	"log"
	"os"
)

func WriteActionXML(strXML string) {
	f, err := os.Create("../../xml/inbound.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(strXML)
	if err2 != nil {
		log.Fatal(err2)
	}
}
