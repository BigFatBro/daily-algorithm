from typing import List
class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        # 单调栈 + 哈希表,用空间换取时间，nums2中，从后向前遍历，计算出每一个值x，第一个比x大的值y,存储到哈希表中
        h_table = {}
        # 栈
        stack = []
        for value in reversed(nums2):
            # 将比当前value小的元素出栈
            while stack and stack[-1]<=value:
                stack.pop()
            h_table[value] = stack[-1] if stack else -1
            # 将当前元素入栈
            stack.append(value)
        # 查询哈希表获取结果
        return [h_table[value] for value in nums1]