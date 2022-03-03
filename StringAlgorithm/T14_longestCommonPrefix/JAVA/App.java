public class App {
    public static void main(String[] args) {
        String[] case1={"flower","flow","flight"};
        System.out.println("case1:"+longestCommonPrefix(case1));
        String[] case2={"dog","racecar","car"};
        System.out.println("case2:"+longestCommonPrefix(case2));
    }

    public static String longestCommonPrefix(String[] strs) {
        if (strs==null||strs.length == 0) {
            return "";
        }
        // 最先开始假设第一个字符串是最长公共前缀
        String prefix = strs[0];
        int strsLen = strs.length;
        for (int i = 1; i < strsLen; i++) {
            prefix = longestCommonPrefix(prefix, strs[i]);
            if (prefix.length()==0) {
                break;
            }
        }
        return prefix;

    }

    //返回两个字符串的最长公共前缀
    public static String longestCommonPrefix(String s1, String s2) {
        int sMinLength = Math.min(s1.length(), s2.length());
        int endIndex = 0;
        while (endIndex < sMinLength && s1.charAt(endIndex)==s2.charAt(endIndex)) {
            endIndex++;
        }
        return s1.substring(0,endIndex);
    }
}
