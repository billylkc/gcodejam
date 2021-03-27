# Bot Trust (10pts, 10pts)

Practice Submissions
You have not attempted this problem.
Last updated: Mar 22 2021, 09:52

# PROBLEM ANALYSIS
Problem
Blue and Orange are friendly robots. An evil computer mastermind has locked them up in separate hallways to test them, and then possibly give them cake.

Each hallway contains 100 buttons labeled with the positive integers {1, 2, ..., 100}. Button k is always k meters from the start of the hallway, and the robots both begin at button 1. Over the period of one second, a robot can walk one meter in either direction, or it can press the button at its position once, or it can stay at its position and not press the button. To complete the test, the robots need to push a certain sequence of buttons in a certain order. Both robots know the full sequence in advance. How fast can they complete it?

For example, let's consider the following button sequence:

   O 2, B 1, B 2, O 4

   Here, O 2 means button 2 in Orange's hallway, B 1 means button 1 in Blue's hallway, and so on. The robots can push this sequence of buttons in 6 seconds using the strategy shown below:

   Time | Orange           | Blue
   -----+------------------+-----------------
   1  | Move to button 2 | Stay at button 1
   2  | Push button 2    | Stay at button 1
   3  | Move to button 3 | Push button 1
   4  | Move to button 4 | Move to button 2
   5  | Stay at button 4 | Push button 2
   6  | Push button 4    | Stay at button 2
   Note that Blue has to wait until Orange has completely finished pushing O 2 before it can start pushing B 1.

# Analysis
BAM! Robots gave us 6 extra seconds of cooperation. Good job, robots!
   -- Cave Johnson (Portal 2)


	Hopefully your robots were more focused on teamwork than Cave Johnson's were, or you would never get all the buttons pressed. Just be glad there were no mashy spike plates or potatoes to contend with! Perhaps in the finals...

	Anyway, if you think about this problem from the perspective of a single robot, the strategy should be pretty intuitive: always move towards the next button and then push it as soon as it comes up in the sequence.

	So the most natural solution to this problem is a straight simulation:

	Keep track of which buttons have been pressed, and look ahead in the sequence to figure out which button is next for each robot.
	One second at a time, have each robot move towards its next button.
	Once it gets to the button, the robot should push it if it's next in the sequence, and just wait otherwise.
	At most 100 buttons needs to be pressed altogether and the distance between buttons is at most 100, so this solution will run plenty fast.

If you get stuck on the implementation, remember that you can see other contestants' source code. There is no better way to learn!

# Input
The first line of the input gives the number of test cases, T. T test cases follow.

Each test case consists of a single line beginning with a positive integer N, representing the number of buttons that need to be pressed. This is followed by N terms of the form "Ri Pi" where Ri is a robot color (always 'O' or 'B'), and Pi is a button position.

# Output
For each test case, output one line containing "Case #x: y", where x is the case number (starting from 1) and y is the minimum number of seconds required for the robots to push the given buttons, in order.

# Limits
1 ≤ Pi ≤ 100 for all i.
Memory limit: 1GB.

# Small dataset (Test set 1 - Visible)
1 ≤ T ≤ 20.
1 ≤ N ≤ 10.
Time limit: 30 seconds.

# Large dataset (Test set 2 - Hidden)
1 ≤ T ≤ 100.
1 ≤ N ≤ 100.
Time limit: 60 seconds.

Sample

# Input
3
4 O 2 B 1 B 2 O 4
3 O 5 O 8 B 100
2 B 2 B 1

# Output
Case #1: 6
Case #2: 100
Case #3: 4
