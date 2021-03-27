# GoroSort (10pts, 20pts)

Practice Submissions
You have not attempted this problem.
Last updated: Mar 22 2021, 10:13

# PROBLEM ANALYSIS
Problem
Goro has 4 arms. Goro is very strong. You don't mess with Goro. Goro needs to sort an array of N different integers. Algorithms are not Goro's strength; strength is Goro's strength. Goro's plan is to use the fingers on two of his hands to hold down several elements of the array and hit the table with his third and fourth fists as hard as possible. This will make the unsecured elements of the array fly up into the air, get shuffled randomly, and fall back down into the empty array locations.

Goro wants to sort the array as quickly as possible. How many hits will it take Goro to sort the given array, on average, if he acts intelligently when choosing which elements of the array to hold down before each hit of the table? Goro has an infinite number of fingers on the two hands he uses to hold down the array.

More precisely, before each hit, Goro may choose any subset of the elements of the array to freeze in place. He may choose differently depending on the outcomes of previous hits. Each hit permutes the unfrozen elements uniformly at random. Each permutation is equally likely.

# Input
The first line of the input gives the number of test cases, T. T test cases follow. Each one will consist of two lines. The first line will give the number N. The second line will list the N elements of the array in their initial order.

# Output
For each test case, output one line containing "Case #x: y", where x is the case number (starting from 1) and y is the expected number of hit-the-table operations when following the best hold-down strategy. Answers with an absolute or relative error of at most 10-6 will be considered correct.

Limits
1 ≤ T ≤ 100;
The second line of each test case will contain a permutation of the N smallest positive integers.
Memory limit: 1GB.

Small dataset (Test set 1 - Visible)
1 ≤ N ≤ 10;
Time limit: 30 seconds.

Large dataset (Test set 2 - Hidden)
1 ≤ N ≤ 1000;
Time limit: 60 seconds.

Analysis
The Solution
This problem is very mathematical in nature, requiring a lot of thought and only a little code to solve correctly. We put it as a problem on the Qualifying Round to give you something different to try without having to worry about time pressure or advancement. We hope you enjoyed it!

For an arbitrary array A, let n(A) be the number of elements that are not already in the correct position. It turns out that the answer you are looking for is simply n(A). Once you realize this, it is easy to calculate with a simple loop, but how do you prove it?


The Proof
Let's first show that the expected number of hits is never more than n(A). Suppose Goro always holds down only those elements that are already in position, and then he randomly permutes the rest. Let x(A) be the expected number of hits required for him to sort A using this strategy.

Lemma: x(A) = n(A) for all A.

We prove this by induction on n(A). If n(A) = 0, then the array is already sorted, and we're done. To set up the induction, let's suppose we have proven the lemma already for smaller values of n(A) and we are now trying to prove it for A. Let pt be the probability that exactly t elements are still out of position after the first hit, and let x't be the expected number of hits required in this case. We make three observations:

p0 * 0 + p1 * 1 + p2 * 2 + ... + pN * N = N - 1. In English, this is saying that the expected number of elements that are out of position after the first hit is exactly N - 1, or equivalently, the expected number of elements that are put into position by the first hit is exactly 1. This follows from "linearity of expectation": Goro is permuting N elements; each one has probability exactly 1/N of ending up in the correct position, and hence, the expected number of elements that end up in the correct position is N * 1/N = 1.
x't = t for t ≤ N - 1. This is true by the inductive hypothesis.
x'N = x(A). If no elements are put into the correct position by the first hit, then we will just randomly permute them all again in the next step, so nothing has changed, and hence x'N = x(A).
Now let's write down a formula for x(A):

1 + p0 * x'0 + p1 * x'1 + ... + pN * x'N
= 1 + p0 * 0 + p1 * 1 + ... + pN * N + pN * (x(A) - N)
= N + pN * (x(A) - N),

which simplifies to (N - x(A)) * (1 - pN) = 0. Since pN < 1, we must have x(A) = N, and the lemma is proven.


To complete the proof, we need to calculate y(A), the expected number of hits required for Goro to sort A if he uses the (still unknown) optimal strategy. Since y(A) ≤ x(A) by definition, we have already proven y(A) ≤ n(A).

To conversely prove n(A) ≤ y(A), it would be nice to just extend the proof of the previous lemma. There is one big technical issue though: it is possible for n(A) to go up if Goro doesn't hold down enough elements, and so it is tricky to set up an induction on n(A). We'll resolve this by having a separate proof for this part, and this time use a slightly different induction hypothesis. We can then follow the previous argument very closely.

Lemma 2: Let K be a non-negative integer. Then for any k ≤ K, the statement y(A) = k is equivalent to the statement n(A) = k.

We will prove this by induction on K. Both y(A) = 0 and n(A) = 0 are equivalent to the array already being sorted, so the K = 0 case is clear. Now, let's suppose we have proven the lemma for K already, and are trying to prove it for K+1. Choose A such that y(A) is the smallest possible value larger than K and consider the optimal strategy for Goro. Let T be the number of elements that are either (a) not in the correct position in A, or (b) permuted when Goro hits the table. Define pi and x'i as we did before. Note that T ≥ n(A) ≥ K+1 by the inductive hypothesis.

As in the previous lemma, we can now prove the following:

p0 * 0 + p1 * 1 + p2 * 2 + ... pT * T ≥ T - 1.
x'i = i for i ≤ K. This follows directly from the inductive hypothesis.
x'i ≥ y(A) for i > K. By the inductive hypothesis, n(A') > K implies y(A') > K, which then implies y(A') ≥ y(A).
As before, we can now write y(A) as

1 + p0 * x'0 + p1 * x'1 + ... pT * x'T
≥ 1 + (p0 * 0 + p1 * 1 + ... + pT * T) + (pK+1 + pK+2 + ... + pT) * (y(A) - T)
≥ T + (pK+1 + pK+2 + ... + pT) * (y(A) - T)

which simplifies to (y(A) - T) * (1 - pK+1 - ... - pT) ≥ 0. The second term has to be positive (if not, then y(A) ≥ min(x'K+1, x'K+2, ...) + 1, which contradicts the third bullet point above is therefore impossible), so we must have y(A) ≥ T ≥ n(A) ≥ K+1. Equality holds only if n(A) = K+1. The first lemma guarantees y(A) ≤ x(A) ≤ K+1 in this case, and the proof is complete!


# Comments
If you go over the proof carefully, you can see there are two things Goro needs to do in order to be optimal. (1) He needs to always hold down elements that are already in the correct position, and (2) he needs to ensure that for each element x that is permuted, the element in x's correct position is also permuted. This means he actually has some choice about what to do.
On a programming contest, of course you do not need to work through a formal proof to implement a correct solution. The best contestants can solve this kind of problem by looking at small examples to see a pattern, and then using intuitive reasoning to see what's going on without formalizing everything. Mastering this kind of reasoning is a difficult art though!

# Sample

Input
 3
 2
 2 1
 3
 1 3 2
 4
 2 1 4 3


Output


Case #1: 2.000000
Case #2: 2.000000
Case #3: 4.000000


Explanation
In test case #3, one possible strategy is to hold down the two leftmost elements first. Elements 3 and 4 will be free to move. After a table hit, they will land in the correct order [3, 4] with probability 1/2 and in the wrong order [4, 3] with probability 1/2. Therefore, on average it will take 2 hits to arrange them in the correct order. After that, Goro can hold down elements 3 and 4 and hit the table until 1 and 2 land in the correct order, which will take another 2 hits, on average. The total is then 2 + 2 = 4 hits.
