package io.github.seeflood.advanced.ds.unionfind;

import org.junit.Assert;
import org.junit.Test;

public class FixedSizeUnionFindTest {
    @Test
    public void testUnionFind() {
        FixedSizeUnionFind<String> uf = new FixedSizeUnionFind<>(10);
        uf.union("a","b");
        uf.union("c","d");
        uf.union("e","f");
        boolean connected = uf.connected("c", "e");
        Assert.assertFalse(connected);
        int i = uf.countComponents();
        Assert.assertEquals(i,3);
        uf.union("g","h");
        uf.union("i","j");
        i = uf.countComponents();
        Assert.assertEquals(i,5);
        uf.union("a","c");
        uf.union("a","e");
        i = uf.countComponents();
        Assert.assertEquals(i,3);
        connected = uf.connected("c", "e");
        Assert.assertTrue(connected);
    }
}