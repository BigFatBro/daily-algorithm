import java.util.HashMap;

public class App {
    public static void main(String[] args) {
        String case1 = "abcabcbb";
        System.out.println(lengthOfLongestSubstring(case1));
        String case2 = "bbbbb";
        System.out.println(lengthOfLongestSubstring(case2));
        String case3 = "pwwkew";
        System.out.println(lengthOfLongestSubstring(case3));
    }

    public static int lengthOfLongestSubstring(String s) {
        if(s.length()==0) return 0;
        HashMap<Character,Integer> lookup = new HashMap<Character, Integer>();
        int n =s.length();
        int left = 0;
        int maxLen = 0;
        for (int i = 0; i < n; i++) {
            if(lookup.containsKey(s.charAt(i))){
                // 从lookup.get(s.charAt(i)) + 1 处开始为最长不重复子串的第一个字符的index
                left = Math.max(left,lookup.get(s.charAt(i)) + 1);
            }
            lookup.put(s.charAt(i), i);
            maxLen = Math.max(maxLen, i-left+1);
        }
        return maxLen;
    }
}