class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        num1_len = len(num1)
        num2_len = len(num2)
        jinwei = 0
        ans = ""
        for i in range(max(num1_len, num2_len)):
            if i < num1_len:
                num1_int = int(num1[num1_len - i - 1])
            else:
                num1_int = 0
            if i < num2_len:
                num2_int = int(num2[num2_len - i - 1])
            else:
                num2_int = 0
            sum = num1_int + num2_int + jinwei
            jinwei = sum // 10
            ans = str(sum % 10) + ans
        if jinwei == 1:
            ans = "1" + ans
        return ans