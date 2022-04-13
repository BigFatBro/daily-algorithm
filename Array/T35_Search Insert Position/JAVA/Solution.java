public class Solution {
    public int searchInsert(int[] nums, int target) {
        int numsLen = nums.length;
        for (int i = 0; i < numsLen; i++) {
            if (nums[i]>=target) {
                return i;
            }
        }
        return numsLen;
    }
}
