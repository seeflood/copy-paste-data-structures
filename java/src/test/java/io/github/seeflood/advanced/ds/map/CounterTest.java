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
    }

    private void assertTrue(boolean b) {
        if (!b) {
            throw new RuntimeException("assertion failed.");
        }
    }
}