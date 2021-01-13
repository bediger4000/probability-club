# Daily Coding Problem: Problem #769 [Hard]

This problem was asked by Two Sigma.

Alice wants to join her school's Probability Student Club.
Membership dues are computed via one of two simple probabilistic games.

The first game: roll a die repeatedly.
Stop rolling once you get a five followed by a six.
Your number of rolls is the amount you pay, in dollars.

The second game: same,
except that the stopping condition is a five followed by a five.

Which of the two games should Alice elect to play?
Does it even matter?
Write a program to simulate the two games and calculate their expected value.

## Build and Run

```sh
$ go build game.go
$ ./game 100000 5 6
Playing 100000 games, until 5 followed by 6
Mean fee 42.06
$ ./game 100000 5 5
Playing 100000 games, until 5 followed by 5
Mean fee 42.00
```

## Analysis

The problem statement is a little loose.

* What size die? In this day of sophisticated D&D players, one never knows.
I'm assuming a 6-sided die.
* Does alice have to pay a minimum of $2?
That is, are the last 2 rolls included in the number of rolls?
I assumed they are, so Alice pays a minimum of $2,
but the way the problem is worded it could be 0, $1 or $2.

Since rolls of a die are pretty much independent,
it doesn't matter which game Alice chooses.
A roll of 5, followed by a 6 is just as likely as a 5 followed by a 5.
Each roll of a d6 has 1/6 chance of getting any given value.

[My code](game.go)

Since a series-of-rolls ending condition is a 1/6 event,
followed by a 1/6 event, the probability of rolling twice
is 1/6\*1/6 = 1/36 = 0.0277.
That's almost exactly what my [distribution of rolls](distribution.go)
program shows for proportion of 2-roll series in 500,000 iterations of
the game.
A 3-roll game shows up with a probability of about 0.0231 in my simulation.
A 3-roll game happens if you don't roll a 5, then roll a 5, then roll a 6:
5/6\*1/6\*1/6 = 0.02314.
I think this verifies that my program gives the correct answer.

Looks like Alice could play either game,
and on average pay a $42 membership fee.

## Interview Analysis

I'm not sure this merits a "Hard" rating.
The only thing that's weird is keeping track of whether the last roll was a 5 or not.
I'd give it a "Medium" maybe.

This seems like a problem statement that invites the candidate to ask questions.
It's possible that the interviewer wanted to elicit the candidate's thought processes,
which would be revealed in questions to the interviewer,
rather than seeing actual coding.

I think the candidate should ask questions.
Just assuming probably gets you nowhere with interviewers that ask this question.

## Probability Distribution

![probability distribution](500.png)

This is the results from a 500,000 game simulation.
Image generated with [gnuplot](http://gnuplot.info/) and a [script](mkdist).
The longest game was 554 rolls. Poor Alice.
