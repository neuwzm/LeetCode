/**
 * Created by zeming.wang on 2018/11/13.
 */


import java.util.Date;
import java.util.HashMap;
import java.util.Map;

/**
 *
 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

 示例:

 给定 nums = [2, 7, 11, 15], target = 9

 因为 nums[0] + nums[1] = 2 + 7 = 9
 所以返回 [0, 1]

 */
public class Sum {

    public static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();

        for (int i = 0; i< nums.length; i++){
            int self = nums[i];
            Integer couple = map.get(self);
            if(couple != null){
                return new int[]{couple, i};
            } else {
                map.put(target - self, i);
            }

        }
        return null;
    }

    public static void main(String[] args) {
        int[] a = new int[]{1,2,1,4,5,6,7,8,9,0};
        int[] r = twoSum(a, 4);
        System.out.println(r[0]);
        System.out.println(r[1]);


    }

}
