package io.github.seeflood.advanced.ds.map;

import org.junit.Assert;
import org.junit.Test;

import static org.junit.Assert.*;

public class LRUCacheTest {

    @Test
    public void test1() {
        LRUCache<Integer, Integer> lRUCache = new LRUCache<>(2);
        lRUCache.put(1, 1); // cache is {1=1}
        lRUCache.put(2, 2); // cache is {1=1, 2=2}
        Integer v = lRUCache.get(1);// return 1
        Assert.assertTrue(v == 1);
        lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
        v = lRUCache.get(2);// returns null (not found)
        Assert.assertTrue(v == null);
        lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
        v = lRUCache.get(1);// return null (not found)
        Assert.assertTrue(v == null);
        v = lRUCache.get(3);    // return 3
        Assert.assertTrue(v == 3);
        v = lRUCache.get(4);    // return 4
        Assert.assertTrue(v == 4);
    }
}