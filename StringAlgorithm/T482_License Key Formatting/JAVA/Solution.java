public class Solution {
    public String licenseKeyFormatting(String s, int k) {
        StringBuffer ans = new StringBuffer();
        int sLen = s.length();
        for (int i = sLen-1, cnt = 0 ; i >=0 ; i--) {
            if (s.charAt(i)!='-') {
                ans.append(Character.toUpperCase(s.charAt(i)));
                cnt++;
                if (cnt%k==0) {
                    ans.append("-");
                }
            }
        }

        if (ans.length() > 0 && ans.charAt(ans.length() - 1) == '-'){
            ans.deleteCharAt(ans.length() - 1);
        }
        return ans.reverse().toString();
    }
}
