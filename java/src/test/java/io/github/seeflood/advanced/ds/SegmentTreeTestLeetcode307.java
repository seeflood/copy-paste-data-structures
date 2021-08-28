package io.github.seeflood.advanced.ds;

import org.junit.Assert;
import org.junit.Test;

/**
 * 用线段树解决Leetcode 307题
 */
public class SegmentTreeTestLeetcode307 {


    public static class NumArray {
        SegmentTree<Integer> tree;

        public NumArray(int[] nums) {
            if (nums.length == 0) {
                tree = null;
                return;
            }
            Integer[] arr = new Integer[nums.length];
            for (int i = 0; i < nums.length; i++) {
                arr[i] = nums[i];
            }
            tree = new SegmentTree<>(arr, SegmentTree.Functions.SUM_INTEGER);
        }

        public void update(int i, int val) {
            if (tree != null) {
                tree.set(i, val);
            }
        }

        public int sumRange(int i, int j) {
            if(tree==null){
                return 0;
            }
            return tree.query(i, j);
        }
    }

    @Test
    public void test1() {
        NumArray a = new NumArray(new int[]{1, 3, 5});
        int i = a.sumRange(0, 2);
        Assert.assertEquals(i, 9);
        a.update(1, 2);
        int i2 = a.sumRange(0, 2);
        Assert.assertEquals(i2, 8);

    }
}