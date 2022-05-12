from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        inf = int(1e9)
        minPrice = inf
        maxEarn = 0
        for price in prices:
            minPrice = min(minPrice, price)
            maxEarn = max(maxEarn, price - minPrice)
        return maxEarn
    
    
    
if __name__ == "__main__":
    prices = [7,1,5,3,6,4]
    print(Solution().maxProfit(prices))
    
# notes: 经测试，在数据少的时候，min()和max()函数的性能不如if语句的性能