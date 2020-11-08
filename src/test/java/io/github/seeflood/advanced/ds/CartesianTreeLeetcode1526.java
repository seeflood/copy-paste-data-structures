package io.github.seeflood.advanced.ds;

import org.junit.Assert;
import org.junit.Test;

/**
 * solve leetcode 1526 with cartesian tree.
 */
public class CartesianTreeLeetcode1526 {

    class Solution {
        int total = 0;

        public int minNumberOperations(int[] target) {
            int n = target.length;
            if (n == 1) {
                return target[0];
            }
            Integer[] arr = new Integer[n];
            for (int i = 0; i < n; i++) {
                arr[i] = target[i];
            }
            CartesianTree<Integer> tree = new CartesianTree<>(arr);
//            dfs
            CartesianTree.Node root = tree.getRoot();
            dfs(root, 0);
            return total;
        }

        private void dfs(CartesianTree.Node<Integer> root, int parentV) {
            int diff = (root.getValue() - parentV) ;
            total += diff;
            if (root.hasLeft()) {
                dfs(root.getLeft(),root.getValue());
            }
            if(root.hasRight()){
                dfs(root.getRight(),root.getValue());
            }
        }
    }

    @Test
    public void test1(){
        int i = new Solution().minNumberOperations(new int[]{1, 2, 3, 2, 1});
        Assert.assertEquals(i,3);
    }

    @Test
    public void test2(){
        int i = new Solution().minNumberOperations(new int[]{3,1,1,2});
        Assert.assertEquals(i,4);
    }
}
