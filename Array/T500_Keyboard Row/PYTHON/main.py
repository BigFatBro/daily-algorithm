from typing import List
class Solution:
    def findWords(self, words: List[str]) -> List[str]:
        ans = []
        t = "12210111011122000010020202"
        for word in words:
            idx = t[ord(word[0].lower()) - ord('a')]
            if all(t[ord(c.lower()) - ord('a')] == idx  for c in word):
                ans.append(word)
        return ans