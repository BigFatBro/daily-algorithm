public class Solution {
    public static void main(String[] args) {
        int case1 = 3;
        int case2 = 4;
        int case3 = 9;
        int case4 = 58;
        int case5 = 1994;
        System.out.println("case1:"+intToRoman(case1));
        System.out.println("case2:"+intToRoman(case2));
        System.out.println("case3:"+intToRoman(case3));
        System.out.println("case4:"+intToRoman(case4));
        System.out.println("case5:"+intToRoman(case5));
    }

    public static String intToRoman(int num) {
        int[] values = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        String[] symbols = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
        StringBuffer ansRoman = new StringBuffer();
        for (int i = 0; i < values.length; i++) {
            while (num>=values[i]) {
                num = num - values[i];
                ansRoman.append(symbols[i]);
            }
            if (num == 0) {
                break;
            }
        }
        return ansRoman.toString();
    }

}
