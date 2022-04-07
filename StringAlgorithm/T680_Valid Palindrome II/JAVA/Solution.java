public class Test{
    public static void main(String[] args) {
        Solution solution = new Solution();
        String case1 = "abca";
        System.out.println("case1:"+solution.validPalindrome(case1));
    }
}

public class Solution {
    public boolean validPalindrome(String s) {
        int sLen = s.length();
        int i=0,j=sLen-1;
        while(i<j){
            if (s.charAt(i)==s.charAt(j)) {
                i++;
                j--;
            }else{
                return validPalindrome(s, i+1, j) || validPalindrome(s, i, j-1);
            }
        }
        return true;

    }

    public boolean validPalindrome(String s, int low, int high){
        for (int i = low,j=high; i < j; i++,j--) {
            if (s.charAt(i)!=s.charAt(j)) {
                return false;
            }
        }
        return true;
    }
}
