package array

type BinaryTree struct {
	root *Node
}

type Node struct {
	data int
	parent *Node
	left *Node
	right *Node
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
func (tree *BinaryTree) InorderTraversal(root *Node, callback func(int)) {
	if root == nil {
		return
	}
	tree.PreorderTraversal(root.left, callback)
	tree.PreorderTraversal(root.right, callback)
	callback(root.data)
}

