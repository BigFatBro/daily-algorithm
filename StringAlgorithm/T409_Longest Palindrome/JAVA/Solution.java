public class Solution {
    public int longestPalindrome(String s) {
        int sLen = s.length();
        int[] cnt = new int[128];
        for (int i = 0; i < sLen; i++) {
            cnt[s.charAt(i)]++;
        }

        int ans = 0;
        for (int i = 0; i < 128; i++) {
            ans += cnt[i]/2*2;
            if ((cnt[i]%2==1)&&(ans%2==0)) {
                ans++;
            }
        }
        return ans;
    }

}
