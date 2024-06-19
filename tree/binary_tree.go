package tree

import (
	"strconv"
	"strings"
)

type BinaryTree struct {
	root *Node
}

type Node struct {
	data   int
	parent *Node
	left   *Node
	right  *Node
}

/* Preorder Traversal */
func (tree *BinaryTree) PreorderTraversal(root *Node, callback func(int)) {
	if root == nil {
		return
	}
	callback(root.data)
	tree.PreorderTraversal(root.left, callback)
	tree.PreorderTraversal(root.right, callback)
}

/* Inorder Traversal */
func (tree *BinaryTree) InorderTraversal(root *Node, callback func(int)) {
	if root == nil {
		return
	}
	tree.PreorderTraversal(root.left, callback)
	callback(root.data)
	tree.PreorderTraversal(root.right, callback)
}

/* Postorder Traversal */
func (tree *BinaryTree) PostorderTraversal(root *Node, callback func(int)) {
	if root == nil {
		return
	}
	tree.PreorderTraversal(root.left, callback)
	tree.PreorderTraversal(root.right, callback)
	callback(root.data)
}

/* Levelorder Traversal */
func (tree *BinaryTree) LevelorderTraversal(root *Node) [][]int {
	var res [][]int
	arr := []*Node{root}
	if root == nil {
		return res
	}

	// Empty arr indicates that there is no node to travel through.
	for len(arr) > 0 {
		// Resetting the length of the arr
		size := len(arr)
		// Storing the nodes that would be popped off from the arr
		curRes := []int{}

		for i := 0; i < size; i++ {
			node := arr[i]
			curRes = append(curRes, node.data)
			// Finding the left and right children of the nodes in the arr, and push them into the arr if they has children nodes.
			if node.left != nil {
				arr = append(arr, node.left)
			}
			if node.right != nil {
				arr = append(arr, node.right)
			}
		}
		// Removing the nodes of the previous layer
		arr = arr[size:]
		res = append(res, curRes)
	}
	return res
}

/* Maximum Depth */
func (tree *BinaryTree) MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return max(tree.MaxDepth(root.left), tree.MaxDepth(root.right)) + 1 // 树的深度是root的高度，而root的高度是 左右孩子中较大者+1
}

/* The closest common ancestor node */
func (tree *BinaryTree) lowestCommonAncestor(root, p, q *Node) *Node {
	if root == nil {
		return nil
	}
	if root.data == p.data || root.data == q.data {
		return root
	}

	var findLeft = tree.lowestCommonAncestor(root.left, p, q)
	var findRight = tree.lowestCommonAncestor(root.right, p, q)

	if findLeft != nil && findRight != nil {
		return root
	} else if findLeft != nil {
		return findLeft
	} else {
		return findRight
	}
}

/* Constructing a Binary Tree ( From Preorder and Inorder Traversal ) */
func (tree *BinaryTree) buildTree1(preorder []int, inorder []int) *Node {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	left := tree.findRootIndex1(preorder[0], inorder)
	root := &Node{
		data:  preorder[0],
		left:  tree.buildTree1(preorder[1:left+1], inorder[:left]),
		right: tree.buildTree1(preorder[left+1:], inorder[left+1:])}
	return root
}
func (tree *BinaryTree) findRootIndex1(target int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

/* Constructing a Binary Tree ( From Postorder and Inorder Traversal ) */
func (tree *BinaryTree) buildTree2(inorder []int, postorder []int) *Node {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	//先找到根节点（后续遍历的最后一个就是根节点）
	nodeValue := postorder[len(postorder)-1]
	//从中序遍历中找到一分为二的点，左边为左子树，右边为右子树
	left := tree.findRootIndex2(inorder, nodeValue)
	//构造root
	root := &Node{
		data:  nodeValue,
		left:  tree.buildTree2(inorder[:left], postorder[:left]), //将后续遍历一分为二，左边为左子树，右边为右子树
		right: tree.buildTree2(inorder[left+1:], postorder[left:len(postorder)-1])}
	return root
}
func (tree *BinaryTree) findRootIndex2(inorder []int, target int) (index int) {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

// Serializes a tree to a single string.
func (tree *BinaryTree) serialize(root *Node) string {
	sb := &strings.Builder{}
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.data))
		sb.WriteByte(',')
		dfs(node.left)
		dfs(node.right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (tree *BinaryTree) deserialize(data string) *Node {
	sp := strings.Split(data, ",")
	var build func() *Node
	build = func() *Node {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &Node{
			data:  val,
			left:  build(),
			right: build(),
		}
	}
	return build()
}

/**
 *  树的子结构
 */
//用于递归遍历 A 中的所有节点，并判断当前节点 A 是否与 B 的根节点相同，相同则调用 helper( ) 进一步校验
func (tree *BinaryTree) isSubStructure(A *Node, B *Node) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	var ret bool

	//当在 A 中找到 B 的根节点时，进入helper递归校验
	if A.data == B.data {
		ret = tree.helper(A, B)
	}

	//ret == false，说明 B 的根节点不在当前 A 树顶中，进入 A 的左子树进行递归查找
	if !ret {
		ret = tree.isSubStructure(A.left, B)
	}

	//当 B 的根节点不在当前 A 树顶和左子树中，进入 A 的右子树进行递归查找
	if !ret {
		ret = tree.isSubStructure(A.right, B)
	}
	return ret

	//利用 || 的短路特性可写成
	//return helper(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
}

// helper 校验 B 是否与 A 的一个子树拥有相同的结构和节点值
func (tree *BinaryTree) helper(a, b *Node) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.data != b.data {
		return false
	}
	//a.data == b.data 递归校验 A B 左子树和右子树的结构和节点是否相同
	return tree.helper(a.left, b.left) && tree.helper(a.right, b.right)
}

/* Symmetric Binary Tree */
func isSymmetric(root *Node) bool {

	var defs func(left *Node, right *Node) bool
	defs = func(left *Node, right *Node) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.data != right.data {
			return false
		}
		return defs(left.left, right.right) && defs(left.right, right.left) // 如果左右都对称就返回true
	}

	return defs(root.left, root.right)
}

/* InvertTree func is for inverting the binary tree */
func InvertTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Recursively invert the left and right subtrees
	left := InvertTree(root.left)
	right := InvertTree(root.right)

	// Swap the left and right subtrees
	root.left = right
	root.right = left

	return root
}
