class Test {
    public static void main(String[] args) {

        Solution solution = new Solution();
        String case1 = "Hello World";
        String case2 = "   fly me   to   the moon  ";
        String case3 = "luffy is still joyboy";
        String case4 = "a";

        System.out.println("case1:"+ solution.lengthOfLastWord(case1));
        System.out.println("case2:"+ solution.lengthOfLastWord(case2));
        System.out.println("case3:"+ solution.lengthOfLastWord(case3));
        System.out.println("case4:"+ solution.lengthOfLastWord(case4));
  
    }
}

class Solution {
    public int lengthOfLastWord(String s) {
        int n = s.length();
       
        int lastWordEnd = n-1;
        //寻找第一个不是空格的位置
        // s中必定存在字母
        while((s.charAt(lastWordEnd) == ' ')){
            lastWordEnd--;
        }

        int lastWordBegin =lastWordEnd;
        
        while(lastWordBegin >=0 && s.charAt(lastWordBegin) != ' '){
            lastWordBegin = lastWordBegin -1;
        }

       
        return lastWordEnd-lastWordBegin;

    }
}
