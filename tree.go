package tree

type Node struct {
	Kind       int
	Value      interface{}
	Parent     *Node
	Prev       *Node
	Next       *Node
	FirstChild *Node
	LastChild  *Node
}

func (n *Node) Children() []*Node {
	children := []*Node{}
	if n.FirstChild == nil {
		return children
	}
	for cur := n.FirstChild; cur != nil; cur = cur.Next {
		children = append(children, cur)
	}
	return children
}

func (n *Node) add(c *Node) {
	c.Parent = n
	if n.LastChild != nil {
		n.LastChild.Next = c
		c.Prev = n.LastChild
		n.LastChild = c
		return
	}
	n.FirstChild = c
	n.LastChild = c
	return
}

// Add adds the children sequentially.
func (n *Node) Add(children ...*Node) {
	for _, child := range children {
		n.add(child)
	}
}

//func (n *node) InsertBefore(ch *node)
