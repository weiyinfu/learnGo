package main

/*
复制双向链表
*/
type Node struct {
	value int
	prev  *Node
	next  *Node
}

func copyLink(node *Node) *Node {
	no := Node{}
	i := node
	j := &no
	for i != nil {
		j.next = &Node{value: i.value, prev: j}
		j = j.next
		i = i.next
	}
	x := no.next
	x.prev = nil
	return x
}
func main() {

}
