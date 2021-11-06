package linkedlistcycle_test

import linkedlistcycle "github.com/ju-popov/structy.net/076-linked-list-cycle"

type testCase struct {
	name     string
	head     *linkedlistcycle.Node
	expected bool
}

func newLinkedList(values ...string) *linkedlistcycle.Node {
	dummyHead := linkedlistcycle.NewNode("")
	current := dummyHead

	for _, value := range values {
		current.Next = linkedlistcycle.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

func test00Head() *linkedlistcycle.Node {
	//nolint:varnamelen
	a := linkedlistcycle.NewNode("a")
	//nolint:varnamelen
	b := linkedlistcycle.NewNode("b")
	//nolint:varnamelen
	c := linkedlistcycle.NewNode("c")

	d := linkedlistcycle.NewNode("d")

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b // cycle

	return a
}

func test01Head() *linkedlistcycle.Node {
	//nolint:varnamelen
	q := linkedlistcycle.NewNode("q")
	//nolint:varnamelen
	r := linkedlistcycle.NewNode("r")
	//nolint:varnamelen
	s := linkedlistcycle.NewNode("s")
	//nolint:varnamelen
	t := linkedlistcycle.NewNode("t")
	//nolint:varnamelen
	u := linkedlistcycle.NewNode("u")

	q.Next = r
	r.Next = s
	s.Next = t
	t.Next = u
	u.Next = q // cycle

	return q
}

func test03Head() *linkedlistcycle.Node {
	//nolint:varnamelen
	q := linkedlistcycle.NewNode("q")
	//nolint:varnamelen
	r := linkedlistcycle.NewNode("r")
	//nolint:varnamelen
	s := linkedlistcycle.NewNode("s")
	//nolint:varnamelen
	t := linkedlistcycle.NewNode("t")
	//nolint:varnamelen
	u := linkedlistcycle.NewNode("u")

	q.Next = r
	r.Next = s
	s.Next = t
	t.Next = u
	u.Next = t // cycle

	return q
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     test00Head(),
		expected: true,
	},
	{
		name:     "test_01",
		head:     test01Head(),
		expected: true,
	},
	{
		name:     "test_02",
		head:     newLinkedList("a", "b", "c", "d"),
		expected: false,
	},
	{
		name:     "test_03",
		head:     test03Head(),
		expected: true,
	},
	{
		name:     "test_04",
		head:     newLinkedList("p"),
		expected: false,
	},
	{
		name:     "test_05",
		head:     newLinkedList(),
		expected: false,
	},
}
