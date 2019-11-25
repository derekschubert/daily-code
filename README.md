# Daily Code
### Daily code challenge solutions &amp;&amp; mock-ups

_Note: These challenges are either from various online sources (ie: dailycodingchallenges.com / leetcode.com / etc) or some problem that I want to figure out personally._

Officially started on Nov 4, 2019 (daily habit made public as a way to push some things to Github ;))

- Day 1: Find all different rectangles in a 2d grid, including diagonals. Some points may be missing.
  - source: [Clément Mihailescu on YouTube](https://www.youtube.com/watch?v=EuPSibuIKIg)
  - solution: [day 1](day-1/rectanglesInGrid.js)

- Day 2: Add two numbers (each represented in a reverse-order linked list) and return it as a linked list.
  - source: [LeetCode](https://leetcode.com/problems/add-two-numbers/)
  - solution: [day 2](day-002/addTwoNumbersInLinkedList.go)

- Day 3: Do any two numbers in a given list add up to equal value k? (determined in O(n) time)
  - source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Google)
  - solution: [day 3](day-003/doTwoNumbersAddUpToK.go) (two solutions based on an input of either a sorted or unsorted list)

- Day 4: Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i
  - source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Uber)
  - solution: [day 4](day-004/productOfAllButSelf.go) (3 solutions: with division & w/o division)

- Day 5: Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.
  - source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Google)
  - solution: [day 5](day-005/serializeBinarySearchTree.js)

- Day 6: Given an array of integers, find the first missing positive integer.
  - source: [Daily Coding Problem (starting to see a trend here...)](https://www.dailycodingproblem.com/) (problem by Stripe)
  - solution: [day 6](day-006/findLowestMissingPositiveInteger.go)
  
- Day 7: Merge k sorted linked lists and return it as one sorted list.
  - source: [LeetCode](https://leetcode.com/problems/merge-k-sorted-lists/)
  - solution: [day 7](day-007/mergeKLinkedLists.go)

- Day 8: Find how much water would be trapped given an array of heights (each index representing an equal width)
  - source: [LeetCode](https://leetcode.com/problems/trapping-rain-water/)
  - solution: [day 8](day-008/trappedWater.go)

- Day 9: Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded - assuming all messages are decodable (ie: 001 not allowed)
  - source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Facebook)
  - solution: [day 9](day-009/decodeNumbers.go)

- Day 10: Given n non-negative integers where each integer represents a point at coordinate (index, int), n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
  - source: [LeetCode](https://leetcode.com/problems/container-with-most-water/)
  - solution: [day 10](day-010/maxWater.go)

- Day 11: Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.
  - source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Airbnb)
  - solution: [day 11](day-011/largestSumOfNonAdjacentNumbers.go) (2 solutions - dynamic (very fast) & recursive)

- Day 12: Doubled up today! 1 - Given a string, find the length of the longest substring without repeating characters. 2 - Given a string, find the length of the longest substring without repeating/mirrored substrings. (see solution file for examples)
  - source: [LeetCode](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
  - solution: [day 12](day-012/longestSubstringWithoutRepeating.go)

- Day 13: Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.
  - source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Apple)
  - solution: [day 13](day-013/jobScheduler.go)

- Day 14: There exists a staircase with N steps, and you can climb up any number of steps from a set of positive integers X. Given N and X, write a function that returns the number of unique ways you can climb the staircase.
  - source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Amazon)
  - solution: [day 14](day-014/waysUpStaircase.go)

- Day 15:	Given an integer k and a string s, find the length of the longest substring that contains at most k distinct characters. (similar but different enough to Day 12 problems)
	- source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Amazon)
  - solution: [day 15](day-015/longestSubstringOfKChars.js) (2 solutions: sliding window/pointers & brute force)

- Day 16: Cheat day! No specific challenge, but did a full implementation of a Binary Search Tree.
  - source: [me :)](https://github.com/derekschubert)
  - solution: [day 16](day-016/bst.js)

- Day 17: Estimate π to 3 decimal places using a Monte Carlo method.  
	- source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Google)
  - solution: [day 17](day-017/monteCarloFindPi.go)

- Day 18: Given a stream of elements too large to store in memory, pick a random element from the stream with uniform probability.
	- source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Facebook)
  - solution: [day 18](day-018/uniformProbability.go)

- Day 19: Given an array of integers and a number k, where 1 <= k <= length of the array,	compute the sum of each subarray of length k.
	- source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Google)
  - solution: [day 19](day-019/totalOfSubarrays.go)

[insert weekend algorithm break here! :D]

- Day 20: Given two singly linked lists that intersect at some point, find the intersecting node. The lists are non-cyclical. 
	- source: [Daily Coding Problem](https://www.dailycodingproblem.com/) (problem by Google)
  - solution: [day 20](day-020/linkedListIntersection.go)