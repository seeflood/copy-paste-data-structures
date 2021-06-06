package io.github.seeflood.advanced.ds;

/**
 * 线段树
 *
 * @param <T>
 */
public class SegmentTree<T> {
    /**
     * 内置的聚合函数
     *
     * @param <T>
     */
    public static class Functions<T> {
        public static AggregateFunction<Integer> MIN_INTEGER = new AggregateFunction<Integer>() {
            @Override
            public Integer aggregate(Integer a, Integer b) {
                return Math.min(a, b);
            }

            @Override
            public Integer update(Integer old, Integer diff) {
                return old + diff;
            }

            @Override
            public Integer getEmpty() {
                return 0;
            }
        };
        public static AggregateFunction<Integer> MAX_INTEGER = new AggregateFunction<Integer>() {
            @Override
            public Integer aggregate(Integer a, Integer b) {
                if (a == null) {
                    return b;
                }
                if (b == null) {
                    return a;
                }
                return Math.max(a, b);
            }

            @Override
            public Integer update(Integer old, Integer diff) {
                if (old == null) {
                    return diff;
                }
                if (diff == null) {
                    return old;
                }
                return old + diff;
            }

            @Override
            public Integer getEmpty() {
                return 0;
            }
        };
        public static AggregateFunction<Integer> SUM_INTEGER = new AggregateFunction<Integer>() {

            @Override
            public Integer aggregate(Integer a, Integer b) {
                if (a == null && b == null) {
                    return 0;
                }
                if (a == null) {
                    return b;
                }
                if (b == null) {
                    return a;
                }
                return a + b;
            }

            @Override
            public Integer update(Integer old, Integer diff) {
                if (old == null && diff == null) {
                    return 0;
                }
                if (old == null) {
                    return diff;
                }
                if (diff == null) {
                    return old;
                }
                return old + diff;
            }

            @Override
            public Integer getEmpty() {
                return 0;
            }
        };
    }


    /**
     * 聚合函数接口
     *
     * @param <T>
     */
    public interface AggregateFunction<T> {
        T aggregate(T a, T b);

        T update(T old, T diff);

        T getEmpty();
    }

    private final AggregateFunction<T> func;
    private final Object[] st;
    private final Object[] lazy;
    private final boolean[] lazyIsUpdate;
    private final int n;

    public SegmentTree(T[] arr, AggregateFunction<T> func) {
        this.func = func;
        this.n = arr.length;
        int size = 2 * nextPowerOf2(n) - 1;
        this.st = new Object[size];
        this.lazy = new Object[size];
        this.lazyIsUpdate = new boolean[size];
        constructTree(arr, 0, n - 1, 0);
    }

    public T query(int from, int to) {
        return query(from, to, 0, n - 1, 0);
    }

    private T query(int qlo, int qhi, int lo, int hi, int idx) {
        //        check arguments
        if (qlo > qhi) {
            throw new IllegalArgumentException("from can't be larger than to");
        }
//        1. cover
        if (qlo <= lo && hi <= qhi) {
            return castIfNotNull(st[idx]);
        }
//        2. non-overlap
        if (qhi < lo || hi < qlo) {
            return func.getEmpty();
        }
//        3. overlap
        int mid = getMid(lo, hi);
        return func.aggregate(query(qlo, qhi, lo, mid, getLeftChild(idx)), query(qlo, qhi, mid + 1, hi, getRightChild(idx)));
    }

    public void set(int index, T value) {
        update(index, index, 0, n - 1, 0, value, false);
    }

    public void set(int from, int to, T value) {
        //        check arguments
        if (from > to) {
            throw new IllegalArgumentException("from can't be larger than to");
        }
        update(from, to, 0, n - 1, 0, value, false);
    }

    private void update(int qlo, int qhi, int lo, int hi, int idx, T value, boolean isUpdate) {
        if (lo > hi) {
            return;
        }
        propagateLazyToChildren(lo, hi, idx);

//        1. cover
        if (qlo <= lo && hi <= qhi) {
            if (isUpdate) {
                st[idx] = func.update(castIfNotNull(st[idx]), value);
            } else {
//                set value
                st[idx] = value;
            }
//             store lazy flag
            lazy[idx] = value;
            lazyIsUpdate[idx] = isUpdate;
            return;
        }
//        2. non-overlap
        if (qhi < lo || hi < qlo) {
            return;
        }
//        3. overlap
        int mid = getMid(lo, hi);
        int leftChild = getLeftChild(idx);
        update(qlo, qhi, lo, mid, leftChild, value, isUpdate);
        int rightChild = getRightChild(idx);
        update(qlo, qhi, mid + 1, hi, rightChild, value, isUpdate);
        st[idx] = func.aggregate(castIfNotNull(st[leftChild]), castIfNotNull(st[rightChild]));
    }

    private void propagateLazyToChildren(int lo, int hi, int idx) {
        if (lo >= hi || lazy[idx] == null) {
            return;
        }
        int leftChild = getLeftChild(idx);
        int rightChild = getRightChild(idx);
//        apply lazy operation to children
        if (lazyIsUpdate[idx]) {
            st[leftChild] = func.update(castIfNotNull(st[leftChild]), castIfNotNull(lazy[idx]));
            st[rightChild] = func.update(castIfNotNull(st[rightChild]), castIfNotNull(lazy[idx]));
        } else {
            st[leftChild] = lazy[idx];
            st[rightChild] = lazy[idx];
        }
//        set children's lazy flag
        lazy[leftChild] = lazy[idx];
        lazyIsUpdate[leftChild] = lazyIsUpdate[idx];
        lazy[rightChild] = lazy[idx];
        lazyIsUpdate[rightChild] = lazyIsUpdate[idx];
//        remove flag since has been used.
        lazy[idx] = null;
    }

    private int getRightChild(int idx) {
        return 2 * idx + 2;
    }

    private int getLeftChild(int idx) {
        return 2 * idx + 1;
    }

    private void constructTree(T[] arr, int lo, int hi, int idx) {
        if (lo > hi) {
            return;
        }
        if (lo == hi) {
//            set
            st[idx] = arr[lo];
            return;
        }
        int mid = getMid(lo, hi);
        int left = 2 * idx + 1;
        constructTree(arr, lo, mid, left);
        int right = 2 * idx + 2;
        constructTree(arr, mid + 1, hi, right);
        st[idx] = func.aggregate(castIfNotNull(st[left]), castIfNotNull(st[right]));
    }

    private T castIfNotNull(Object o) {
        if (o == null) {
            return null;
        }
        return (T) o;
    }

    private int getMid(int lo, int hi) {
        return lo + (hi - lo) / 2;
    }

    private int nextPowerOf2(int i) {
        if (i <= 0) {
            return 0;
        }
        if ((i & (i - 1)) == 0) {
            return i;
        }
        while ((i & (i - 1)) != 0) {
            i &= (i - 1);
        }
        return i << 1;
    }

}
