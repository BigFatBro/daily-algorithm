var search = function(nums, target) {
    let mid,left=0, right = nums.length-1;
    while(left<=right){
        mid = left+((right - left) >> 1);
        if (nums[mid]>target) {
            right = mid-1;
        }else if (nums[mid]<target) {
            left = mid +1;
        }else{
            return mid;
        }

    }
    return -1;

};

let nums = [-1,0,3,5,9,12], target = 9;
console.log("output:", search(nums, target))