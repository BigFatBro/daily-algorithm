class Solution {
    public boolean checkRecord(String s) {
        int countA = 0, countL = 0 ;
        int sLen = s.length();
        for (int i = 0; i < sLen; i++) {
            if (s.charAt(i)=='A') {
                countA++;
                countL = 0;
                if (countA>=2) {
                    return false;
                }
            }else if (s.charAt(i)=='L') {
                countL++;
                if (countL>=3) {
                    return false;
                }
            }else{
                countL = 0;
            }
        }
        return true;
    }
}