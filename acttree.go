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

func (n *Node) Next() (kid *Node, kidSiblings []*Node) {
	var k []*Node
	if n.firstKid != nil {
		for x := n.firstKid.nextSibling; x != nil; x = x.nextSibling {
			k = append(k, x)
		}
	}
	return n.firstKid, k
}

type Tree struct {
	Root *Node //first node of the tree
}

func New() *Tree {
	return new(Tree)
}

// AddNode adds a new kid or sibling node to the tree.
// The new node will be added after the param node.
// When the param node has no kid, new node is added as kid,
// or it is added as a sibling.
// Note that the new node can be the root of another tree,
// making it a merge(join) function of trees.
func (t *Tree) AddNode(n, prev *Node) (added *Node) {
	if t.Root == nil { //add first node to tree
		t.Root = n
		return n
	}

	if prev.firstKid == nil {
		prev.firstKid = n
		n.parent = prev
		return n
	}

	for x.nextSibling != nil {
		x = x.nextSibling
	}

	x.nextSibling = n
	n.prevSibling = x
	n.parent = x.parent
	return n
}

// Add is a convenience wrapper for AddNode(&Node{Value: v}, prev)
func (t *Tree) Add(v interface{}, prev *Node) (added *Node) {
	return t.AddNode(&Node{Value: v}, prev)
}

// Cut cuts a subtree off the original, with the given node
// being root of the new tree.
func (t *Tree) Cut(n *Node) (subTree *Tree) {
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
		} else { //x is sibling
			prev.nextSibling = n.nextSibling
			if n.nextSibling != nil {
				n.nextSibling.prevSibling = prev
			}
		}
	}

	n.parent = nil
	n.prevSibling = nil
	n.nextSibling = nil

	var nt = new(Tree)
	nt.Root = n
	return nt
}

type WalkThroughMethod int

const (
	FirstKidOnly WalkThroughMethod = iota
	DepthFirst
	BreathFirst
)

// WalkThrough goes over nodes in tree by the order defined in WalkThroughMethod.
// It will call the user-defined node handler enter() and leave() on each node.
// enter() is called before all its later nodes called in WalkThroughMethod order.
// leave() is called after all its later nodes called.
// enter() and leave() must return false to keep the WalkThrough loop continue.
// Returns true means to stop visiting the other nodes, job done.
func (t *Tree) WalkThrough(n *Node, m WalkThroughMethod,
	enter func(*Node) bool, leave func(*Node) bool) {

	var stopWalk = false
	if stopWalk {
		return
	}

	s := n //starting point
	if s == nil {
		s = t.Root
	}
	if s == nil {
		return
	} //empty tree

	if enter != nil {
		stopWalk = enter(s)
		if stopWalk {
			return
		}
	}

	switch m {
	case DepthFirst:
		if s.firstKid != nil {
			t.WalkThrough(s.firstKid, DepthFirst, enter, leave)
			if stopWalk {
				return
			}
		}
		if s.nextSibling != nil {
			t.WalkThrough(s.nextSibling, DepthFirst, enter, leave)
			if stopWalk {
				return
			}
		}
	case FirstKidOnly:
		if s.firstKid != nil {
			t.WalkThrough(s.firstKid, FirstKidOnly, enter, leave)
			if stopWalk {
				return
			}
		}
	case BreathFirst:
		if s.nextSibling != nil {
			t.WalkThrough(s.nextSibling, BreathFirst, enter, leave)
			if stopWalk {
				return
			}
		}
		if s.firstKid != nil {
			t.WalkThrough(s.firstKid, BreathFirst, enter, leave)
			if stopWalk {
				return
			}
		}
	}

	if leave != nil {
		stopWalk = leave(s)
		if stopWalk {
			return
		}
	}
}

// Remove deletes a node and all its kids from tree.
func (t *Tree) Remove(n *Node) int {
	st := t.Cut(n)
	if st == nil {
		return 0
	}

	c := 0
	st.WalkThrough(nil, DepthFirst, nil, func(x *Node) bool {
		c++

		//clean up for preventing memory leak
		x.parent = nil
		x.prevSibling = nil
		x.nextSibling = nil
		x.firstKid = nil
		x.Value = nil

		return false
	})

	return c
}
