package io.github.seeflood.advanced.ds.impl;

import io.github.seeflood.advanced.ds.UnionFind;

import java.util.HashMap;
import java.util.Map;

public class FixedSizeUnionFind<T> implements UnionFind<T> {
    int n;
    int components;
    int[] father;
    int[] size;
    Object[] idxToElm;
    Map<T, Integer> elmToIdx = new HashMap<>();

    public FixedSizeUnionFind(int size) {
        this.n = size;
        this.components = size;
        this.father = new int[size];
        this.size = new int[size];
        this.idxToElm = new Object[size];
        for (int i = 0; i < size; i++) {
            father[i] = i;
            this.size[i] = 1;
        }
    }

    @Override
    public void union(T a, T b) {
        int aRoot = findRootIdx(a);
        int bRoot = findRootIdx(b);
        if (aRoot == bRoot) {
            return;
        }
        if (size[aRoot] >= size[bRoot]) {
            size[aRoot] += size[bRoot];
            father[bRoot] = aRoot;
        } else {
            size[bRoot] += size[aRoot];
            father[aRoot] = bRoot;
        }
        components--;
    }

    @Override
    public T findRoot(T a) {
        int rootIdx = findRootIdx(a);
        return (T) idxToElm[rootIdx];
    }

    @Override
    public boolean connected(T a, T b) {
        return findRootIdx(a) == findRootIdx(b);
    }

    private int findRootIdx(T a) {
        int i = getIdx(a);
        int cur = i;
        while (father[cur] != cur) {
            int old = cur;
            cur = father[cur];
            father[old] = father[cur];
        }
        return cur;
    }

    private int getIdx(T a) {
        if (elmToIdx.containsKey(a)) {
            return elmToIdx.get(a);
        }
        int i = elmToIdx.size();
        if (i >= n) {
            throw new IllegalArgumentException("Container is full,can not afford new element!");
        }
        elmToIdx.put(a, i);
        idxToElm[i] = a;
        return i;
    }

    @Override
    public int countComponents() {
        return components - (n - elmToIdx.size());
    }
}
