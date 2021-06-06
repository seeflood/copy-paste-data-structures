package io.github.seeflood.advanced.ds;

import java.util.ArrayDeque;
import java.util.Comparator;
import java.util.Deque;

/**
 * 笛卡尔树,可以使用CartesianTree.getRoot()来得到根，以便遍历树
 *
 * @param <T>
 */
public class CartesianTree<T> {
    public static class Node<T> {
        private int from = -1;
        private int to = -1;
        private int idx = -1;
        private T value = null;
        private Node<T> left = null, right = null;

        public Node(int idx, T value, int from, int to) {
            this.idx = idx;
            this.value = value;
            this.from = from;
            this.to = to;
        }

        public Node(int idx, T value) {
            this.idx = idx;
            this.value = value;
        }

        public Node() {
        }

        public int getFrom() {
            return from;
        }

        public int getTo() {
            return to;
        }

        public int getIdx() {
            return idx;
        }

        public T getValue() {
            return value;
        }

        public boolean hasLeft() {
            lazyAddRange();
            return left != null;
        }

        private void lazyAddRange() {
            if (left != null && left.from == -1) {
                left.from = this.from;
                left.to = this.idx - 1;
            }
            if (right != null && right.from == -1) {
                right.from = this.idx + 1;
                right.to = this.to;
            }
        }

        public boolean hasRight() {
            lazyAddRange();
            return right != null;
        }

        public Node<T> getLeft() {
            lazyAddRange();
            return left;
        }

        public Node<T> getRight() {
            lazyAddRange();
            return right;
        }
    }

    final Comparator<T> DEFAULT_COMPARATOR = new Comparator<T>() {
        @Override
        public int compare(T o1, T o2) {
            return ((Comparable) o1).compareTo(o2);
        }
    };

    T[] arr;
    Node dummy = new Node();

    /**
     * Build Cartesian Tree,e.g. arr=[5,2,3,1,4],
     * then tree will look like:<br>
     * <pre>
     *          1
     *      2       4
     *   5    3
     *   </pre>
     *
     * @param arr
     */
    public CartesianTree(T[] arr) {
        this.arr = arr;
        init(arr, DEFAULT_COMPARATOR);

    }

    /**
     * Build Cartesian Tree,e.g. arr=[5,2,3,1,4],
     * then tree will look like:<br>
     * <pre>
     *          1
     *      2       4
     *   5    3
     *   </pre>
     *
     * @param arr
     */
    public CartesianTree(T[] arr, Comparator<T> comparator) {
        this.arr = arr;
        init(arr, comparator);
    }

    private void init(T[] arr, Comparator<T> comparator) {
        if (arr == null || comparator == null) {
            throw new NullPointerException();
        }
        if (arr.length == 0) {
            throw new IllegalArgumentException("arr is empty,can't build a tree on it.");
        }
        //         build tree with stack
        dummy.right = new Node(0, arr[0]);
        Deque<Node<T>> stack = new ArrayDeque<>();
        stack.addLast(dummy.right);
        int n = arr.length;
        for (int i = 1; i < n; i++) {
//            walk up stack to find smaller node
            Node<T> prev = null;
            while (!stack.isEmpty() && comparator.compare(stack.peekLast().value, arr[i]) > 0) {
                prev = stack.pollLast();
            }
            Node<T> toAdd = new Node<>(i, arr[i]);
            if (!stack.isEmpty()) {
                stack.peekLast().right = toAdd;
            }else{
                dummy.right=toAdd;
            }
            stack.addLast(toAdd);
            toAdd.left = prev;
        }

//       lazy add range index
        dummy.right.from = 0;
        dummy.right.to = n - 1;
    }

    public T[] getArr() {
        return arr;
    }

    public Node getRoot() {
        return dummy.right;
    }
}
