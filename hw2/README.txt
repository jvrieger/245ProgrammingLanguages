NAME:
=====
	Julia Rieger

Programs Files:
===============
    Part 1:
        structs.go
    Part 2:
        readfiles.go
    Part 3:
        slices.go
	
How to Compile and Run:
===============
    Part 1:
       go run structs.go
    Part 2:
       go run readfiles.go
    Part 3:
       go run slices.go
       
Reflection:
===========
    Populating the launches slice took the longest because I had to assign every
    attribute of the launch instance read from the file, also it took me a long 
    time to fully figure out and understand how to convert strings to numbers,
    but I'm glad I did it because now I really get it and it probably won't take
    me as long next time. For a long time I tried to figure out how to format a
    launch where it will not have a double space if there is no category (ie the
    value of category is ""), since the space before and after the category will
    still print and it will appear double.

Questions:
==========
    1. The original and copy were not equal after changing a value in the copy to
        be different than the original. This is because Go compares each field of
        the struct seperately. Java on the other hand compares references instead
        of the objects themselves if you use the '==' operator. If you want to 
        actually compare the value of java objetcs you have to use the equals() or
        compareTo() methods, which the coder must usually implement themselves. When
        manually implementing these methods the coder can specify what attribute of
        the object they are comparing. To answer the question, if the launch type
        (class in java) had a compareTo(), it could compare the values of each attribute
        in which case changing one of them would also result in the original and the
        copy being unequal. However if the compareTo() only compared specific attributes,
        and the one being changed in the copy was not one of those, the original and 
        the copy would be equal. 

    2. Yes, item 1 in the original slice (launches) and the slice with the first 5
        items (first5) are equal. After changing the value of a field in the struct
        in position 1 of launches, the value of the same field in the struct in
        position 1 of first5 also changed. This means they still evaluated to equal
        even after changing the value of one of the structs. This initially confused 
        me because I know Go uses the Pass By Value semantic, where arguments are 
        passed by value so functions receive a copy of each argument and don't just
        reference the variable thats being passed. A possible explanation to this could
        be that Go doesn't do this for copying values to variables, ie when assigning
        values to slices Go assigns references instead of an actual copy of the value. 
        It could also be that since in my implementation of first5 (making a subslice 
        of the original slice, launches) making a subslice actually just copies the
        reference pointers in the original slice instead of creating whole new instances.
        This would make sense because it would save memory since the computer wouldn't
        have to create as many new variables/structs.

I Worked With:
==============
N/A

Approximate Hours worked:
=========================
3 hrs

Known Bugs or Limitations:
==========================
    None that I know of