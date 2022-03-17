public class Test{
    public static void main(String[] args) {
        Solution solution = new Solution();
        char[] s ={'h','e','l','l','o'};
        solution.reverseString(s);
        for (char c : s) {
            System.out.printf("%c ", c);;
        }

    }
}
public class Solution {
    public void reverseString(char[] s) {
        //双指针法
        int sLen = s.length;
        int stopCondition = s.length/2;
        for (int i = 0; i < stopCondition; i++) {
            char temp = s[i];
            s[i] = s[sLen-i-1];
            s[sLen-i-1] = temp;
        }
    }
}
