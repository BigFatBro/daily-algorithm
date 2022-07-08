from typing import List
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        # 双指针法
        slow = 0
        for fast in range(0,len(nums)):
            if nums[fast] != 0:
              nums[slow], nums[fast] = nums[fast], nums[slow]
              slow += 1