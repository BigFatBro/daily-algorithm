public class Solution {
    public boolean repeatedSubstringPattern(String s) {
        int sLen = s.length();
        for (int i = 1; i*2 <= sLen; i++) {
            if (sLen % i == 0) {
                boolean match = true;
                for (int j = i; j < sLen; j++) {
                    if (s.charAt(j) != s.charAt(j-i)) {
                        match = false;
                        break;
                    }
                }
                if (match) {
                    return true;
                }
            }
            
        }
        return false;

    }
    
}
