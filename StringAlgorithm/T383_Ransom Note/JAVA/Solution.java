class Solution {
    public boolean canConstruct(String ransomNote, String magazine) {
        if (magazine.length()<ransomNote.length()) {
            return false;
        }
        int[] cnt = new int[26];
        int rLen = ransomNote.length();
        int mLen = magazine.length();
        for (int i = 0; i < mLen; i++) {
            cnt[magazine.charAt(i)-'a']++;
        }
        for (int i = 0; i < rLen; i++) {
            cnt[ransomNote.charAt(i)-'a']--;
            if (cnt[ransomNote.charAt(i)-'a']<0) {
                return false;
            }
        }

        return true;

    }
}