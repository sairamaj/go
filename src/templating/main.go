package main

import(
	"fmt"
	"text/template"
	"bytes"
)

var DATA string = `
<xml>{{.Count}} items are made of {{.Material}}</xml>
`

func main(){
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse(DATA)
	if err != nil { panic(err) }

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, sweaters)
	
	if err != nil { panic(err) }
	fmt.Println("____________________")
	fmt.Println(tpl.String())
	fmt.Println("____________________")
}