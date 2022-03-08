import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.HashMap;
import java.lang.StringBuffer;


class Test{
    public static void main(String[] args) {
        Solution solution = new Solution();
        String case1 = "23";
        String case2 = "";
        String case3 = "2";
        System.out.println(solution.letterCombinations(case1));
        System.out.println(solution.letterCombinations(case2));
        System.out.println(solution.letterCombinations(case3));
    }
}


class Solution {
    public List<String> letterCombinations(String digits) {
        List<String> combinations = new ArrayList<>();
        if (digits.length()==0 || digits == null) {
            return combinations;
        }

        Map<Character, String> phoneMap = new HashMap<Character, String>() {{
            put('2', "abc");
            put('3', "def");
            put('4', "ghi");
            put('5', "jkl");
            put('6', "mno");
            put('7', "pqrs");
            put('8', "tuv");
            put('9', "wxyz");
        }};
        backtrack(combinations, phoneMap, digits, 0, new StringBuffer());
        return combinations;

        
    }

    public void backtrack(List<String> combinations, Map<Character, String> phoneMap, String digits, int index, StringBuffer combination) {
        if (index == digits.length()) {
            combinations.add(combination.toString());
        }else{
            char t = digits.charAt(index);
            String letters = phoneMap.get(t);
            for (int i = 0; i < letters.length(); i++) {

                combination.append(letters.charAt(i));
                backtrack(combinations, phoneMap, digits, index+1, combination);
                combination.deleteCharAt(index);
                
            }
            
        }
        
    }
}