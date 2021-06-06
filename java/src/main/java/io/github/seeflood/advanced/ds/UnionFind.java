package io.github.seeflood.advanced.ds;

public interface UnionFind<T> {

    void union(T a,T b);

    /**
     * find root
     * @param a
     * @return
     */
    T findRoot(T a);

    boolean connected(T a,T b);

    int countComponents();
}
