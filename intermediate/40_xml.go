package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	City    string   `xml:"city,omitempty"`
	Email   string   `xml:"email"`
	Address Address  `xml:"address"`
	//Email   string   `xml:"-"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	//person := Person{Name: "John", Age: 30, City: "Batum", Email: "john@mail.com"}
	person := Person{Name: "John", Email: "john@mail.com", Address: Address{City: "NY", State: "NY"}}
	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling data to XML:", err)
	}
	fmt.Println("XML data", string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling data to XML:", err)
	}
	fmt.Println(string(xmlData1))

	//xmlRaw := `<person><name>John</name><age>30</age></person>`
	xmlRaw := `<person><name>John</name><age>30</age><address><city>Batum</city><state>Ajara</state></address></person>`
	var personXml Person
	err = xml.Unmarshal([]byte(xmlRaw), &personXml)
	if err != nil {
		log.Fatalln("Error unmarshalling xml:", err)
	}
	fmt.Println(personXml)
	fmt.Println("Local string", personXml.XMLName.Local)
	fmt.Println("Namespace", personXml.XMLName.Space)

	book := Book{
		ISBN:   "23-452-366-778",
		Title:  "Go",
		Author: "Tewrjk",
		Pseudo: "Pseudo",
		PseudoAttr: "Pseudo Attr",
	}

	xmlDataAttr, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling data:", err)
	}
	fmt.Println(string(xmlDataAttr))
}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title"`
	Author     string   `xml:"author"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}
