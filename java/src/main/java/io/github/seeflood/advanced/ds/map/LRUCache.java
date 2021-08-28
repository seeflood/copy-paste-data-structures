package io.github.seeflood.advanced.ds.map;

import java.util.Iterator;
import java.util.LinkedHashMap;
import java.util.Map;

public class LRUCache<K, V> {
    LinkedHashMap<K, V> map;
    int                 cap;

    public LRUCache(int capacity) {
        this.cap = capacity;
        this.map = new LinkedHashMap<>(cap, 0.75f, true);
    }

    public V get(K key) {
        if (!map.containsKey(key)) {
            return null;
        }
        return map.get(key);
    }

    public void put(K key, V value) {
        if (cap <= 0) {
            return;
        }
        map.put(key, value);
        if (map.size() > cap) {
            Iterator<Map.Entry<K, V>> itr = map.entrySet().iterator();
            itr.next();
            itr.remove();
        }
    }
}