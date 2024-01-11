NAME:
=====
	Julia Rieger

Programs Files:
===============
    Part 1:
        lucas/src/main.rs
    Part 2:
        students/src/main.rs
        data file: students/src/students.txt
    Perfect Numbers:
        perfect/src/main.rs
    Stack Depth:
        stack/src/main.rs
	
How to Compile and Run:
=======================
    Part 1:
       cd to the directory -> lucas/src/
       cargo run main.rs
    Part 2:
       cd to the directory -> students/src/
       cargo run main.rs
    Perfect Numbers:
       cd to the directory -> perfect/src/
       cargo run main.rs
    Stack Depth:
       cd to the directory -> stack/src/
       cargo run main.rs
       
Tests:
======
    Part 1:
        to test with L values from 1 through 10, you must manually
        change the L value in the code (line 6: let l = <value>). Then execute the code:
        for l = 1, the output is:
            "The largest term we can reach with 32 bits is 46: 1836311903
             The ratio of the largest term we can compute with 32 bits and the next term is: 0.6180339887498949"

        for l = 10, the output is:
            "The largest term we can reach with 32 bits is 42: 1758135565
             The ratio of the largest term we can compute with 32 bits and the next term is: 0.6180339887498949"
      
    Part 2:
        to test simply follow compile and run instructions
        main.rs will use the file "students.txt"
        students.txt is in the following format:
        <firstname1> <lastname1> <id1>
        <firstname2> <lastname2> <id2>
        ...

        to add additional test cases follow the same format
        valid ids are those which use ASCII values of the student's initials (capital letters)
        ex/ initials "AA" would have the id 6565

    Perfect Numbers Table:
    -5 [9]
    -4 [5, 14, 44, 110, 152, 884, 2144, 8384, 18632, 116624, 8394752, 15370304]
    -3 []
    -2 [3, 10, 136, 32896]
    -1 [2, 4, 6, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 
        131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432]
    Perfect [1, 6, 28, 496, 8128, 33550336]
    1 []
    2 [20, 104, 464, 650, 1952, 130304, 522752, 8382464]
    3 [18]
    4 [12, 70, 88, 1888, 4030, 5830, 321128, 521728, 1848964, 8378368]
    5 []

    Stack Depth:
    Feel free to uncomment test code for rec_1, rec_2, rec_3, or rec_4. Due to the nature of the testing 
    (running until death), only one function can be tested at a time. Therefore only test one function call
    at a time with the auxiliary code directly above when necessary.

Discussion:
===========
    Part 1: Explain how you avoid panic (or death) in your program? 
    The way I computed the next term in the sequence is by adding the two previous terms. Therefore to avoid creating
    a next term that overflowed the memory available for an i32 (2^n-1 -1 = 2147483647) I added a check (lines 19-26) 
    where I casted the two previous terms to i64s, which can certainly hold enough bits to add the two previous terms.
    If the two previous terms exceeded 2147483647 I forced the function to stop computing and return the largest term
    it can handle. The technique I used does not apply for every mathematical function because its based on the fact
    that the next term is calculated using exactly 2 pre-existing terms. This allowed my technique to work nicely
    with my recursive function. For the lucas sequence it was safe to assume that an i64 would be able to hold the sum 
    of the 2 previous terms, but for a different function that may not be the case.

    Max Stack Depth: 
    1. changing the parameter from u32 to u64 did not seem to affect the max depth achievable, because both runs
    averaged around 74,786 - 74,840 stack frames before overflow. However changing the parameter to u128 changed
    noticably, now max frames averages 65,446 - 65,497. A u32 takes 4 bytes of memory, a u64 8 bytes, and a u128 
    16 bytes. Because the growth in memory from u32 to u128 is exponential, it makes sense that the smaller bit 
    options would be more similar and the larger options would noticably decrease the number of max stack frames
    possible. 
    2. Adding an i128 decreased the max frames to averaging 58,172 - 58,210. This decreased the stack frames by about
    16,609. This stands with my previous findings that adding parameters which require more bytes in memory will 
    allow for less stack frames. 
    3. Adding an array of 100 i32 decreased the max frames to averaging 16,358 - 16,370. This decreased the stack 
    frames by about 58,470, which is far more than half of the default u32 recursion. This could be equivalent to 
    adding 100 i32 parameters to the function, so it makes sense why the stack frames decreased by such a dramatic margin.
    4. Adding a Vector to hold i128 decreased the max frames to averaging 47,598 - 47,630. Strangely this was the same whether
    the vector was poplulated with 1000 random numbers or whether it was empty. My best guess for this is that in rust
    an empty vector will allocate memory in the stack to prepare for the programmer to add eleemnets to it, and that 
    allocated amount happened to be enough to freely hold 1000 i128s. While the stack frames decreased by about 27,210,
    it was not nearly to the degree that the array in question 3 decreased the stack frames. This doesn't follow the
    line of reasoning we have been using up until now because it doesn't make sense that 100 i32s would take 
    dramatically more memory than 1000 i128s. My best guess to this is that because we are using a vector, the i128s are
    getting stored on the heap at runtime instead of in the stack and therefore aren't taking any stack memory. This makes
    a lot more sense than my consideration that the stack may pre-allocate memory for its contents, because that could
    be wasteful. i.e. the array stores its contents at compile time on the stack, therefore taking more stack space and 
    limiting the number of stack frames available, while the vector stores its contents at runtime on the heap, taking 
    less stack space and allowing for more stack frames. Therefore for recusive functions it may make more sense to use vectors
    than arrays.

Reflection:
===========
   Getting the recursion algorithm right for part 1 took me longer than it should have probably, but in the end I got it. 
   I think it's cool that we got to show the golden ratio for each lucas sequence. It wasn't too hard to come up with a
   way to avoid killing the program by overflowing the stack, but explaining how that happened in the discussion was more
   difficult. Part 2 was straight forward and helped me learn a lot more about actually coding and implement pretty common
   tasks in rust. Part 3 was intuitive but took longer to work out runtime and how to get the resulting table, and part 4
   was a good practice in stack vs heap memory allocation.


Approximate Hours worked:
=========================
7 hrs

Known Bugs or Limitations:
==========================
N/A?
