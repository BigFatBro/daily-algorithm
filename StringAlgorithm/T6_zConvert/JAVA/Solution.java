class Solution {

    public static void main(String[] args) {
        System.out.println(convert("PAYPALISHIRING", 3));
        System.out.println(convert("PAYPALISHIRING", 4));
    }


    public static String convert(String s, int numRows) {
        int n = s.length(),r=numRows;
        if (r==1 || r>=n) {
            return s;
        }

        //一个周期需要的字符数
        int t = r+r-2;
        //一个周期需要1+r-2=r-1列
        //需要的周期数n/t,然后向上取整，最后一个周期视作完整周期
        //需要的矩阵列数
        int c=(n+t-1)/t*(r-1);
        //创建一个 r行 c 列的矩阵
        char[][] mat = new char[r][c];
        for (int i = 0,j=0,k=0; i < n; i++) {
            mat[j][k] = s.charAt(i);
            if (i % t < r-1) {
                //向下移动
                ++j;
            }else{
                //向右上移动
                --j;
                ++k;
            }
        }
        StringBuffer ans = new StringBuffer();
        for (char[] row : mat) {
            for (char ch : row) {
                if(ch!=0){
                    ans.append(ch);
                }
                
            }
        }
        return ans.toString();

    }
}