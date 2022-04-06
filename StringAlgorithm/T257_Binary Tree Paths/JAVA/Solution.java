import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;


class TreeNode{
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {};
    TreeNode(int val) {
        this.val = val;
    }
    TreeNode(int val, TreeNode left, TreeNode right){
        this.val = val;
        this.left = left;
        this.right = right;
    }
}
class Solution{
    public List<String> binaryTreePaths(TreeNode root) {
        //深度优先遍历
        List<String> paths = new ArrayList<>();
        consructPaths(root, "", paths);
        return paths;
    }

    public void consructPaths(TreeNode root, String path, List<String> paths){
        if (root!=null) {
            StringBuffer pathB = new StringBuffer(path);
            pathB.append(Integer.toString(root.val));
            if (root.left == null && root.right==null) {
                paths.add(pathB.toString());
                
            } else {
                pathB.append("->");
                consructPaths(root.left, pathB.toString(), paths);
                consructPaths(root.right, pathB.toString(), paths);
            }
        }
    }

    public List<String> binaryTreePathsMethod2(TreeNode root) {
        //广度优先搜索
        List<String> paths = new ArrayList<>();
        if (root == null) {
            return paths;
        }

        Queue<TreeNode> nodeQueue = new LinkedList<>();
        Queue<String> pathQueue = new LinkedList<>();

        nodeQueue.offer(root);
        pathQueue.offer(Integer.toString(root.val));

        while (!nodeQueue.isEmpty()) {
            TreeNode node = nodeQueue.poll();
            String path = pathQueue.poll();

            if(node.left == null && node.right == null){
                paths.add(path);
            }else{
                if(node.left != null){
                    nodeQueue.offer(node.left);
                    pathQueue.offer(new StringBuffer(path).append("->").append(node.left.val).toString());
                }
                if(node.right!=null){
                    nodeQueue.offer(node.right);
                    pathQueue.offer(new StringBuffer(path).append("->").append(node.right.val).toString());
                }
            }
        }

        return paths;
    }
}