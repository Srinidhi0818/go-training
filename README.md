# go-training

# Data Types In GO
1.	Integers: Signed and Unsigned
2.	Floats
3.	Byte
4.	Rune
5.	String
6.	Boolean
7.	Arrays
8.	Structs
9.	Slices
10.	Maps
11.	Functions
12.	Pointers
13.	Channels

In signed integer there are:
•	Int
•	Int8
•	Int16
•	Int32
•	Int64
In unsigned integer there are:
•	Uint
Uint8
•	Uint16
•	Uint32
•	Uint64
Floats has 2 types:
•	Float 64
•	Float 32

# STRING FORMATTING
Printf - used to print the formatted string
Sprintf - used to create and return the string without printing it

%d - to print int
%f - to print floats
%t - to format bool
%s - to format string
%c - to format characters
%b - to format binary
%x - to format hex value
%T - to get the type
%p - to format pointers
%v - to print value of the struct
%q - to use " "
%e and %E - to format float in scientific notation
\t - to tab

# SLICES
len() function - returns the length of the slice (the number of elements in the slice)
cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
Slices can also be copied

# POINTERS
The default value or zero-value of a pointer is always nil
Declaration and initialization of the pointers can be done into a single line
The address of a value is storedd in a variable by using the ampersand (&) symbol.

# MAPS
To create an empty map, use the builtin make: make(map[key-type]val-type)
Set key/value pairs using typical name[key] = val syntax
Get a value for a key with name[key]
The builtin len returns the number of key/value pairs when called on a map
The builtin delete removes key/value pairs from a map
Maps appear in the form map[k:v k:v] when printed with fmt.Println
Attempting to add keys to a nil map will result in a runtime error
Can initialize a map using the built-in make() function

# STRUCTS

Go’s structs are typed collections of fields. They’re useful for grouping data together to form records
Can access individual fields of a struct using the dot (.) operator
Can get a pointer to a struct using the & operator 
Can also use the built-in new() function to create an instance of a struct. The new() function allocates enough memory to fit all the struct fields, sets each of them to their zero value and returns a pointer to the newly allocated struct 