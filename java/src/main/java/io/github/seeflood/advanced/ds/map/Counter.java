package io.github.seeflood.advanced.ds.map;

import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

/**
 * An easy-to-use counter like that in python
 *
 * @param <T>
 */
public class Counter<T> implements Iterable<T> {
    private Map<T, Integer> m = new HashMap<>();

    public Counter() {
    }

    public Counter(T[] array) {
        if (array == null) {
            return;
        }
        for (int i = 0; i < array.length; i++) {
            add(array[i]);
        }
    }

    public Counter(List<T> list) {
        if (list == null) {
            return;
        }
        for (int i = 0; i < list.size(); i++) {
            add(list.get(i));
        }
    }

    /**
     * @param t
     * @return counter value after this addition
     */
    public int add(T t) {
        return add(t, 1);
    }

    public int add(T key, int diff) {
        int c;
        if (!m.containsKey(key)) {
            c = diff;
        } else {
            c = m.get(key) + diff;
        }
        m.put(key, c);
        return c;
    }

    /**
     * counter-1.  It can be negative after substraction
     *
     * @param t
     * @return counter value after this operation
     */
    public int substract(T t) {
        return add(t, -1);
    }

    public int get(T t) {
        if (!m.containsKey(t)) {
            return 0;
        }
        return m.get(t);
    }

    @Override
    public Iterator<T> iterator() {
        return m.keySet().iterator();
    }
}