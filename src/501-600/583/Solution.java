/**
 *
 * 583. Delete Operation for Two Strings

 Given two words word1 and word2, find the minimum number of steps required to make word1 and word2 the same, where in each step you can delete one character in either string.

 Example 1:
 Input: "sea", "eat"
 Output: 2

 Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".

 Note:
 The length of given words won't exceed 500.
 Characters in given words can only be lower-case letters.

 */



/**
 * Created by zeming.wang on 2018/1/8.
 */
class Solution {

    public static void main(String[] args) {
        System.out.println(minDistance("aaaa","aabb"));
    }

    public static int minDistance(String word1, String word2) {
        int min = 1000;

        for(int i = 0; i < word1.length(); i++){
            int index_2 = 0;
            int match_num = 0;
            while(index_2 < word2.length()){
                if(word1.charAt(i) == word2.charAt(index_2)){
                    match_num++;
                }
            }
        }

        return 1;
    }
}
