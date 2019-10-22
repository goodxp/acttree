// Copyright 2019 goodxp(goodxp@gmail.com). All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package actTree implements a double-linked tree data structure.
// Features: (tiny footprint intended)
//  - create a new tree: actTree.New()
//  - add a node or merge a subtree: Add(), AddNode()
//  - go over all nodes by preferred order: WalkThrough()
//  - goto a neighbor node: Prev(), Next()
//  - split the tree at a node being the root of a subtree: Cut()
//  - remove a node and its kids: Remove()
package actTree

type Node struct {
	parent      *Node
	firstKid    *Node
	prevSibling *Node
	nextSibling *Node

	Value interface{} //user-defined content value to associate with the node
}

func (n *Node) Prev() (parent *Node, prevSibling *Node) {
	return n.parent, n.prevSibling
}

func (n *Node) Next() (kid *Node, nextSibling *Node) {
	return n.firstKid, n.nextSibling
}

func (n *Node) Add(v interface{}) (added *Node) {
	return n.addNode(&Node{Value: v}, false)
}

func (n *Node) addNode(ret *Node, asRootSib bool) (added *Node) {
	if n.firstKid == nil {
		n.firstKid = ret
		ret.parent = n
		return ret
	}

	x := n.firstKid
	if asRootSib {
		x = n
	}

	for x.nextSibling != nil {
		x = x.nextSibling
	}
	x.nextSibling = ret
	ret.prevSibling = x
	ret.parent = x.parent
	return ret
}

func (n *Node) HasSib() bool {
	if n.prevSibling != nil || n.nextSibling != nil {
		return true
	}
	return false
}

type Tree struct {
	Root *Node //first node of the tree
}

func New() *Tree {
	return new(Tree)
}

// AddNode adds a new kid or sibling node to the tree.
// The new node will be added after the prev node.
// When the prev node has no kid, new node is added as firstKid,
// or it is added as a sibling of firstKid.
// Note that the new node can be the root of another tree,
// making it a merge(join) function of trees.
func (t *Tree) AddNode(n, prev *Node) (added *Node) {
	if t.Root == nil { //add first node to tree
		t.Root = n
		return n
	}
	if prev == nil {
		return t.Root.addNode(n, true)
	}
	return prev.addNode(n, false)
}

// Add is a convenience wrapper for AddNode(&Node{Value: v}, prev)
func (t *Tree) Add(v interface{}, prev *Node) (added *Node) {
	return t.AddNode(&Node{Value: v}, prev)
}

// Remove deletes a node and all its kids from tree.
func (t *Tree) Remove(n *Node) int {
	r := t.cut(n)
	if r == nil {
		return 0
	}

	c := 0

	for r.nextSibling != nil {
		c += t.Remove(r.nextSibling)
	}

	for r.firstKid != nil {
		c += t.Remove(r.firstKid)
	}

	//clean up for preventing memory leak
	n.parent = nil
	n.prevSibling = nil
	n.nextSibling = nil
	n.firstKid = nil
	n.Value = nil

	n = nil
	c++

	return c
}

// Cut cuts a subtree off the original, with the given node
// being root of the new tree.
func (t *Tree) Cut(n *Node) (subTree *Tree) {
	nr := t.cut(n)
	if nr != nil {
		var nt = new(Tree)
		nt.Root = nr
		return nt
	}
	return nil
}

func (t *Tree) cut(n *Node) *Node {
	if (n == nil) || (t.Root == nil) {
		return nil
	}

	prev := n.prevSibling
	isFirstKid := false
	if prev == nil {
		prev = n.parent
		isFirstKid = true
	}

	if prev == nil {
		t.Root = nil
	} else {
		if isFirstKid {
			prev.firstKid = n.nextSibling
			if n.nextSibling != nil {
				n.nextSibling.prevSibling = nil
			}
		} else {
			prev.nextSibling = n.nextSibling
			if n.nextSibling != nil {
				n.nextSibling.prevSibling = prev
			}
		}
	}

	n.parent = nil
	n.prevSibling = nil
	n.nextSibling = nil

	return n
}

// WalkThrough iterates over "n" and its kid nodes.
// enter() and leave() called before and after visiting all kid nodes.
// return from enter() and leave(): true = end WalkThrough; false = continue.
func WalkThrough(n *Node, enter func(*Node) bool, leave func(*Node) bool) {
	if n == nil {
		return
	}

	if walk(n, enter, leave) {
		return
	}

	if n.prevSibling == nil && n.parent == nil { //t.Root
		for n_ := n.nextSibling; n_ != nil; n_ = n_.nextSibling {
			if walk(n_, enter, leave) {
				return
			}
		}
	}
}

func walk(n *Node, enter func(*Node) bool, leave func(*Node) bool) bool {
	if enter != nil {
		if enter(n) {
			return true
		}
	}

	for n_ := n.firstKid; n_ != nil; n_ = n_.nextSibling {
		if walk(n_, enter, leave) {
			return true
		}
	}

	if leave != nil {
		if leave(n) {
			return true
		}
	}

	return false
}
