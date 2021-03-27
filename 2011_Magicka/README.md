# Magicka (10pts, 15pts)

Practice Submissions
You have not attempted this problem.
Last updated: Mar 22 2021, 10:10

# PROBLEM ANALYSIS
Introduction
Magicka™ is an action-adventure game developed by Arrowhead Game Studios. In Magicka you play a wizard, invoking and combining elements to create Magicks. This problem has a similar idea, but it does not assume that you have played Magicka.

Note: "invoke" means "call on." For this problem, it is a technical term and you don't need to know its normal English meaning.

# Problem
As a wizard, you can invoke eight elements, which are the "base" elements. Each base element is a single character from {Q, W, E, R, A, S, D, F}. When you invoke an element, it gets appended to your element list. For example: if you invoke W and then invoke A, (we'll call that "invoking WA" for short) then your element list will be [W, A].

We will specify pairs of base elements that combine to form non-base elements (the other 18 capital letters). For example, Q and F might combine to form T. If the two elements from a pair appear at the end of the element list, then both elements of the pair will be immediately removed, and they will be replaced by the element they form. In the example above, if the element list looks like [A, Q, F] or [A, F, Q] at any point, it will become [A, T].

We will specify pairs of base elements that are opposed to each other. After you invoke an element, if it isn't immediately combined to form another element, and it is opposed to something in your element list, then your whole element list will be cleared.

For example, suppose Q and F combine to make T. R and F are opposed to each other. Then invoking the following things (in order, from left to right) will have the following results:

QF → [T] (Q and F combine to form T)
QEF → [Q, E, F] (Q and F can't combine because they were never at the end of the element list together)
RFE → [E] (F and R are opposed, so the list is cleared; then E is invoked)
REF → [] (F and R are opposed, so the list is cleared)
RQF → [R, T] (QF combine to make T, so the list is not cleared)
RFQ → [Q] (F and R are opposed, so the list is cleared)
Given a list of elements to invoke, what will be in the element list when you're done?

Analysis
This problem can be solved with a simulation. First, we have to remember what elements combine to make other elements. A map of some sort, like a hash map, is a great way of doing this. Next we have to track the opposed elements, remembering that one element can be opposed to multiple other elements; a set of pairs, while not particularly efficient for this purpose, will do the trick.

Finally, the simulation itself. For each character, first we check to see if it combines with the last item on the element list, and combine it if so. If it doesn't combine, then we iterate through the elements already in the list and see if it's opposed to any of them -- if so, we clear the list. Finally, if neither of those conditions was met, we append it to the list. Here is some Pythonesque pseudocode that solves the problem:

# Let combo_list contain all the combinations as 3-letter strs.
# Let opposed_list contain all the opposed elements as 2-letter strs.
# Let invoke be a str containing the elements to invoke.
combos = dict()
opposed = dict()
	for x in combo_list:
combos[x[0] + x[1]] = x[2]
combos[x[1] + x[0]] = x[2]
for x in opposed_list:
opposed.add(x[0] + x[1])
opposed.add(x[1] + x[0])
# Now combos contains a mapping from each pair to the thing it
# creates.  If one of the combinations was "ABC", then
# combos["AB"] = "C" and combos["BA"] = "C".
# opposed is filled in a similar way.

element_list = []
for element in invoke:
# If element_list isn't empty, the last element might combine
# with the element being invoked.
if element_list:
last_two = element_list[-1] + element
if last_two in combos:
element_list[-1] = combos[last_two]
continue

# Now we iterate through element_list to see if anything there
# is opposed to the element being invoked.
wipe_list = False
for e in element_list:
if (e + element) in opposed:
wipe_list = True
if wipe_list:
element_list = []
continue

# There was no combination and no erasing: just append the
# element to the list.
element_list.append(element)

# Input
The first line of the input gives the number of test cases, T. T test cases follow. Each test case consists of a single line, containing the following space-separated elements in order:

First an integer C, followed by C strings, each containing three characters: two base elements followed by a non-base element. This indicates that the two base elements combine to form the non-base element. Next will come an integer D, followed by D strings, each containing two characters: two base elements that are opposed to each other. Finally there will be an integer N, followed by a single string containing N characters: the series of base elements you are to invoke. You will invoke them in the order they appear in the string (leftmost character first, and so on), one at a time.

# Output
For each test case, output one line containing "Case #x: y", where x is the case number (starting from 1) and y is a list in the format "[e0, e1, ...]" where ei is the ith element of the final element list. Please see the sample output for examples.

# Limits
1 ≤ T ≤ 100.
Each pair of base elements may only appear together in one combination, though they may appear in a combination and also be opposed to each other.
No base element may be opposed to itself.
Unlike in the computer game Magicka, there is no limit to the length of the element list.
Memory limit: 1GB.

Small dataset (Test set 1 - Visible)
0 ≤ C ≤ 1.
0 ≤ D ≤ 1.
1 ≤ N ≤ 10.
Time limit: 30 seconds.

Large dataset (Test set 2 - Hidden)
0 ≤ C ≤ 36.
0 ≤ D ≤ 28.
1 ≤ N ≤ 100.
Time limit: 60 seconds.

Sample

# Input

 5
 0 0 2 EA
 1 QRI 0 4 RRQR
 1 QFT 1 QF 7 FAQFDFQ
 1 EEZ 1 QE 7 QEEEERA
 0 1 QW 2 QW


# Output

Case #1: [E, A]
Case #2: [R, I, R]
Case #3: [F, D, T]
Case #4: [Z, E, R, A]
Case #5: []


Magicka™ is a trademark of Paradox Interactive AB. Paradox Interactive AB does not endorse and has no involvement with Google Code Jam.
