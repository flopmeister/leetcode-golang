## String
|#  | Title           |  Solution       |  Time           | Space           | Difficulty    | Tag          | Note| 
|---|---------------- | --------------- | --------------- | --------------- | ------------- |--------------|-----|
8 | [String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/) | [golang](./golang/string_to_integer_atoi.go) | O(n) | O(1) | Easy ||
12 | [Integer to Roman](https://leetcode.com/problems/integer-to-roman) | [golang](./golang/integer_to_roman.go) | O(1) | O(1) | Medium ||
13 | [Roman to Integer](https://leetcode.com/problems/roman-to-integer) | [golang](./golang/roman_to_integer.go) | O(n) | O(1) | Easy ||
273 | [Integer to String](https://leetcode.com/problems/integer-to-english-words) | [golang](./golang/integer_to_english_words.go) |  O(n/10) | O(1)| Hard || 
686 | [Repeated String Match](https://leetcode.com/contest/leetcode-weekly-contest-52/problems) | [goalng](./golang/repeated_string_match.go) | O(n)|O(1)|Easy ||
383 | [Ransom Note](https://leetcode.com/problems/ransom-note) | [golang](./golang/ransom_note.go) | O(m+n) | O(m) | Easy ||
22 | [Generate Parentheses](https://leetcode.com/problems/generate-parentheses) | [golang](./golang/generate_parentheses.go) | TBA | TBA | Medium || 不是最优解。
28 | [Implement strStr()](https://leetcode.com/problems/implement-strstr) | [golang](./golang/implement_strStr.go) | O(n^2) | O(1) | Easy || 下次使用KMP
38 | [Count and Say](https://leetcode.com/problems/count-and-say) | [golang](./golang/count_and_say.go) | O(n^2) | O(n) | Easy || 用一个pointer记录char出现次数 
344 | [Reverse String](https://leetcode.com/problems/reverse-string) | [golang](./golang/reverse_string.go) | O(n) | O(1) | Easy ||
657 | [Judge Route Circle](https://leetcode.com/problems/judge-route-circle) | [golang](./golang/judge_route_circle.go) | O(n) | O(1) | Easy ||
58 | [Length of Last Word](https://leetcode.com/problems/length-of-last-word) | [golang](./golang/length_of_last_word.go) | O(n) | O(1) | Easy ||

## DP
|#  | Title           |  Solution       |  Time           | Space           | Difficulty    | Tag          | Note| 
|---|---------------- | --------------- | --------------- | --------------- | ------------- |--------------|-----|
689 | [Maximum Sum of 3 Non-Overlapping Subarrays](https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/) | [golang](./golang/max_sum_of_three_subarrays.go)| O(n) | O(n) | Hard ||
32 | [Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses) | [golang](./golang/longest_valid_parentheses.go) | O(n) | O(n) | Hard ||
121 | [Best Time to Buy and Sell Stock ](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [golang](./golang/best_time_to_buy_and_sell_stock.go)| O(n) | O(1) | Easy ||
714 | [Best Time to Buy and Sell Stock with Transaction Fee ](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) | [golang](./golang/best_time_to_buy_and_sell_with_transaction_fee.go)| O(n) | O(1) | Medium ||
123 | [Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii) | [golang](./golang/best_time_to_buy_and_sell_stock_iii.go)| O(n) | O(n) | Easy ||
188 | [Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv) | [golang](./golang/best_time_to_buy_and_sell_stock_iv.go)| O(kn) | O(k) | Hard ||
309 | [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown) | [golang](./golang/best_time_to_buy_and_sell_stock_with_cool_down.go) [python](./python/best_time_to_buy_and_sell_stock_with_cool_down.py)| O(n) | O(1) | Medium ||
303 | [Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable) | [golang](./golang/range_sum_query_immutable.go.go)| O(1) | O(n) | Easy ||
304| [Range Sum Query 2D - Immutable](https://leetcode.com/problems/range-sum-query-2d-immutable) | [golang](./golang/range_sum_query_2D_immutable.go.go)| O(1) | O(mn) | Medium ||
198| [House Robber](https://leetcode.com/problems/house-robber) | [golang](./golang/housr_robber.go) | O(n)| O(1) | Easy ||
70 | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs) | [golang](./golang/climbing_stairs.go) | O(n)| o(n)| Easy ||
53 | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/description/) | [golang](./golang/maximum_subarray.go) |O(n)|O(n)|Easy ||
264 | [Ugly Number II](https://leetcode.com/problems/ugly-number-ii) | [golang](./golang/ugly_number_ii.go) | O(n) | O(1) | Medium ||
139 | [Word Break](https://leetcode.com/problems/word-break) | [golang](./golang/word_break.go)| O(n^2) | O(n) | Medium ||
140 | [Word Break II](https://leetcode.com/problems/word-break-ii) | [golang](./golang/word_break_ii.go)| O(n^2) | O(n) | Hard ||3 solutions in this file
64 | [Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum) | [golang](./golang/minimum_path_sum.go) | O(m*n) | O(m*n) | Medium || Can be further improved to space O(1)
62 | [Unique Paths](https://leetcode.com/problems/unique-paths) | [golang](./golang/unique_paths.go) |O(m*n) | O(m*n) | Medium || Can be further improved to space O(m)
53 | [Unique Paths II](https://leetcode.com/problems/unique-paths-ii) | [golang](./golang/unique_paths_ii.go) | O(m*n) | O(m*n) || can be further improved to space O(m)
494 | [Target Sum](https://leetcode.com/problems/target-sum) | [golang](./golang/target_sum.go) | O(n^2) |O(n) | Medium || two solutions are provided
322 | [Coin Change](https://leetcode.com/problems/coin-change) | [golang](./golang/coin_change.go) | O(s*n) | O(s) | Medium ||
312 | [Burst Balloons](https://leetcode.com/problems/burst-balloons) | [golang](./golang/burst_balloons.go) | O(n^3) | O(n) | Hard || 参考了discuss 解出题目
673 | [Number of Longest Increasing Subsequence](https://leetcode.com/problems/number-of-longest-increasing-subsequence) | [golang](./golang/number_of_longest_increasing_subsequence.go) | O(n^2) | O(n) | Medium || 
691 | [Stickers to Spell Word](https://leetcode.com/problems/stickers-to-spell-word) | [golang](./golang/stickers_to_spell_word.go) | TBA | TBA | Hard || 还没做出来。。。 

## Array
|#  | Title           |  Solution       |  Time           | Space           | Difficulty    | Tag          | Note| 
|---|---------------- | --------------- | --------------- | --------------- | ------------- |--------------|-----|
122 | [Best Time to Buy and Sell Stock II ](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | [golang](./golang/best_time_to_buy_and_sell_stock_ii.go)| O(n) | O(1) | Easy ||
670 | [Maximum Swap](https://leetcode.com/problems/maximum-swap) | [golang](./golang/maximum_swap.go) | O(n) | O(n) | Medium || 两次循环很巧妙
674 | [Longest Continuous Increasing Subsequence](https://leetcode.com/problems/longest-continuous-increasing-subsequence) | [golang](./golang/longest_continuous_increasing_sequence.go) | O(1) | O(1) | Easy ||
26 | [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array) | [golang](./golang/remove_duplicates_from_sorted_array.go) | O(n) | O(1) | Easy ||
27 | [Remove Element](https://leetcode.com/problems/remove-element) | [golang](./golang/remove_element.go) | O(n) | O(1) | Easy|| 
31 | [Next Permutation](https://leetcode.com/problems/next-permutation/description/) | [golang](./golang/next_permutation.go) | O(n) | O(1) | Medium || 
33 | [Search in Rotated Array](https://leetcode.com/problems/search-in-rotated-sorted-array) | [golang](./golang/search_in_rotated_array.go) | O(lg(n)) | O(1) | Medium | Binary Search |
34 | [Search for a Range](https://leetcode.com/problems/search-for-a-range) | [golang](./golang/search_for_a_range.go) | O(lg(n)) | O(1) | Medium | Binary Search |
35 | [Search Insert Position](https://leetcode.com/problems/search-insert-position) | [golang](./golang/search_insert_position.go) | O(lg(n)) | O(1) | Easy | Binary Search |
36 | [Valid Soduku](https://leetcode.com/problems/valid-sudoku) | [golang](./golang/valid_soduku.go) | O(n^2) | O(n^2) | Medium ||
39 | [Combination Sum](https://leetcode.com/problems/combination-sum) | [golang](./golang/combination_sum.go) | TBA | TBA | Medium ||
40 | [Combination Sum II](https://leetcode.com/problems/combination-sum-ii) | [golang](./golang/combination_sum_ii.go) | TBA | O(n) | Medium || 
41 | [First Missing Positive](https://leetcode.com/problems/first-missing-positive) | [golang](./golang/first_missing_positive.go) | O(n) | O(1) | Hard || 这个sort 有些特殊，刚好index跟value 得match
42 | [Trapping Water](https://leetcode.com/problems/trapping-rain-water) | [golang](./golang/trapping_water.go) | O(n^2) | O(n) | Hard | 'Array' 'Two Pointers' 'Stack' | 不是最优解，https://leetcode.com/problems/trapping-rain-water/solution/ 讲解很不错
45 | [Jump Game II](https://leetcode.com/problems/jump-game-ii) | [golang](./golang/jump_game_ii.go) | O(n) | O(1) | Hard | 'BFS' |
55 | [Jump Game](https://leetcode.com/problems/jump-game) | [golang](./golang/jump_game.go) | O(n) | O(1) | Medium | DP |
46 | [Permutations](https://leetcode.com/problems/permutations) | [golang](./golang/permutations.go) | O(n) | O(n) | Medium ||
47 | [Permutations II](https://leetcode.com/problems/permutations-ii) | [golang](./golang/permutations_ii.go) | O(n) | O(n) | Medium ||
48 | [Rotate Image](https://leetcode.com/problems/rotate-image) | [golang](./golang/rotate_image.go) | O(n) | O(1) | Medium || 衍生：anti-clockwise rotation problem.
49 | [Group Anagrams](https://leetcode.com/problems/group-anagrams) | [golang](./golang/group_anagrams.go) | O(n*glogg) | O(n) | Medium || 
561 | [Array Partition I](https://leetcode.com/problems/array-partition-i) | [golang](./golang/array_partition_i.go) | O(n) | O(1) | Easy ||
54 | [Spiral Matrix](https://leetcode.com/problems/spiral-matrix) | [golang](./golang/spiral_matrix.go) | O(n) | O(n) | Medium ||
59 | [Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii) | [golang](./golang/spiral_matrix_ii.go) | O(n) | O(n) | Medium || 2D matrix transverse 问题 
56 | [Merge Intervals](https://leetcode.com/problems/merge-intervals) | [golang](./golang/merge_intervals.go) | O(nlgn) | O(1) | Medium | Sort | implement golang sort interface 即可
57 | [Inset Interval](https://leetcode.com/problems/insert-interval) | [golang](./golang/insert_interval.go) | O(n) | O(1) | Hard | Sort 或者遍历 | solution 2 来自leetcode 感觉更好
217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate) | [golang](./golang/contains_duplicate.go) | O(n) | O(n) | Easy | HashTable |

## Bit Manipulation
|#  | Title           |  Solution       |  Time           | Space           | Difficulty    | Tag          | Note| 
|---|---------------- | --------------- | --------------- | --------------- | ------------- |--------------|-----|
461 | [Hamming Distance ](https://leetcode.com/problems/hamming-distance) | [golang](./golang/hamming_distance.go)| O(4n) | O(1) | Easy ||
136 | [Single Number](https://leetcode.com/problems/single-number) | [golang](./golang/single_number.go) | O(n) | O(1) | Easy |Array|

## Tree
|#  | Title           |  Solution       |  Time           | Space           | Difficulty    | Tag          | Note| 
|---|---------------- | --------------- | --------------- | --------------- | ------------- |--------------|-----|
617 | [Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/description/) | [golang](./golang/merge_two_binary_trees.go)| O(n) | O(n) | Easy ||

## Math
|#  | Title           |  Solution       |  Time           | Space           | Difficulty    | Tag          | Note| 
|---|---------------- | --------------- | --------------- | --------------- | ------------- |--------------|-----|
263 | [Ugly Number](https://leetcode.com/problems/ugly-number) | [golang](./golang/ugly_number.go) | O(n) | O(1) | Easy ||
66 | [Plus One](https://leetcode.com/problems/plus-one) | [golang](./golang/plus_one.go) | O(n) | O(1) | Easy | Array |
67 | [Add Binary](https://leetcode.com/problems/add-binary) | [golang](./golang/add_binary.go) | O(n) | O(1) | Easy| String |
458 | [Poor Pigs](https://leetcode.com/problems/poor-pigs) | [golang](./golang/poor_pigs.go) | O(1) | O(1) | Easy| ...|

## LinkedList
|#  | Title           |  Solution       |  Time           | Space           | Difficulty    | Tag          | Note| 
|---|---------------- | --------------- | --------------- | --------------- | ------------- |--------------|-----|
24 | [Linked List](https://leetcode.com/problems/swap-nodes-in-pairs) | [golang](./golang/swap_nodes_in_pairs.go) | O(n) | O(1) | Medium ||
61 | [Rotate List](https://leetcode.com/problems/rotate-list) | [golang](./golang/rotate_list.go) | O(n) | O(1) | Medium | Two Pointers |
23 | [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists) | [golang](./golang/merge_k_sorted_lists.go) | O(nlgn) | O(n) | HARD|| Can be further reduced to O(nlgk) and O(1)

## Brainteaser
292 | [Nim Game](https://leetcode.com/problems/nim-game) | [golang](./golang/nim_game.go) | O(1) | O(0) | Easy ||