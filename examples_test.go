package tree_test

import (
	"fmt"

	"github.com/rwxrob/tree-go"
)

func ExampleNewNode() {

	node := tree.NewNode()
	node.Kind = 2
	node.Value = "something"

	fmt.Printf(
		"Type: %T\nKind: %v\nValue: %v",
		node, node.Kind, node.Value,
	)

	// Output:
	// Type: *tree.node
	// Kind: 2
	// Value: something

}

func ExampleNode() {

	r := tree.NewNode()
	r.Kind = 1
	r.Value = "foo"

	c := tree.NewNode()
	c.Kind = 2
	c.Value = "bar"

	o := tree.NewNode()
	o.Kind = 2
	o.Value = "other"

	l := tree.NewNode()
	l.Kind = 2
	l.Value = "last"

	r.AddChild(c)
	r.AddChild(o)
	r.AddChild(l)

	newFirst := tree.NewNode()
	newFirst.Kind = 3
	newFirst.Value = "FIRST"

	fmt.Printf(`
Root:  %T %v %v %v %v %v
Child: %T %v %v %v %v %v
Other: %T %v %v %v %v %v
Last:  %T %v %v %v %v %v
Number of Children: %v
First Child: %v
Last Child:  %v
`,
		r, r.Kind, r.Value, r.Parent(), r.Prev(), r.Next(),
		c, c.Kind, c.Value, c.Parent().Value, c.Prev(), c.Next().Value,
		o, o.Kind, o.Value, o.Parent().Value, o.Prev().Value, o.Next().Value,
		l, l.Kind, l.Value, l.Parent().Value, l.Prev().Value, l.Next(),
		len(r.Children()), r.ChildFirst().Value, r.ChildLast().Value,
	)

	// Output:
	// Root:  *tree.node 1 foo <nil> <nil> <nil>
	// Child: *tree.node 2 bar foo <nil> other
	// Other: *tree.node 2 other foo bar last
	// Last:  *tree.node 2 last foo other <nil>
	// Number of Children: 3
	// First Child: bar
	// Last Child:  last

}
