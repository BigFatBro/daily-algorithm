import java.util.Map;
import java.util.Stack;
import java.util.Map;
import java.util.HashMap;


class Test{
    public static void main(String[] args) {
        String case1 = "()[]{}";
        String case2 = "(]";
        String case3 = "([)]";
        String case4 = "{[]}";

        Solution solution = new Solution();
        System.out.println("case1:" + solution.isValid(case1));
        System.out.println("case2:" + solution.isValid(case2));
        System.out.println("case3:" + solution.isValid(case3));
        System.out.println("case4:" + solution.isValid(case4));
    }
}

class Solution {
    public boolean isValid(String s) {
        
        Stack<Character> stack = new Stack<>();

        for (int i = 0; i < s.length(); i++) {
            char currentChar = s.charAt(i);
            Character ch = (Character) s.charAt(i);
            if (currentChar ==')' || currentChar == ']' || currentChar == '}') {
                //出栈
                if (stack.empty()) {
                    //栈为空
                    return false;
                    
                } else {
                    //栈非空
                    Character chTop = (Character) stack.pop();
                    char c = chTop.charValue();
                    if (c != '(' && c!='[' && c!='{') {
                        return false;
                    } else {
                        if(c == '(' && currentChar!=')'){
                            return false;
                        }
                        if(c == '[' && currentChar!=']'){
                            return false;
                        }
                        if(c == '{' && currentChar!='}'){
                            return false;
                        }
                    }
                    
                }
            } else {
                //入栈
                stack.push(ch);
            }
        }
        //最后栈非空，返回false
        if(!stack.empty()){
            return false;
        }

        return true;

    }
}