class Test{
    public static void main(String[] args) {
        String num1 = "11", num2 = "123";
        Solution solution = new Solution();
        System.out.println("case1:"+ solution.addStrings(num1, num2));
    }
}


public class Solution {
    public String addStrings(String num1, String num2) {
        int i=num1.length()-1,j=num2.length()-1, jinwei=0;
        StringBuffer ans = new StringBuffer();
        while (i>=0||j>=0||jinwei!=0) {
            int x,y;
            if (i>=0) {
                x = num1.charAt(i)-'0';
            }else{
                x = 0;
            }
            if (j>=0) {
                y = num2.charAt(j)-'0';
            } else {
                y=0;
            }

            int res=x+y+jinwei;
            ans.append(res%10);
            jinwei = res/10;
            i--;
            j--;
        }
        ans.reverse();
        return ans.toString();
    }
}
