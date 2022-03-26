class Solution {
    public int countSegments(String s) {
        int count=0;
        int sLen = s.length();
        for (int i = 0; i < sLen; i++) {
            if ((s.charAt(i)!=' ')&&(i==0||s.charAt(i-1)==' ')) {
                count++;
            }
        }
        return count;
    }
}