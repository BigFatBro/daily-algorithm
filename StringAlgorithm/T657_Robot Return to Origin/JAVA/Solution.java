class Solution {
    public boolean judgeCircle(String moves) {
        int x=0,y=0, sLen=moves.length();
        for (int i = 0; i < sLen; i++) {
            if (moves.charAt(i)=='U') {
                y++;
            }else if (moves.charAt(i)=='D') {
                y--;
            }else if (moves.charAt(i)=='L') {
                x--;
            }else{
                x++;
            }
        }

        if (x==0&&y==0) {
            return true;
        }else{
            return false;
        }

    }
}