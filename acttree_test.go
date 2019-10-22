package actTree

import "testing"

type v struct {
	x int
}

func TestAdd(t *testing.T) {
	var tree *Tree
	tree = New()

	v1 := v{x: 1}
	v2 := v{x: 2}
	v3 := v{x: 3}

	v4 := v{x: 211}
	v5 := v{x: 212}

	v6 := v{x: 221}
	v7 := v{x: 222}

	n1 := tree.Add(v1, nil)
	n2 := tree.Add(v2, n1)
	tree.Add(v3, n2)

	n4 := tree.Add(v4, n1)
	tree.Add(v5, n4)
	n6 := tree.Add(v6, n1)
	tree.Add(v7, n6)

	got := 0
	want := 7

	WalkThrough(tree.Root, func(n *Node) bool {
		got++

		p, _ := n.Prev()
		if n.parent != p {
			t.Errorf("parent not retrieved %v", n)
		}

		switch n.Value.(v).x {
		case 1:
			if tree.Root != n {
				t.Errorf("root %v shall be %v", tree.Root, n)
			}
			if n.parent != nil {
				t.Errorf("root has no parent %v", n)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 2 {
					t.Errorf("firstKid of %v shall be 2", n)
				}
			}
		case 2:
			if n.parent.Value.(v).x != 1 {
				t.Errorf("parent of %v shall be 1", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling == nil {
				t.Errorf("%v shall have 1 nextSibling", n)
			} else {
				if n.nextSibling.Value.(v).x != 211 {
					t.Errorf("nextSibling of %v shall be 211", n)
				}
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 3 {
					t.Errorf("firstKid of %v shall be 3", n)
				}
			}
		case 3:
			if n.parent.Value.(v).x != 2 {
				t.Errorf("parent of %v shall be 2", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		case 4:
			if n.parent.Value.(v).x != 1 {
				t.Errorf("parent of %v shall be 1", n.parent.Value.(v).x)
			}
			if n.prevSibling == nil {
				t.Errorf("%v shall have 1 prevSibling", n)
			} else {
				if n.prevSibling.Value.(v).x != 2 {
					t.Errorf("prevSibling of %v shall be 2", n)
				}
			}
			if n.nextSibling == nil {
				t.Errorf("%v shall have 1 nextSibling", n)
			} else {
				if n.nextSibling.Value.(v).x != 221 {
					t.Errorf("nextSibling of %v shall be 221", n)
				}
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 212 {
					t.Errorf("firstKid of %v shall be 212", n)
				}
			}
		case 5:
			if n.parent.Value.(v).x != 211 {
				t.Errorf("parent of %v shall be 211", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		case 6:
			if n.parent.Value.(v).x != 1 {
				t.Errorf("parent of %v shall be 1", n.parent.Value.(v).x)
			}
			if n.prevSibling == nil {
				t.Errorf("%v shall have 1 prevSibling", n)
			} else {
				if n.prevSibling.Value.(v).x != 211 {
					t.Errorf("prevSibling of %v shall be 211", n)
				}
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 222 {
					t.Errorf("firstKid of %v shall be 222", n)
				}
			}
		case 7:
			if n.parent.Value.(v).x != 221 {
				t.Errorf("parent of %v shall be 221", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		}
		return false
	}, nil)

	if got != want {
		t.Errorf("added %v nodes got %v", want, got)
	}
}

func TestWalkThrough(t *testing.T) {
	var tree *Tree
	tree = New()

	v1 := v{x: 1}
	v2 := v{x: 2}
	v3 := v{x: 3}

	v4 := v{x: 211}
	v5 := v{x: 212}

	n1 := tree.Add(v1, nil)
	tree.Add(v2, n1)
	tree.Add(v3, n1)
	n4 := tree.Add(v4, n1)
	tree.Add(v5, n4)

	got := 0
	want := 4

	WalkThrough(tree.Root, func(n *Node) bool {
		got++
		if n.Value.(v).x == 211 {
			return true
		}
		return false
	}, nil)

	if got != want {
		t.Errorf("shall stop on %v but on %v", want, got)
	}
}

func TestCut(t *testing.T) {
	var st *Tree
	st = New()
	st.Cut(st.Add(v{x: 1004},
		st.Add(v{x: 1003},
			st.Add(v{x: 1002},
				st.Add(v{x: 1001}, nil)))))

	var tree *Tree
	tree = New()

	v1 := v{x: 1}
	v2 := v{x: 2}
	v3 := v{x: 3}

	v4 := v{x: 211}
	v5 := v{x: 212}

	v6 := v{x: 221}
	v7 := v{x: 222}

	n1 := tree.Add(v1, nil)
	n2 := tree.Add(v2, n1)
	tree.Add(v3, n2)

	n4 := tree.Add(v4, n1)
	tree.Add(v5, n4)

	tree.AddNode(st.Root, n1)

	n6 := tree.Add(v6, n1)
	tree.Add(v7, n6)

	stnew := tree.Cut(st.Root) //for testing Cut()

	got := 0
	want := 7

	WalkThrough(tree.Root, func(n *Node) bool {
		got++

		p, _ := n.Prev()
		if n.parent != p {
			t.Errorf("parent not retrieved %v", n)
		}

		switch n.Value.(v).x {
		case 1:
			if tree.Root != n {
				t.Errorf("root %v shall be %v", tree.Root, n)
			}
			if n.parent != nil {
				t.Errorf("root has no parent %v", n)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 2 {
					t.Errorf("firstKid of %v shall be 2", n)
				}
			}
		case 2:
			if n.parent.Value.(v).x != 1 {
				t.Errorf("parent of %v shall be 1", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling == nil {
				t.Errorf("%v shall have 1 nextSibling", n)
			} else {
				if n.nextSibling.Value.(v).x != 211 {
					t.Errorf("nextSibling of %v shall be 211", n)
				}
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 3 {
					t.Errorf("firstKid of %v shall be 3", n)
				}
			}
		case 3:
			if n.parent.Value.(v).x != 2 {
				t.Errorf("parent of %v shall be 2", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		case 4:
			if n.parent.Value.(v).x != 1 {
				t.Errorf("parent of %v shall be 1", n.parent.Value.(v).x)
			}
			if n.prevSibling == nil {
				t.Errorf("%v shall have 1 prevSibling", n)
			} else {
				if n.prevSibling.Value.(v).x != 2 {
					t.Errorf("prevSibling of %v shall be 2", n)
				}
			}
			if n.nextSibling == nil {
				t.Errorf("%v shall have 1 nextSibling", n)
			} else {
				if n.nextSibling.Value.(v).x != 221 {
					t.Errorf("nextSibling of %v shall be 221", n)
				}
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 212 {
					t.Errorf("firstKid of %v shall be 212", n)
				}
			}
		case 5:
			if n.parent.Value.(v).x != 211 {
				t.Errorf("parent of %v shall be 211", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		case 6:
			if n.parent.Value.(v).x != 1 {
				t.Errorf("parent of %v shall be 1", n.parent.Value.(v).x)
			}
			if n.prevSibling == nil {
				t.Errorf("%v shall have 1 prevSibling", n)
			} else {
				if n.prevSibling.Value.(v).x != 211 {
					t.Errorf("prevSibling of %v shall be 211", n)
				}
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 222 {
					t.Errorf("firstKid of %v shall be 222", n)
				}
			}
		case 7:
			if n.parent.Value.(v).x != 221 {
				t.Errorf("parent of %v shall be 221", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		}
		return false
	}, nil)

	if got != want {
		t.Errorf("added %v nodes got %v", want, got)
	}

	got = 0
	want = 3
	WalkThrough(stnew.Root, func(n *Node) bool {
		got++

		p, _ := n.Prev()
		if n.parent != p {
			t.Errorf("parent not retrieved %v", n)
		}

		switch n.Value.(v).x {
		case 1001:
			if stnew.Root != n {
				t.Errorf("root %v shall be %v", stnew.Root, n)
			}
			if n.parent != nil {
				t.Errorf("root has no parent %v", n)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 1002 {
					t.Errorf("firstKid of %v shall be 1002", n)
				}
			}
		case 1002:
			if n.parent.Value.(v).x != 1001 {
				t.Errorf("parent of %v shall be 1001", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 1003 {
					t.Errorf("firstKid of %v shall be 1003", n)
				}
			}
		case 1003:
			if n.parent.Value.(v).x != 1002 {
				t.Errorf("parent of %v shall be 1002", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		}
		return false
	}, nil)

	if got != want {
		t.Errorf("new tree has %v nodes got %v", want, got)
	}
}

func TestRemove(t *testing.T) {
	var st *Tree
	st = New()
	st.Add(v{x: 1003}, st.Add(v{x: 1002}, st.Add(v{x: 1001}, nil)))

	var tree *Tree
	tree = New()

	v1 := v{x: 1}
	v2 := v{x: 2}
	v3 := v{x: 3}

	v4 := v{x: 211}
	v5 := v{x: 212}

	v6 := v{x: 221}
	v7 := v{x: 222}

	n1 := tree.Add(v1, nil)
	n2 := tree.Add(v2, n1)
	tree.Add(v3, n2)

	n4 := tree.Add(v4, n1)
	tree.Add(v5, n4)

	tree.AddNode(st.Root, n1)

	n6 := tree.Add(v6, n1)
	tree.Add(v7, n6)

	c := tree.Remove(st.Root) //for testing Remove()

	got := 0
	want := 7

	WalkThrough(tree.Root, func(n *Node) bool {
		got++

		p, _ := n.Prev()
		if n.parent != p {
			t.Errorf("parent not retrieved %v", n)
		}

		switch n.Value.(v).x {
		case 1:
			if tree.Root != n {
				t.Errorf("root %v shall be %v", tree.Root, n)
			}
			if n.parent != nil {
				t.Errorf("root has no parent %v", n)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 2 {
					t.Errorf("firstKid of %v shall be 2", n)
				}
			}
		case 2:
			if n.parent.Value.(v).x != 1 {
				t.Errorf("parent of %v shall be 1", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling == nil {
				t.Errorf("%v shall have 1 nextSibling", n)
			} else {
				if n.nextSibling.Value.(v).x != 211 {
					t.Errorf("nextSibling of %v shall be 211", n)
				}
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 3 {
					t.Errorf("firstKid of %v shall be 3", n)
				}
			}
		case 3:
			if n.parent.Value.(v).x != 2 {
				t.Errorf("parent of %v shall be 2", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		case 4:
			if n.parent.Value.(v).x != 1 {
				t.Errorf("parent of %v shall be 1", n.parent.Value.(v).x)
			}
			if n.prevSibling == nil {
				t.Errorf("%v shall have 1 prevSibling", n)
			} else {
				if n.prevSibling.Value.(v).x != 2 {
					t.Errorf("prevSibling of %v shall be 2", n)
				}
			}
			if n.nextSibling == nil {
				t.Errorf("%v shall have 1 nextSibling", n)
			} else {
				if n.nextSibling.Value.(v).x != 221 {
					t.Errorf("nextSibling of %v shall be 221", n)
				}
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 212 {
					t.Errorf("firstKid of %v shall be 212", n)
				}
			}
		case 5:
			if n.parent.Value.(v).x != 211 {
				t.Errorf("parent of %v shall be 211", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		case 6:
			if n.parent.Value.(v).x != 1 {
				t.Errorf("parent of %v shall be 1", n.parent.Value.(v).x)
			}
			if n.prevSibling == nil {
				t.Errorf("%v shall have 1 prevSibling", n)
			} else {
				if n.prevSibling.Value.(v).x != 211 {
					t.Errorf("prevSibling of %v shall be 211", n)
				}
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid == nil {
				t.Errorf("%v shall have a firstKid", n)
			} else {
				if n.firstKid.Value.(v).x != 222 {
					t.Errorf("firstKid of %v shall be 222", n)
				}
			}
		case 7:
			if n.parent.Value.(v).x != 221 {
				t.Errorf("parent of %v shall be 221", n.parent.Value.(v).x)
			}
			if n.prevSibling != nil {
				t.Errorf("%v shall have no prevSibling", n)
			}
			if n.nextSibling != nil {
				t.Errorf("%v shall have no nextSibling", n)
			}
			if n.firstKid != nil {
				t.Errorf("%v shall have no firstKid", n)
			}
		}
		return false
	}, nil)

	if got != want {
		t.Errorf("added %v nodes got %v", want, got)
	}

	got = c
	want = 3
	if got != want {
		t.Errorf("new tree has %v nodes got %v", want, got)
	}
}

func TestPrevNext(t *testing.T) {
	var tree *Tree
	tree = New()

	v1 := v{x: 1}
	v2 := v{x: 2}
	v3 := v{x: 3}

	v4 := v{x: 211}
	v5 := v{x: 212}

	v6 := v{x: 221}
	v7 := v{x: 222}

	n1 := tree.Add(v1, nil)
	n2 := tree.Add(v2, n1)
	tree.Add(v3, n2)

	n4 := tree.Add(v4, n1)
	tree.Add(v5, n4)

	n6 := tree.Add(v6, n1)
	tree.Add(v7, n6)

	got := 0
	want := 3
	var tail *Node

	for n := tree.Root; n != nil; n, _ = n.Next() {
		tail = n
		got++
	}

	if got != want {
		t.Errorf("added %v nodes got %v", want, got)
	}

	got = 0
	for n := tail; n != nil; n, _ = n.Prev() {
		got++
	}

	if got != want {
		t.Errorf("added %v nodes got %v", want, got)
	}
}
