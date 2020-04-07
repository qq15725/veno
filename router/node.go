package router

import "strings"

type Node struct {
	pattern  string
	part     string
	children []*Node
	isWild   bool
}

func (node *Node) matchChild(part string) *Node {
	for _, child := range node.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (node *Node) matchChildren(part string) []*Node {
	children := make([]*Node, 0)
	for _, child := range node.children {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}
	return children
}

func (node *Node) insertChild(pattern string, parts []string, depth int) {
	if len(parts) == depth {
		node.pattern = pattern
	} else {
		part := parts[depth]
		child := node.matchChild(part)
		if child == nil {
			child = &Node{part: part, isWild: part[0] == ':' || part[0] == '*'}
			node.children = append(node.children, child)
		}
		child.insertChild(pattern, parts, depth+1)
	}
}

func (node *Node) insert(pattern string) {
	node.insertChild(pattern, parsePattern(pattern), 0)
}

func (node *Node) searchChild(parts []string, depth int) *Node {
	if len(parts) == depth || strings.HasPrefix(node.part, "*") {
		if node.pattern == "" {
			return nil
		}
		return node
	}
	part := parts[depth]
	children := node.matchChildren(part)
	for _, child := range children {
		target := child.searchChild(parts, depth+1)
		if target != nil {
			return target
		}
	}
	return nil
}

func (node *Node) search(pattern string) *Node {
	return node.searchChild(parsePattern(pattern), 0)
}

func parsePattern(pattern string) []string {
	parts := make([]string, 0)
	for _, item := range strings.Split(pattern, "/") {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
