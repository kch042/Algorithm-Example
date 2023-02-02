# 0-1 knapsack problem (0-1 ks)

## Review
- Greedy method can be used to solve fractional ks problem
- Dynamic programming can be used to solved 0-1 ks only for small bag weight limit and the weight of items must be integers

The most generic way to solve this problem is to use brute force method which take O(2^n) time
However, we could improve the performance by branch and bound method (B&B)

## Method
Draw the decision tree of ks problem, in which each leaf node represents a permutation (candidate solution) to the problem.

We can imagine that we are searching the tree for the optimal answer W may decide not to go further down some subtrees if we know all of its leaves has total value less than
the largest value we have found so far, thus saving some running time.

To know if a subtree is worth searching, we could compute an upper bound for the total value of all leaves in the subtree
the above op is called "Bounding".

If the computed upper bound is less than the largest value found so far which means that all leaves has value less than best found so far
then, we may disregard the subtree.

Else, we need to add this subtree into our consideration of exploration which is called "Branching".
