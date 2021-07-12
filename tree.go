package tree

// AddChild fulfills the Node interface by adding the specified Node as
// a new child. See NewNode for more about how to create a Node to be
// added.

// Children must return a slice of Nodes even if there are no children.

type Node interface {
	Parent() Node
	Prev() Node
	Next() Node
	ChildFirst() Node
	ChildLast() Node
	AddChild(Node)
	Children() []Node
	InsertBefore(Node)
	InsertAfter(Node)
	Delete()
}

type node struct {
	Kind  int
	Value interface{}

	parent *node
	prev   *node
	next   *node
	first  *node
	last   *node
}

func (n *node) Parent() *node     { return n.parent }
func (n *node) Prev() *node       { return n.prev }
func (n *node) Next() *node       { return n.next }
func (n *node) ChildFirst() *node { return n.first }
func (n *node) ChildLast() *node  { return n.last }

func (n *node) Children() []*node {
	chs := []*node{}
	if n.first == nil {
		return chs
	}
	for cur := n.first; cur != nil; cur = cur.next {
		chs = append(chs, cur)
	}

	return chs
}

// NewNode returns a struct that fulfills the Node, fmt.String, and
// yaml.Marshaller/Unmarshaller interfaces.
func NewNode() *node { return new(node) }

func (n *node) AddChild(ch *node) {
	ch.parent = n
	if n.last != nil {
		n.last.next = ch
		ch.prev = n.last
		n.last = ch
		return
	}
	n.first = ch
	n.last = ch
	return
}

//func (n *node) InsertBefore(ch *node)
