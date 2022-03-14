import java.util.Map;
import java.util.HashMap;

class Test{
    public static void main(String[] args) {
        Solution solution = new Solution();
        String case1_s = "egg", case1_t = "add";
        String case2_s = "foo", case2_t = "bar";
        String case3_s = "paper", case3_t = "title";
        System.out.println("case1:"+solution.isIsomorphic(case1_s, case1_t));
        System.out.println("case2:"+solution.isIsomorphic(case2_s, case2_t));
        System.out.println("case3:"+solution.isIsomorphic(case3_s, case3_t));

        System.out.println("case1:"+solution.isIsomorphicV2(case1_s, case1_t));
        System.out.println("case2:"+solution.isIsomorphicV2(case2_s, case2_t));
        System.out.println("case3:"+solution.isIsomorphicV2(case3_s, case3_t));
    }
}

class Solution {
    public boolean isIsomorphic(String s, String t) {
        if (s.length()!=t.length()) {
            return false;
        }
        
        Map<Character,Character> sMap = new HashMap<>();
        Map<Character,Character> tMap = new HashMap<>();
        int sLen = s.length();
        for (int i = 0; i < sLen; i++) {
            char sChar = s.charAt(i), tChar = t.charAt(i);
            if ((sMap.containsKey(sChar) && sMap.get(sChar) != tChar ) || (tMap.containsKey(tChar) && tMap.get(tChar) != sChar)) {
                return false;
            }else{
                sMap.put(sChar, tChar);
                tMap.put(tChar, sChar);
            }
        }
       
        return true;
    }

    public boolean isIsomorphicV2(String s, String t) {
        if (s.length()!=t.length()) {
            return false;
        }
        int[] sMap = new int[128];
        int[] tMap = new int[128];
        StringBuilder sPattern = new StringBuilder();
        StringBuilder tPattern = new StringBuilder();

        int sLen = s.length();

        for (int i = 0; i < sLen; i++) {
            char sChar = s.charAt(i),tChar = t.charAt(i);
            if (sMap[sChar]==0) {
                sMap[sChar] = i+1;
            }
            if (tMap[tChar]==0) {
                tMap[tChar] = i+1;
            }
            sPattern.append(sMap[sChar]);
            tPattern.append(tMap[tChar]);
        }

        if (sPattern.toString().equals(tPattern.toString())) {
            return true;
        }else{
            return false;
        }
    }

}