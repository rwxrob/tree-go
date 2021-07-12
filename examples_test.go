package tree_test

import (
	"fmt"

	"github.com/rwxrob/tree-go"
)

func ExampleNode() {

	r := &tree.Node{Kind: 1, Value: "foo"}
	c := &tree.Node{Kind: 2, Value: "bar"}
	o := &tree.Node{Kind: 2, Value: "other"}
	l := &tree.Node{Kind: 2, Value: "last"}
	r.Add(c, o, l)

	fmt.Printf(`
Root:  %v %v %v %v %v
Child: %v %v %v %v %v
Other: %v %v %v %v %v
Last:  %v %v %v %v %v
Number of Children: %v
First Child: %v
Last Child:  %v
`,
		r.Kind, r.Value, r.Parent, r.Prev, r.Next,
		c.Kind, c.Value, c.Parent.Value, c.Prev, c.Next.Value,
		o.Kind, o.Value, o.Parent.Value, o.Prev.Value, o.Next.Value,
		l.Kind, l.Value, l.Parent.Value, l.Prev.Value, l.Next,
		len(r.Children()), r.FirstChild.Value, r.LastChild.Value,
	)

	// Output:
	// Root:  1 foo <nil> <nil> <nil>
	// Child: 2 bar foo <nil> other
	// Other: 2 other foo bar last
	// Last:  2 last foo other <nil>
	// Number of Children: 3
	// First Child: bar
	// Last Child:  last

}
