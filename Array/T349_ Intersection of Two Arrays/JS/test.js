var set_intersection = function(set1,set2){
    if (set1.size > set2.size) {
        return set_intersection(set2, set1);
    }
    const ans = new Set();
    //遍历较小的集合，判断其是否在大集合中
    for (const num of set1) {
        if (set2.has(num)) {
            ans.add(num);
        }
    }
    return [...ans];
    
}

var intersection = function(nums1, nums2){
    //先去重,存储到集合中
    const set1 = new Set(nums1);
    const set2 = new Set(nums2);
    return set_intersection(set1, set2);
}

const nums1 = [1,2,2,1];
const nums2 = [2,2];
console.log(intersection(nums1, nums2))