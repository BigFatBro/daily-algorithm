import java.util.Map;
import java.util.HashMap;
class Test{
    public static void main(String[] args) {
        Solution solution =new Solution();
        String case1Pattern = "abba", case1S = "dog cat cat dog";
        String case2Pattern = "abba", case2S = "dog cat cat fish";
        String case3Pattern = "aaaa", case3S = "dog cat cat dog";

        System.out.println("case1:"+ solution.wordPattern(case1Pattern, case1S));
        System.out.println("case2:"+ solution.wordPattern(case2Pattern, case2S));
        System.out.println("case3:"+ solution.wordPattern(case3Pattern, case3S));
    }
}

class Solution {
    public boolean wordPattern(String pattern, String s) {
        // 切分s
        String[] sArrry = s.split(" ");
        // System.out.println(sArrry.length);
        // System.out.println(sArrry[0]);
        // System.out.println(sArrry[1]);
        // System.out.println(sArrry[2]);
        // System.out.println(sArrry[3]);
        if (pattern.length() != sArrry.length) {
            return false;
        }

        Map<Character,String> patternMap = new HashMap<>();
        Map<String,Character> sMap = new HashMap<>();

        int patternLen = pattern.length();
        for (int i = 0; i < patternLen; i++) {
            char patternChar = pattern.charAt(i);
            String sString = sArrry[i];
            if ((patternMap.containsKey(patternChar) && !patternMap.get(patternChar).equals(sString)) || (sMap.containsKey(sString) && sMap.get(sString) != patternChar)) {
                // System.out.println(patternMap.containsKey(patternChar));
                // System.out.println(patternMap.get(patternChar).equals(sString));
                //System.out.println("FALSE:"+ " " + i + " " + patternChar + " "+sString);
                return false;
            }else{
                patternMap.put(patternChar, sString);
                sMap.put(sString, patternChar);
            }

        }

        return true;

    }
}