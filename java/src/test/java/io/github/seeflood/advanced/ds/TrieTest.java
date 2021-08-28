package io.github.seeflood.advanced.ds;

import org.junit.Assert;
import org.junit.Test;

public class TrieTest {
    @Test
    public void test1() {
        Trie<Integer> trie = new Trie<>();
        trie.put("apple", 11);
        Integer v = trie.get("apple");
        Assert.assertTrue(v == 11);
        v = trie.get("app");
        Assert.assertTrue(v == null);
        boolean b = trie.startsWith("app");// return True
        Assert.assertTrue(b);
        trie.put("app", 1);
        v = trie.get("app");
        Assert.assertTrue(v == 1);
        b = trie.startsWith("apples");
        Assert.assertFalse(b);
        String pre = trie.longestPrefixOf("apples");
        Assert.assertTrue(pre.equals("apple"));
        pre = trie.longestPrefixOf("apple");
        Assert.assertTrue(pre.equals("apple"));
        v = trie.remove("apple");
        Assert.assertTrue(v == 11);
        pre = trie.longestPrefixOf("apples");
        Assert.assertTrue(pre.equals("app"));
        v = trie.get("apple");
        Assert.assertTrue(v == null);
        v = trie.get("app");
        Assert.assertTrue(v == 1);
    }

}