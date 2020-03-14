package study

import (
	"hash/crc32"
	"sort"
)

// 参考实现: https://medium.com/@sent0hil/consistent-hashing-a-guide-go-implementation-fe3421ac3e8f
// GitHub上的一个实现: https://github.com/stathat/consistent

// Node is a single entity in a ring.
type Node struct {
	Id     string
	HashId uint32
}

type Nodes []*Node

// Ring is a network of distributed nodes.
type Ring struct {
	Nodes Nodes
}

// Initializes new distribute network of nodes or a ring.
func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}

// Adds node to the ring.
func (r *Ring) AddNode(id string) {
	r.Nodes = append(r.Nodes, NewNode(id))
	sort.Sort(r.Nodes)
}

// Removes node from the ring if it exists, else returns ErrNodeNotFound.
func (r *Ring) RemoveNode(id string) {}

func (r *Ring) Get(key string) string {
	searchFunc := func(i int) bool {
		return r.Nodes[i].HashId >= crc32.ChecksumIEEE([]byte(key))
	}
	i := sort.Search(len(r.Nodes), searchFunc)
	return r.Nodes[i].Id
}

func NewNode(id string) *Node {
	return &Node{
		Id:     id,
		HashId: crc32.ChecksumIEEE([]byte(id)),
	}
}

func (n Nodes) Len() int {
	return len(n)
}
func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n Nodes) Less(i, j int) bool {
	return n[i].HashId < n[j].HashId
}

func main()  {

}