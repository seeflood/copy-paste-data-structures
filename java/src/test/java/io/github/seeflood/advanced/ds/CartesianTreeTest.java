package io.github.seeflood.advanced.ds;

import junit.framework.TestCase;
import org.junit.Assert;
import org.junit.Test;

public class CartesianTreeTest extends TestCase {


    @Test
    public void testConstruct(){
        Integer[]arr=new Integer[]{1,2,3};
        CartesianTree<Integer> itgTree = new CartesianTree<>(arr);
        CartesianTree.Node root = itgTree.getRoot();
        Assert.assertEquals(root.getValue(),1);
        Assert.assertEquals(root.getIdx(),0);
        Assert.assertFalse(root.hasLeft());
        Assert.assertTrue(root.hasRight());
        CartesianTree.Node left = root.getLeft();
        Assert.assertNull(left);
        CartesianTree.Node right = root.getRight();
        Assert.assertEquals(right.getValue(),2);
        Assert.assertEquals(right.getIdx(),1);
        Assert.assertFalse(right.hasLeft());
        Assert.assertTrue(right.hasRight());
        Assert.assertEquals(right.getRight().getValue(),3);
        Assert.assertEquals(right.getRight().getIdx(),2);
        String[] strArr=new String[]{"1","2","3"};
        CartesianTree<String> strTree = new CartesianTree<>(strArr);
        CartesianTree.Node root1 = strTree.getRoot();
        Assert.assertEquals(root1.getValue(),"1");
    }
}