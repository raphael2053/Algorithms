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

func LevelorderTraversal(root *Node) [][]int {
	//定义一个存储结果的二维切片
	var res [][]int
	//定义一个含有根节点的队列
	arr := []*Node{root}
	//根节点为空直接返回空
	if root == nil {
		return res
	}
	//当队列为空就返回结果，不为空就进入循环
	for len(arr) > 0 {
		//重置队列长度
		size := len(arr)
		//定义一个切片来存储出队列的值
		curRes := []int{}
		//遍历队列
		for i := 0; i < size; i++ {
			//临时变量存储队列中的元素
			node := arr[i]
			//将队列中元素的值加入到这个切片里。
			curRes = append(curRes, node.data)
			//寻找队列里元素的左右孩子，如果不为空就入队
			if node.left != nil {
				arr = append(arr, node.left)
			}
			if node.right != nil {
				arr = append(arr, node.right)
			}
		}
		//遍历队列结束后，把上一层对应的元素出队。
		arr = arr[size:]
		//把所有结果加入到最后的结果上
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

/**
* 最近公共祖先
 */
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
