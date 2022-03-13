import javax.swing.plaf.basic.BasicInternalFrameTitlePane.TitlePaneLayout;

class Test {
    public static void main(String[] args) {
        // System.out.println((int)('B'-'A'));
        Solution solution = new Solution();
        String case1 = "A", case2 = "AB", case3 = "ZY";
        System.out.println("case1:" + solution.titleToNumber(case1));
        System.out.println("case2:" + solution.titleToNumber(case2));
        System.out.println("case3:" + solution.titleToNumber(case3));
    }
}

class Solution {
    public int titleToNumber(String columnTitle) {
        int ans = 0;
        int sLen = columnTitle.length();
        for (int i = 0; i < sLen; i++) {
            ans = ans * 26 + (int) (columnTitle.charAt(i) - 'A') + 1;
        }

        return ans;

    }
}