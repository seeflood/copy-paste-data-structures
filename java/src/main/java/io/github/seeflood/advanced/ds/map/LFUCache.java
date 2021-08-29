package io.github.seeflood.advanced.ds.map;

import java.util.HashMap;
import java.util.Map;

public class LFUCache<K, V> {
    static class Node<K, V> {
        K          k;
        V          v;
        int        freq = 0;
        Node<K, V> prev, next;

        Node(K k, V v) {
            this.k = k;
            this.v = v;
        }
    }

    static class LinkedList<K, V> {
        Node<K, V> dummy = new Node<>(null, null);
        int        cnt   = 0;

        LinkedList() {
            dummy.next = dummy;
            dummy.prev = dummy;
        }

        void remove(Node<K, V> node) {
            node.prev.next = node.next;
            node.next.prev = node.prev;
            node.next = null;
            node.prev = null;
            cnt--;
        }

        int size() {
            return cnt;
        }

        public void addLast(Node<K, V> node) {
            Node<K, V> tail = dummy.prev;
            tail.next = node;
            node.prev = tail;
            dummy.prev = node;
            node.next = dummy;
            cnt++;
        }

        public Node<K, V> pollFirst() {
            Node<K, V> first = dummy.next;
            dummy.next = first.next;
            first.next.prev = dummy;
            first.next = null;
            first.prev = null;
            cnt--;
            return first;
        }
    }

    private int                            cap;
    private Map<K, Node<K, V>>             k2node    = new HashMap<>();
    private Map<Integer, LinkedList<K, V>> freq2List = new HashMap<>();
    private int                            minFreq   = 1;

    public LFUCache(int capacity) {
        this.cap = capacity;
    }

    public V get(K key) {
        if (!k2node.containsKey(key)) {
            return null;
        }
        Node<K, V> node = k2node.get(key);
        refresh(node);
        return node.v;
    }

    private void refresh(Node<K, V> node) {
        // remove node in old freq list
        if (node.freq > 0) {
            LinkedList<K, V> list = freq2List.get(node.freq);
            list.remove(node);
            if (list.size() == 0) {
                freq2List.remove(node.freq);
                if (minFreq == node.freq) {
                    minFreq++;
                }
            }
        }
        // add the node into the new freq list
        minFreq = Math.min(minFreq, ++node.freq);
        LinkedList<K, V> list = freq2List.get(node.freq);
        if (list == null) {
            list = new LinkedList<>();
            freq2List.put(node.freq, list);
        }
        list.addLast(node);
    }

    public void put(K key, V value) {
        if (this.cap <= 0) {
            return;
        }
        if (k2node.containsKey(key)) {
            Node<K, V> node = k2node.get(key);
            node.v = value;
            refresh(node);
            return;
        }
        if (k2node.size() == cap) {
            // have to remove the LFU node
            LinkedList<K, V> list = freq2List.get(minFreq);
            Node<K, V> toRemove = list.pollFirst();
            if (list.size() == 0) {
                freq2List.remove(minFreq);
            }
            k2node.remove(toRemove.k);
        }
        // add new node
        Node<K, V> toAdd = new Node<>(key, value);
        refresh(toAdd);
        k2node.put(key, toAdd);
    }

    public int size() {
        return k2node.size();
    }
}