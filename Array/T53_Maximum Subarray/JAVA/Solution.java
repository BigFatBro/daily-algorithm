public class Test {
    public static void main(String[] args) {
        int[] nums = {-2,1,-3,4,-1,2,1,-5, 4};
        Solution solution = new Solution();
        System.out.println("case1:"+solution.maxSubArray(nums));
    }
    
}

class Solution {
    public int maxSubArray(int[] nums) {
        //贪心算法：遍历数组的每一个值，计算当前元素之前的子数列的和，若之前和小于0，则丢弃当前元素之前的数列
        //一个网友的评论：这道题目的思想是： 走完这一生 如果我和你在一起会变得更好，那我们就在一起，否则我就丢下你。 我回顾我最光辉的时刻就是和不同人在一起，变得更好的最长连续时刻
        int maxSum = Integer.MIN_VALUE;
        int curSum = 0;
        int preSum = 0;
        for (int i = 0; i < nums.length; i++) {
            if (preSum > 0) {
                curSum = preSum + nums[i];
            }else{
                curSum = nums[i];
            }
            if (maxSum < curSum) {
                maxSum = curSum;
            }
            preSum = curSum;
            
        }

        return maxSum;

    }
}
