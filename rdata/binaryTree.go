/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-18 14:52
 */
package rdata

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (n *Node) Print() {
	fmt.Println(n.Value)
}

func (n *Node) AddValue(value int) {
	if n == nil {
		fmt.Println("the node is nil")
		return
	}
	n.Value = value
}

//前序遍历 根-左-右
func (n *Node) PreOrder() {

	if n == nil {
		return
	}
	n.Print()
	n.Left.PreOrder()
	n.Right.PreOrder()
}

//中序遍历 左-根-右
func (n *Node) MiddleOrder() {
	if n == nil {
		return
	}
	n.Left.MiddleOrder()
	n.Print()
	n.Right.MiddleOrder()
}

//后序遍历 左右根
func (n *Node) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	n.Print()
}
