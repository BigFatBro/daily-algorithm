import java.io.ObjectOutputStream.PutField;
import java.util.Map;
import java.util.HashMap;

public class Solution {
    public static void main(String[] args) {
        String case1 = "III";
        String case2 = "IV";
        String case3 = "IX";
        String case4 = "LVIII";
        String case5 = "MCMXCIV";
        System.out.println("case1:"+romanToInt(case1));
        System.out.println("case2:"+romanToInt(case2));
        System.out.println("case3:"+romanToInt(case3));
        System.out.println("case4:"+romanToInt(case4));
        System.out.println("case5:"+romanToInt(case5));
    }

    public static int romanToInt(String s) {
        if (s.length() == 0 || s == null) {
            return 0;
        }
        Map<String, Integer> map = new HashMap<String, Integer>();
        map.put("CD",400);
        map.put("CM",900);
        map.put("XL",40);
        map.put("XC",90);
        map.put("IV",4);
        map.put("IX",9);
        map.put("M",1000);
        map.put("D",500);
        map.put("C",100);
        map.put("L",50);
        map.put("X",10);
        map.put("V",5);
        map.put("I",1);
        
        int ans = 0;
        int i = 0;
        //遍历到倒数第二个字母
        while (i<s.length()-1) {
            String sTwo = s.substring(i, i+2);
            if (map.containsKey(sTwo)) {
                ans = ans + map.get(sTwo);
                i = i+2;
            }else{
                String sOne = s.substring(i, i+1);
                ans = ans + map.get(sOne);
                i = i+1;
            }
            
            
        }
        //处理最后一个字母
        if (i==s.length()-1) {
            String sOne = s.substring(i, i+1);
            ans = ans + map.get(sOne);  
        }
        return ans;
        


    }
    
}
