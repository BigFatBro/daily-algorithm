public class Solution {
    public int firstUniqChar(String s) {
        int sLen = s.length();
        int[] cnt = new int[26];

        for (int i = 0; i < sLen; i++) {
            cnt[s.charAt(i)-'a']++;
        }
        for (int i = 0; i < sLen; i++) {
            if (cnt[s.charAt(i)-'a']==1) {
                return i;
            }
        }

        return -1;

    }
}
