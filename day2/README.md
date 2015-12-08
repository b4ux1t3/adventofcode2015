#Day Two: Day 2: I Was Told There Would Be No Math

[Link.](http://adventofcode.com/day/2)

#The Problem
The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?

#Instructions

The only thing you have to do for this program is to make sure you have a file that is named `input` in the same directory as the executable (or as the `.go` file if you use `go run`).

Feel free to mess around with different files.

`input.copy` is a copy of the original input file which allows you to mess around with `input` without losing the original solution.
