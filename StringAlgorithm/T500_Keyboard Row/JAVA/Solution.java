import java.util.ArrayList;

public class Solution {
    public String[] findWords(String[] words) {
        rrayList<String> ans = new ArrayList<>();
        //能用数组就不要用map
        String rowIdx = "12210111011122000010020202";
        for (String word : words) {
            boolean isValid = true;
            //第一个字母的行号
            char idx = rowIdx.charAt(Character.toLowerCase(word.charAt(0))-'a');
            for (int i = 1; i < word.length(); i++) {
                if (rowIdx.charAt(Character.toLowerCase(word.charAt(i))-'a')!=idx) {
                    isValid = false;
                    break;
                }
            }
            if (isValid) {
                ans.add(word);
            }
        }

        String[] ansArray = new String[ans.size()];
        for (int i = 0; i < ans.size(); i++) {
            ansArray[i] = ans.get(i);
        } 
        return ansArray;
    }
}
