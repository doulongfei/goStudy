public class TreeNode {
    int val = 0;
    TreeNode left = null;
    TreeNode right = null;

    public TreeNode(int val) {
        this.val = val;

    }

}
public class Solution {
    public boolean HasSubtree(TreeNode root1,TreeNode root2) {
        if(root2==null) return false;
        if(root1==null && root2!=null) return false;
        boolean flag = false;
        if(root1.val==root2.val){
            // 两个节点相同时，验证以两个节点开始的子树是否相同
            flag = isSubTree(root1,root2);
        }
        if(!flag){
            // 递归左子树进行判断
            flag = HasSubtree(root1.left, root2);
            if(!flag){
                // 递归右子树进行判断
                flag = HasSubtree(root1.right, root2);
            }
        }
        return flag;
    }
    private boolean isSubTree(TreeNode root1,TreeNode root2){
        if(root2==null) {
            return true;
        }
        if(root1==null&&root2!=null){
            return false;
        }
        if(root1.val==root2.val){
            return isSubTree(root1.left,root2.left)&&isSubTree(root1.right,root2.right);
        }else{
            return false;
        }

    }
}