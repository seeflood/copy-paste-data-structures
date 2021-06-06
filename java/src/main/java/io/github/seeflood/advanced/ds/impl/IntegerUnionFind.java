package io.github.seeflood.advanced.ds.impl;

import io.github.seeflood.advanced.ds.UnionFind;

public class IntegerUnionFind implements UnionFind<Integer> {
    int n;
    int components;
    int[] father;
    int[] size;

    public IntegerUnionFind(int capacity) {
        this.n = capacity;
        this.components = capacity;
        this.father = new int[capacity];
        this.size = new int[capacity];
        for (int i = 0; i < n; i++) {
            this.father[i] = i;
            this.size[i] = 1;
        }
    }

    @Override
    public void union(Integer a, Integer b) {
        Integer aRoot = findRoot(a);
        Integer bRoot = findRoot(b);
        if (aRoot.equals(bRoot)) {
            return;
        }
        if (size[aRoot] < size[bRoot]) {
            father[aRoot] = bRoot;
            size[bRoot] += size[aRoot];
        } else {
            father[bRoot] = aRoot;
            size[aRoot] += size[bRoot];
        }
        components--;
    }

    @Override
    public Integer findRoot(Integer element) {
        if (element == null) {
            throw new NullPointerException();
        }
        if (element < 0 || element >= n) {
            throw new IllegalArgumentException("element is out of bound.");
        }
        int cur = element;
        while (father[cur] != cur) {
            int old = cur;
            cur = father[cur];
            father[old] = father[cur];
        }
        return cur;
    }

    @Override
    public boolean connected(Integer a, Integer b) {
        //         not possible for NPE
        return findRoot(a).equals(findRoot(b));
    }

    @Override
    public int countComponents() {
        return components;
    }
}
