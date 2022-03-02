public class App {
    public static void main(String[] args) {
        String case1 = "babad";
        System.out.println("case1:"+ longestPalindrome(case1));
        String case2 = "cbbd";
        System.out.println("case2:"+ longestPalindrome(case2));
    }

    public static boolean isPalindromic(String s) {
        int sLen = s.length();
        for (int i = 0; i < sLen; i++) {
            if (s.charAt(i) != s.charAt(sLen - i - 1)) {
                return false;
            }
        }
        return true;
    }

    public static String longestPalindrome(String s) {
        int sLen = s.length();
        String ans = "";
        int max = 0;
        for (int i = 0; i < sLen; i++) {
            for (int j = i + 1; j <= sLen; j++) {
                String subStr = s.substring(i, j);
                if (isPalindromic(subStr) && subStr.length() > max) {
                    ans = subStr;
                    max = subStr.length();
                }

            }
        }

        return ans;
    }
}