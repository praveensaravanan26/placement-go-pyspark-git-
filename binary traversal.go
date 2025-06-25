type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}
func inorder(root *TreeNode) {
 if root == nil {
 return
 }
 inorder(root.Left)
 fmt.Println(root.Val)
 inorder(root.Right)
}