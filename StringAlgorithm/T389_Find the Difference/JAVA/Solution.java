public class Solution {
    public char findTheDifference(String s, String t) {
        int sLen = s.length();
        int[] cnt = new int[26];
        for (int i = 0; i < sLen; i++) {
            cnt[s.charAt(i)-'a']++;
        }

        int i = 0;
        while (true) {
            cnt[t.charAt(i)-'a']--;
            if (cnt[t.charAt(i)-'a'] < 0) {
                return t.charAt(i);
            }
        }


    }
}
