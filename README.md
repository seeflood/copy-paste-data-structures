[comment]: <> (<a href="foo-full_res.png"><picture>)

[comment]: <> (  <source srcset="1.jpg" media="&#40;prefers-color-scheme: dark&#41;" />)

[comment]: <> (  <source srcset="2.jpeg" media="&#40;prefers-color-scheme: light&#41;, &#40;prefers-color-scheme: no-preference&#41;" />)

[comment]: <> (  <img src="2.jpeg" />)

[comment]: <> (</picture></a>)

# Copy-Paste Data Structures
Advanced data structures in Go and Java,which can be easily copy-pasted into your project or competitive programming code.

chinese:

>一个java和Go的高级数据结构库，设计目标是任何数据结构都能单独复制粘贴出来、单独使用，文件和文件之间没有任何互相依赖，以便在打online-judge比赛（比如leetcode周赛，google比赛）的时候能够直接粘贴使用，或者把某个数据结构粘贴到自己的工作项目中做魔改。

## Examples For Golang
- deque
  - [linked_list](go/pkg/deque/linked_list_test.go)
  - [slice_deque](go/pkg/deque/slice_deque_test.go)
- heap
  - [int_heap](go/pkg/heap/int_heap_test.go)
  - [slice_heap](go/pkg/heap/slice_heap_test.go)
- set
  - [bit_set](go/pkg/set/bit_set_test.go)
  - [hash_set](go/pkg/set/hash_set_test.go)
- map
  - [lru_cache](go/pkg/map/test/lru_cache_leetcode146_test.go)
- [union find](go/pkg/union_find/union_find_test.go)
- tree
  - [trie](go/pkg/tree/trie_test.go)
  - [segment tree](go/pkg/tree/segment/segment_tree_test.go)
  - [cartesian tree](go/pkg/tree/cartesian/cartesian_tree_test.go)

## Examples For Java
You can check unit tests in [test directory](java/src/test) as examples.

### Cartesian Tree Example
Let's take [leetcode-1526](https://leetcode-cn.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/) as an example,we can use cartesian-tree to solve it.

```$xslt
 class Solution {
        int total = 0;

        public int minNumberOperations(int[] target) {
            int n = target.length;
            if (n == 1) {
                return target[0];
            }
            Integer[] arr = new Integer[n];
            for (int i = 0; i < n; i++) {
                arr[i] = target[i];
            }
            CartesianTree<Integer> tree = new CartesianTree<>(arr);
//            dfs
            CartesianTree.Node root = tree.getRoot();
            dfs(root, 0);
            return total;
        }

        private void dfs(CartesianTree.Node<Integer> root, int parentV) {
            int diff = (root.getValue() - parentV) ;
            total += diff;
            if (root.hasLeft()) {
                dfs(root.getLeft(),root.getValue());
            }
            if(root.hasRight()){
                dfs(root.getRight(),root.getValue());
            }
        }

         // copy-paste cartesian-tree here
    }

```

### Segment Tree Example
Code using segment tree to solve the problem [leetcode-307](https://leetcode-cn.com/problems/range-sum-query-mutable/) :
```$xslt
    // copy-paste segment tree here
    class NumArray {
        SegmentTree<Integer> tree;

        public NumArray(int[] nums) {
            if (nums.length == 0) {
                tree = null;
                return;
            }
            Integer[] arr = new Integer[nums.length];
            for (int i = 0; i < nums.length; i++) {
                arr[i] = nums[i];
            }
            tree = new SegmentTree<>(arr, SegmentTree.Functions.SUM_INTEGER);
        }

        public void update(int i, int val) {
            if (tree != null) {
                tree.set(i, val);
            }
        }

        public int sumRange(int i, int j) {
            if(tree==null){
                return 0;
            }
            return tree.query(i, j);
        }
    }

```

## Design Goals

### Online-Judge Ready
Users can copy any single data structure(e.g. segment tree) into their solution codes for online-judge(like leetcode) without import any other files.

To achieve this goal,any file in this project imports nothing but those in standard library,and there is no depandency relationship between files in this project.

### Copy-Paste Ready


### Production Ready
As a SDE in a tech company,I used to copy-paste a trie into my project to solve a business problem,because in that case the Trie in Apache-Common-Collections didn't fit the situation.

Adding a new open-source project into the private maven repository in a big company might not be that easy,because you always have to persuade lots of leaders to make approval.

So copy-paste a data structure into your project can be another choice,especially when you want to make some changes to the data structure.
