import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class Test{
    public static void main(String[] args) {
        int[] nums = {-10,-3,0,5,9};
        Solution solution = new Solution();
        TreeNode root = solution.sortedArrayToBST(nums);
        List<Integer> list = solution.floorTraverse(root);
        System.out.println("case1: ");
        
        for (int i = 0; i < list.size(); i++) {
            System.out.println(list.get(i));
        }
    }
}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {}
    TreeNode(int val) {this.val = val;}
    TreeNode(int val, TreeNode left, TreeNode right){
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

class Solution {
    // 通过中序遍历序列，重建二叉搜索树（结果不唯一）
    public TreeNode sortedArrayToBST(int[] nums) {
        return helper(nums, 0, nums.length-1);
    }

    public TreeNode helper(int[] nums, int left, int right){
        // 递归结束
        if (left>right) {
            return null;
        }
        int mid = (left+right)/2;
        TreeNode root = new TreeNode(nums[mid]);
        root.left = helper(nums, left, mid - 1);
        root.right = helper(nums, mid + 1, right);
        return root;
    }

    // 层序遍历
    public List<Integer> floorTraverse(TreeNode root) {
        List<Integer> ans = new ArrayList<>();
        Queue<TreeNode> temp = new LinkedList<>();
        // 根节点入队
        temp.add(root);

        while (!temp.isEmpty()) {
            //队首元素出队
            TreeNode tempNode= temp.remove();
            ans.add(Integer.valueOf(tempNode.val));
            if (tempNode.left != null) {
                temp.add(tempNode.left);
            }
            if (tempNode.right != null) {
                temp.add(tempNode.right);
            }
        }

        return ans;
    }
}
