package io.github.seeflood.advanced.ds.map;

import org.junit.Test;

public class CounterTest {

    @Test
    public void test1() {
        Counter<Integer> counter = new Counter<>();
        assertTrue(counter.get(1) == 0);
        assertTrue(counter.add(1) == 1);
        assertTrue(counter.get(111) == 0);
        assertTrue(counter.add(1) == 2);
        assertTrue(counter.get(111) == 0);
        assertTrue(counter.substract(111) == -1);
        assertTrue(counter.add(111) == 0);
        for (Integer k : counter) {
            System.out.println(k + "," + counter.get(k));
        }
    }

    /**
     * use Counter to solve https://leetcode-cn.com/problems/count-special-quadruplets/
     */
    @Test
    public void testSolveLeetcode() {
        CounterTest ct = new CounterTest();
        int v = ct.countQuadruplets(new int[] {1, 1, 1, 3, 5});
        assertTrue(v == 4);
    }

    int MAX = 100;

    public int countQuadruplets(int[] nums) {
        int n = nums.length, ans = 0;
        Counter<Integer> single = new Counter<>();
        Counter<Integer> dbl = new Counter<>();
        Counter<Integer> triple = new Counter<>();
        for (int i = 0; i < n; i++) {
            ans += triple.get(nums[i]);
            for (Integer k : dbl) {
                if (k + nums[i] <= MAX) {
                    triple.add(k + nums[i], dbl.get(k));
                }
            }
            for (Integer k : single) {
                if (k + nums[i] <= MAX) {
                    dbl.add(k + nums[i], single.get(k));
                }
            }
            single.add(nums[i]);
        }

        return ans;
    }

    private void assertTrue(boolean b) {
        if (!b) {
            throw new RuntimeException("assertion failed.");
        }
    }
}