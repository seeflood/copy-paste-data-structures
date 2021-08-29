package io.github.seeflood.advanced.ds.map;

import org.junit.Assert;
import org.junit.Test;

import static org.junit.Assert.*;

public class LFUCacheTest {

    @Test
    public void test1() {
        LFUCache<Integer, Integer> lfu = new LFUCache(2);
        lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
        lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
        Integer v = lfu.get(1);// return 1
        Assert.assertTrue(v == 1);
        // cache=[1,2], cnt(2)=1, cnt(1)=2
        lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
        // cache=[3,1], cnt(3)=1, cnt(1)=2
        v = lfu.get(2);// return null (not found)
        Assert.assertNull(v);
        v = lfu.get(3);      // return 3
        Assert.assertTrue(v == 3);
        // cache=[3,1], cnt(3)=2, cnt(1)=2
        lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
        // cache=[4,3], cnt(4)=1, cnt(3)=2
        Assert.assertNull(lfu.get(1));      // return null (not found)
        Assert.assertTrue(lfu.get(3) == 3);      // return 3
        // cache=[3,4], cnt(4)=1, cnt(3)=3
        Assert.assertTrue(lfu.get(4) == 4);    // return 4
        // cache=[3,4], cnt(4)=2, cnt(3)=3
    }
}