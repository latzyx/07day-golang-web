package gee

import (
	"strings"
)

type node struct {
	Pattern  string  // 等待匹配的路由
	part     string  // 路由的一部分
	childern []*node // 子节点
	isWild   bool    // 是否成功匹配
}

// 第一个匹配成功的节点,用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.childern {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点,用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range nodes {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 匹配成功进行插入
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			Pattern: part,
			isWild:  part[0] == ':' || part[0] == '*',
		}
		n.childern = append(n.childern, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
