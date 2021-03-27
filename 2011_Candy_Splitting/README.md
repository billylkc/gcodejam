# Candy Splitting (10pts, 15pts)

Practice Submissions
You have not attempted this problem.
Last updated: Mar 22 2021, 10:04

# PROBLEM ANALYSIS

Sean and Patrick are brothers who just got a nice bag of candy from their parents. Each piece of candy has some positive integer value, and the children want to divide the candy between them. First, Sean will split the candy into two piles, and choose one to give to Patrick. Then Patrick will try to calculate the value of each pile, where the value of a pile is the sum of the values of all pieces of candy in that pile; if he decides the piles don't have equal value, he will start crying.

Unfortunately, Patrick is very young and doesn't know how to add properly. He almost knows how to add numbers in binary; but when he adds two 1s together, he always forgets to carry the remainder to the next bit. For example, if he wants to sum 12 (1100 in binary) and 5 (101 in binary), he will add the two rightmost bits correctly, but in the third bit he will forget to carry the remainder to the next bit:

    1100
  + 0101
  ------
    1001

So after adding the last bit without the carry from the third bit, the final result is 9 (1001 in binary). Here are some other examples of Patrick's math skills:

5 + 4 = 1
7 + 9 = 14
50 + 10 = 56
Sean is very good at adding, and he wants to take as much value as he can without causing his little brother to cry. If it's possible, he will split the bag of candy into two non-empty piles such that Patrick thinks that both have the same value. Given the values of all pieces of candy in the bag, we would like to know if this is possible; and, if it's possible, determine the maximum possible value of Sean's pile.

# Analysis
Analysis
The main step needed to solve this problem is to understand Patrick's strange addition algorithm. For example, can we describe what happens when Patrick adds up several numbers instead of just two? It turns out we can: we should write all numbers in binary, align them by their least significant bit, and write 1 in those positions where we have an odd number of 1 bits in the summands (the numbers being added).

Consider this example: suppose Patrick needs to add 5, 7 and 9 together. First, he adds up 5 and 7 by writing them in binary and adding digit-by-digit without carry, as described in the problem statement:

    101
  + 111
  -----
    010
The result is 010 in binary, which is 2. Now, he adds up 2 and 9:

  0010
+ 1001
------
1011

The result is 1011 in binary, which is 11. It is most instructive to look at what happened to the least significant bit: after adding up two of the numbers, we had a 0 in the least significant bit since both of the summands had a 1 there. However, we have a 1 in the least significant bit of the overall result since the third number had a 1 there as well. It's not hard to see that this generalizes as described above: for any bit, it will be equal to 1 in the overall sum if and only if this bit is set to 1 in an odd number of summands.

Having established that, we can now understand Sean's task better. He needs to separate the given set of numbers into two parts in such a way that for every bit position, either:

In both parts, an odd number of summands have this bit set to 1, so that the corresponding bit in the sum is 1 for both parts, or
In both parts, an even number of summands have this bit set to 1, so that the corresponding bit in the sum is 0 for both parts.
But saying that we need two numbers to be either both odd, or both even, is equivalent to saying that their sum must be even!

That allows us to reformulate Sean's task simply as: he needs to separate the given set of numbers into two parts in such a way that for every bit position, an even number of summands have the bit set to 1 across both parts put together. Suddenly, we understand that this condition doesn't rely on the way we separate the numbers into two parts at all! Either it is true for every bit position that an even number of all summands have the bit set to 1, in which case any separation into two non-empty piles will make Patrick happy, or there is some bit which is 1 in an odd number of summands, in which case there's no way to make Patrick happy.

For example, suppose Sean has pieces of candy with values 5, 7, 9 and 11. If he separates them into 5 and 7 for himself, and 9 and 11 for Patrick, Patrick will add 5 and 7 as 5+7=1012+1112=0102=2, and 9 and 11 as 9+11=10012+10112=00102=2, so Patrick is happy. But even if Sean takes 7, 9 and 11 and leaves just 5 to Patrick, Patrick will add 7, 9 and 11 as 7+9+11=01112+10012+10112=01012=5, so he is still happy! It's not difficult to verify that in all other cases Patrick is happy as well.

All the above reasoning can be made simpler if you notice that Patrick's strange addition is exactly bitwise exclusive or. The condition for Patrick's happiness can be reformulated as the bitwise exclusive or of all candy values being equal to zero. In many programming languages, bitwise exclusive or is already built in with the "^" operator, making this very easy to check!

Now, how does the overall solution for the problem work? First, we need to check if Patrick will be happy - as shown above, this does not depend on Sean's piles. If not, then we just output "NO" for this testcase. If yes, then Sean should maximize his pile, and this is achieved by taking all pieces of candy except the one with the smallest value.

# Input
The first line of the input gives the number of test cases, T. T test cases follow. Each test case is described in two lines. The first line contains a single integer N, denoting the number of candies in the bag. The next line contains the N integers Ci separated by single spaces, which denote the value of each piece of candy in the bag.

# Output
For each test case, output one line containing "Case #x: y", where x is the case number (starting from 1). If it is impossible for Sean to keep Patrick from crying, y should be the word "NO". Otherwise, y should be the value of the pile of candies that Sean will keep.

Limits
1 ≤ T ≤ 100.
1 ≤ Ci ≤ 106.
Memory limit: 1GB.

Small dataset (Test set 1 - Visible)
2 ≤ N ≤ 15.
Time limit: 30 seconds.

Large dataset (Test set 2 - Hidden)
2 ≤ N ≤ 1000.
Time limit: 60 seconds.


# Input

2
5
1 2 3 4 5
3
3 5 6

# Output
Case #1: NO
Case #2: 11
