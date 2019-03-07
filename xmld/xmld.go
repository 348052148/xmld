package xmld

import (
	"encoding/xml"
	"io"
)

type Document struct {
	root *Element
}

type Element struct {
	Name string
	child   []*Element
	Attr	[]Attr
	Data   string
}

type Attr struct {
	Name string
	Value string
}

func NewElement(name string, attrs []xml.Attr) *Element  {
	var attrList []Attr
	for _,attr := range attrs {
		attrList = append(attrList,Attr{
			//区分Space Local
			Name:attr.Name.Local,
			Value:attr.Value,
		})
	}
	return &Element{
		Name:name,
		Attr:attrList,
	}
}

func (e *Element)AddData(value string)  {
	e.Data = value
}

func (e *Element)AddChild(node *Element)  {
	e.child = append(e.child, node)
}

func (e *Element)GetChildAllByTagName(name string) []*Element  {
	var elems []*Element
	for _,elem := range e.child {
		if elem.Name == name {
			elems = append(elems, elem)
		}
	}
	return elems
}

func (e *Element)GetChildByTagName(name string) *Element  {
	for _,elem := range e.child {
		if elem.Name == name {
			return elem
		}
	}
	return nil
}

func NewDocument() *Document {
	return &Document{
		root:nil,
	}
}

func (doc *Document)GetRoot() *Element  {
	return doc.root
}

func (doc *Document)Read(reader io.Reader)  {
	dec := xml.NewDecoder(reader)
	var stack stack
	for {
		t, err := dec.RawToken()
		if err == io.EOF || err!=nil {
			break
		}
		//stack.push(doc.root)
		top := stack.peek()
		switch t := t.(type) {
		case xml.StartElement:
			e := NewElement(t.Name.Local, t.Attr)
			if (doc.root == nil) {
				doc.root = e
			}else {
				top.(*Element).AddChild(e)
			}
			stack.push(e)
		case xml.EndElement:
			stack.pop()
		case xml.CharData:
			if e,ok := top.(*Element); ok {
				e.AddData(string(t))
			}
		case xml.Comment:
			//fmt.Println(t)
		case xml.Directive:
			//fmt.Println(t)
		case xml.ProcInst:
			//fmt.Println(t)
		}
	}
}