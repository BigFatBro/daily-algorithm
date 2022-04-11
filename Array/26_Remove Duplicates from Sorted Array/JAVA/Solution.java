public class Solution {
    public int removeDuplicates(int[] nums) {
        int numsLen = nums.length;
        if (numsLen==0) {
            return 0;
        }
        int slow=1;
        for (int fast = 1; fast < numsLen; fast++) {
            if (nums[fast]!=nums[fast-1]) {
                nums[slow]=nums[fast];
                slow++;
                
            }
            
        }
        return slow;
    }
    
}
