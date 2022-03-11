class Test{
    public static void main(String[] args) {
        Solution solution = new Solution();
        String a = "11", b = "1";
        System.out.println(solution.addBinary2(a, b));
    }
}

class Solution {
    //method 1
    public String addBinary(String a, String b) {
        return Integer.toBinaryString(Integer.parseInt(a,2) + Integer.parseInt(b, 2));
    }

    //method 2
    public String addBinary2(String a, String b) {
        StringBuilder ans = new StringBuilder();
        //ca表示进位
        int ca=0;

        for (int i = a.length()-1, j = b.length()-1; i>=0||j>=0; i--,j--) {
            int sum=ca;
            sum += (i>=0 ? a.charAt(i)-'0' : 0 );
            sum += (j>=0 ? b.charAt(j)-'0' : 0);
            ans.append(sum%2);
            ca = sum/2;
            
        }

        //验证最后是否有进位
        ans.append(ca == 1 ? ca : "");

        return ans.reverse().toString();
    }
}