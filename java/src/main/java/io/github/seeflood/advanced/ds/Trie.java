package io.github.seeflood.advanced.ds;

public class Trie<V> {
    private class Node<V> {
        private V         value;
        private Node<V>[] next;
        private int       cnt = 0;

        public Node(V value, int R) {
            this.value = value;
            next = new Node[R];
        }
    }

    public static final int     DEFAULT_R = 256;
    private             int     R;
    private             Node<V> root;

    public Trie() {
        R = DEFAULT_R;
        root = new Node<>(null, R);
    }

    public Trie(int r) {
        R = r;
        root = new Node<>(null, R);
    }

    public V get(String key) {
        if (key == null || key.length() == 0) {
            return root.value;
        }
        Node<V> node = getNode(root, key);
        if (node == null) {
            return null;
        }
        return node.value;
    }

    private Node<V> getNode(Node<V> root, String key) {
        Node<V> parent = root;
        int i = 0;
        int len = key.length();
        while (parent != null && i < len) {
            char c = key.charAt(i);
            parent = parent.next[c];
            i++;
        }
        if (i < len) {
            return null;
        }
        return parent;
    }

    public void put(String key, V value) {
        if (key == null || key.length() == 0) {
            root.value = value;
            root.cnt++;
            return;
        }
        // find the leaf node
        Node<V> parent = root;
        int len = key.length();
        for (int i = 0; i < len; i++) {
            parent.cnt++;
            char c = key.charAt(i);
            if (parent.next[c] == null) {
                parent.next[c] = new Node<>(null, R);
            }
            parent = parent.next[c];
        }
        // put value
        parent.cnt++;
        parent.value = value;
    }

    public String longestPrefixOf(String key) {
        if (key == null) {
            return null;
        }
        Node<V> parent = root;
        int i = 0, len = key.length(), max = 0;
        while (parent != null && i < len) {
            if (parent.value != null) {
                max = i;
            }
            char c = key.charAt(i);
            parent = parent.next[c];
            i++;
        }
        if (parent != null && parent.value != null) {
            max = len;
        }
        return key.substring(0, max);
    }

    public V remove(String key) {
        if (key == null || key.length() == 0) {
            V v = root.value;
            root.value = null;
            root.cnt--;
            return v;
        }
        // find the leaf node
        Node<V> parent = root;
        int i = 0, len = key.length();
        while (parent != null && i < len) {
            parent.cnt--;
            char c = key.charAt(i);
            parent = parent.next[c];
            i++;
        }
        // remove value
        if (parent == null) {
            return null;
        }
        parent.cnt--;
        V v = parent.value;
        parent.value = null;
        return v;
    }

    public boolean startsWith(String prefix) {
        if (prefix == null || prefix.length() == 0) {
            return true;
        }
        Node<V> node = getNode(root, prefix);
        if (node == null) {
            return false;
        }
        return node.cnt > 0;
    }
}