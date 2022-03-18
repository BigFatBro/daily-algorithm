public class Solution {
    //双指针法
    public String reverseVowels(String s) {
        int n = s.length();
        if (n<=1) {
            return s;
        }
        char[] arr=s.toCharArray();
        int left=0, right = n-1;
        while (left<right) {
            if (isVowel(arr[left]) && isVowel(arr[right])) {
                char tmp = arr[left];
                arr[left] = arr[right];
                arr[right] = tmp;

                left++;
                right--;
            }

            if(!isVowel(arr[left])) left++;
            if (!isVowel(arr[right])) right--;
            
        }

        return new String(arr);

        
    }

    public boolean isVowel(char c) {
        return c == 'a' || c == 'A'
                || c == 'e' || c == 'E'
                || c == 'i' || c == 'I'
                || c == 'o' || c == 'O'
                || c == 'u' || c == 'U';
    }
}
