class Test {
    public static void main(String[] args) {
        Solution solution = new Solution();
        int case1=1,case2=27,case3=701;
        System.out.println("case1:"+solution.convertToTitle(case1));
        System.out.println("case2:"+solution.convertToTitle(case2));
        System.out.println("case3:"+solution.convertToTitle(case3));
    }
}

public class Solution {
    public String convertToTitle(int columnNumber) {
        StringBuilder ans = new StringBuilder();
        while(columnNumber>0){
            columnNumber--;
            ans.append((char)('A'+columnNumber%26));
            columnNumber = columnNumber / 26;
        }
        return ans.reverse().toString();

    }
    
}
