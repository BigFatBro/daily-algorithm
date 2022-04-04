import java.util.HashMap;
import java.util.Map;
import java.util.List;
import java.util.ArrayList;

class Solution {
    public String[] findRestaurant(String[] list1, String[] list2) {
        Map<String,Integer> map = new HashMap<>();
        for (int i = 0; i < list1.length; i++) {
            map.put(list1[i], i);
        }

        int minIndexSum = Integer.MAX_VALUE;
        List<String> ret = new ArrayList<String>();
        for (int i = 0; i < list2.length; i++) {
            if (map.containsKey(list2[i])) {
                int indexSum = i + map.get(list2[i]);
                if (indexSum == minIndexSum) {
                    ret.add(list2[i]);
                }else if(indexSum < minIndexSum){
                    minIndexSum = indexSum;
                    ret.clear();
                    ret.add(list2[i]);
                }
            }
        }
        return ret.toArray(new String[ret.size()]);

    }
}