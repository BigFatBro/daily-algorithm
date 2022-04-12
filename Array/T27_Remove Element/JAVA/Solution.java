public class Solution {
    public int removeElement(int[] nums, int val) {
        int numsLen = nums.length;
        int slow = 0;
        for (int fast = 0; fast < numsLen; fast++) {
            if (nums[fast] != val) {
                nums[slow] = nums[fast];
			    slow++;
            }
        }
        return slow;

    }
}
