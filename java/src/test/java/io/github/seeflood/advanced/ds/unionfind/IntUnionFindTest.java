package io.github.seeflood.advanced.ds.unionfind;

import org.junit.Assert;
import org.junit.Test;

public class IntUnionFindTest {
    @Test
    public void testUnionFind() {
        IntUnionFind uf = new IntUnionFind(10);
        uf.union(0, 1);
        uf.union(2, 3);
        uf.union(4, 5);
        boolean connected = uf.connected(2, 4);
        Assert.assertFalse(connected);
        int i = uf.countComponents();
        Assert.assertEquals(i, 7);
        uf.union(6, 7);
        uf.union(8, 9);
        i = uf.countComponents();
        Assert.assertEquals(i, 5);
        uf.union(0, 2);
        uf.union(0, 4);
        i = uf.countComponents();
        Assert.assertEquals(i, 3);
        connected = uf.connected(2, 4);
        Assert.assertTrue(connected);
    }

}