NAME: Julia Rieger
HW: 3
DATE: 10-19-23 to 10-26-23
==================

Programs Files:
===============
    Part 1:
        scope.go
    Part 2:
        closures.go
    Part 3:
        iterator.go

How to Compile + Run:
=====================
    Part 1:
       go run scope.go
    Part 2:
       go run closures.go
    Part 3:
       go run iterator.go
	
Reflection: 
===========
     Part 1 was not difficult because we had directly gone over similar scenarios in class and it was easy to imagine cases that would be different
     in a static versus dynamic scoping environment. Part 2 was hard for me to understand. I had to read the instructions many times over and a big 
     portion of the time I spent on it was because I didn't get what I was supposed to be doing. I think I eventually got it though, and considering,
     coding it was not the biggest challenge. Part 3 was a little more difficult than part 2 because there was a little less instruction, I believe 
     I was able to make the iterator generator function correctly but I'm still not confident in whether my test cases are implemented correctly,
     I was confused on how to make a loop with the iterators.

Discussion:
===========
     Part 1: Proving Go uses static scoping
     We can say variables are bound to their global identifier, in static scoping this global identifier is the "top-level" scope/environment,
     and in dynamic scoping this global identifier is the most recent scope/environment on the runtime stack. In static scoping the compiler
     searches for bindings in first the current scope, then in global scope, then in decreasingly smaller scopes. This is why the output of
     scope.go is "value of a when g() is called: 1". When the compiler was in the f() stack frame and needed to return a, it first searched the
     local scope (f()) for a binding to a, of which it did not find one. It then searched the global scope for a, where it found a = 1 in line
     12. Because static scoping prioritizes global scope and scope size over the most recent stack, it returned 1 instead of 2. Thus g() returns
     1 and 1 is printed. Were this written in a dynamically scoped language, we would search for variable binding in first the current
     block, and then all the successive functions called in order of most to least recent. This means while we are in the f() stack frame
     and need to return a, we would first search the local scope (f()) not finding a value for a, and next search the most recent function in 
     the stack, which is g() (because g() calls f()). In g(), a is initialized to 2, and therefore 2 would be returned from f(), and then g(),
     and 2 would be printed in main. In this program, if 1 were printed that means the global scope was accessed/searched first, and if 2 were printed
     that means the most recent function call was accessed/searched first. Because Go printed 1, this means the global scope was accessed first and
     therefore Go used static scoping.

     Part 2:
     My tests and the output show that the nextInSeries functions do not interact because I provided examples (specifically lines 76-77) where I called
     a nextInSeries function that was made before the functions running (in lines 63-75) were made. When I called the maker function to create 3 more
     nextInSeries Functions in lines 60-62, that did not interfere with the current progress/status of the index variable for the original 2 functions
     in lines 38 and 48. Additionally when I called additional tests, I switched back and forth between calling odd, geometric, and arithmetic functions,
     and in each case the current state/index was preserved because the terms they returned were correct in the order of their specific series. In short,
     this proved the index information was preserved even while making other nextInSeries functions. I can gave an unbounded number of nextInSeries functions
     as long as I can create new functions/variables to assign them to. The maker function can return an unlimited amount of functions which store their
     own current index values, therefore no matter how many nextInSeries are created, the index value data will not be lost and therefore not limit 
     how many functions can be made.

     Part 3:
     I think my iterator is inferior to the range command because since it is written in Go (with my limited knowledge of Go), when you want to loop through
     a slice with the iterator you have to write a bulky loop. Additionally you may need to use extra variables because of this loop. If you are not looking
     to use a loop and simply want to make a single iteration, you have to use new variable names when declaring the element and validity with the iterator
     method call, because the Go compiler gets mad at the reuse of e and v. The range command is much more readable and probably has optimizations that make
     it faster than my program. They will both get the job done. 

     When you change the contents of the slice that my iterator works on, the iterator will use the updated contents of the slice if the indices are the same,
     for example if you replace a slice of [1, 2, 3, 4, 5] with the contents [6, 7, 8, 9, 10]. Because you are replacing the same number of indices, my iterator
     function doesn't trip up because the index value being incremented will be the same. Every time the iterator returns the next value it runs the function new,
     which is why changing the slice contents after the iterator is called/bound does not change the value of the elements being returned (again as long as indices
     aren't changed).

     For this same reason, my iterator fails to print newly appended items. The appended items do not appear if the index becomes greater than the length of
     the slice when the iterator was originally called/bound. However if some items are deleted from the slice, making it so the index of the appended item(s)
     is less than or equal to the length of the original slice, the newly appended item values will be returned correctly, since the iterator will still reach
     those new items.

Approximate Hours worked:
=========================
    4hrs

Known Bugs or Limitations:
==========================
    N/A???