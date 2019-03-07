package main

import (
	"fmt"
	"strings"
	"test/xmlcode/xmld"
)


func main() {
	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	doc := xmld.NewDocument()
	doc.Read(strings.NewReader(data))
	fmt.Println(doc.GetRoot().GetChildByTagName("Email").
		GetChildByTagName("Addr").Data)
}
