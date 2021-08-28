package io.github.seeflood.advanced.ds.unionfind;

/**
 * UnionFind for int
 */
public class IntUnionFind {
    private int   cnt;
    private int[] parent;
    private int[] size;

    public IntUnionFind(int size) {
        this.cnt = size;
        parent = new int[size];
        this.size = new int[size];
        for (int i = 0; i < size; i++) {
            parent[i] = i;
            this.size[i] = 1;
        }
    }

    public void union(int a, int b) {
        int aroot = findRoot(a);
        int broot = findRoot(b);
        if (aroot == broot) {
            return;
        }
        cnt--;
        if (size[aroot] <= size[broot]) {
            parent[aroot] = broot;
            size[broot] += size[aroot];
        } else {
            parent[broot] = aroot;
            size[aroot] += size[broot];
        }
    }

    /**
     * find root
     *
     * @param a
     * @return
     */
    public int findRoot(int a) {
        int root = a;
        while (parent[root] != root) {
            root = parent[root];
        }
        // path compression
        while (parent[a] != root) {
            int tmp = parent[a];
            parent[a] = root;
            a = tmp;
        }
        return root;
    }

    public boolean connected(int a, int b) {
        if (a >= parent.length || b >= parent.length || a < 0 || b < 0) {
            return false;
        }
        return findRoot(a) == findRoot(b);
    }

    public int countComponents() {
        return cnt;
    }
}