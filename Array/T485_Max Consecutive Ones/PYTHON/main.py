from typing import List


class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        maxCnt = cnt = 0
        for i, num in enumerate(nums):
            if num == 1:
                cnt += 1
            else:
                maxCnt = max(maxCnt, cnt)
                cnt = 0
        maxCnt = max(cnt, maxCnt)
        return maxCnt
