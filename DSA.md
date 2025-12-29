## Phase 1: Core Syntax & Linear Structures (Week 1)

### Dec 29 (Mon): The Runtime Environment
- **Study**
  - Go: `main` function structure, `package main`, `fmt.Println`
  - JavaScript: `console.log`, `let` vs `const`
- **Syntax Drill**
  - Write "Hello World"
  - Write a function that adds two integers
- **Problem**
  - FizzBuzz (focus on `if/else` and modulo `%`)
- **Milestone**
  - Compile and run a Go program locally
  - Run a JS script in Node.js

### Dec 30 (Tue): Loops and Conditionals
- **Study**
  - Go: `for` (the only loop)
  - JS: `for`, `for...of`, `while`
- **Syntax Drill**
  - Print numbers 1 to 10
  - Print even numbers from an array
- **Problem**
  - Sum of Array Elements
- **Milestone**
  - Understand why Go uses `_` (blank identifier) in `range` loops

### Dec 31 (Wed): Arrays vs Slices
- **Study**
  - Go: `make([]int, len, cap)`, `append`
  - JS: Arrays, `.push()`, `.pop()`
- **Syntax Drill**
  - Create an array of size 5
  - Append 3 elements
  - Print the length
- **Problem**
  - Find Max Element in Array
- **Milestone**
  - Explain the difference between `len` and `cap` in Go

### Jan 1 (Thu): Strings and Runes
- **Study**
  - Go: `string` vs `byte` vs `rune`
  - JS: Template literals, `.split()`
- **Syntax Drill**
  - Iterate over a string with emojis (e.g., `HeLlO 🌍`)
- **Problem**
  - Valid Palindrome (simple version)
- **Milestone**
  - Successfully handle a multi-byte character in Go

### Jan 2 (Fri): Functions and Multiple Returns
- **Study**
  - Go: Multiple return values `func foo() (int, error)`
  - JS: Destructuring (`const [val, err] = foo()`)
- **Syntax Drill**
  - Function returning integer division result and remainder
- **Problem**
  - Convert Temperature (Celsius → Kelvin/Fahrenheit)
- **Milestone**
  - Write a Go function that returns an error on invalid input

### Jan 3 (Sat): Pointers vs References
- **Study**
  - Go: Pointers (`*int`, `&x`)
  - JS: Reference behavior for objects/arrays
- **Syntax Drill**
  - Swap function
    - Go: Use pointers
    - JS: Swap array elements
- **Problem**
  - Reverse String (in-place)
- **Milestone**
  - Visualize stack vs heap memory

### Jan 4 (Sun): Review & Practice
- **Tasks**
  - Re-solve Reverse String and FizzBuzz
- **Problem**
  - Move Zeroes (two-pointer approach)
- **Milestone**
  - O(N) time and O(1) space solution

---

## Phase 2: Hashing & Frequency Counting (Week 2)

### Jan 5 (Mon): Maps and Objects
- **Study**
  - Go: `map[key]value`
  - JS: `Map` vs `Object`
- **Syntax Drill**
  - Character frequency counter
- **Problem**
  - Contains Duplicate
- **Milestone**
  - Use `struct{}` as a Set in Go

### Jan 6 (Tue): The Anagram Pattern
- **Study**
  - Fixed-size array as map (26 lowercase letters)
- **Problem**
  - Valid Anagram
- **Milestone**
  - Implement using both hashmap and fixed array

### Jan 7 (Wed): Two Sum
- **Study**
  - Complement strategy (`target - current`)
- **Problem**
  - Two Sum
- **Milestone**
  - O(N) solution using a map

### Jan 8 (Thu): Grouping Data
- **Study**
  - Map keys and hashing
- **Problem**
  - Group Anagrams
- **Milestone**
  - Handle array-as-key issue in JS using `.join('#')`

### Jan 9 (Fri): Sets and Uniqueness
- **Study**
  - JS `Set`, Go set via map
- **Problem**
  - Intersection of Two Arrays
- **Milestone**
  - Helper function to extract map keys in Go

### Jan 10 (Sat): Sliding Window (Fixed)
- **Study**
  - Sliding window concept
- **Problem**
  - Find All Anagrams in a String
- **Milestone**
  - Slide window without full recomputation

### Jan 11 (Sun): Review & Assessment
- **Tasks**
  - Re-solve Two Sum and Contains Duplicate
- **Deep Dive**
  - Go map resizing and random iteration order

---

## Phase 3: Pointers, Recursion & Lists (Week 3)

### Jan 12 (Mon): Linked List Basics
- **Study**
  - Node structures
- **Problem**
  - Reverse Linked List
- **Milestone**
  - Draw pointer changes before coding

### Jan 13 (Tue): Merging Lists
- **Study**
  - Dummy node technique
- **Problem**
  - Merge Two Sorted Lists
- **Milestone**
  - Use dummy head to simplify logic

### Jan 14 (Wed): Fast & Slow Pointers
- **Study**
  - Floyd’s Cycle Detection
- **Problem**
  - Linked List Cycle
- **Milestone**
  - Explain why fast catches slow

### Jan 15 (Thu): Stack & Binary Search
- **Study**
  - Stack using arrays/slices
- **Problem**
  - Valid Parentheses
- **Milestone**
  - Handle (), {}, and []
- **Problem**
  - Binary Search
- **Milestone**
  - Correct "not found" handling

### Jan 17 (Sat): Binary Search (Applied)
- **Problem**
  - Search Insert Position / First Bad Version
- **Milestone**
  - Distinguish `lo < hi` vs `lo <= hi`

### Jan 18 (Sun): Recursion 101
- **Study**
  - Base vs recursive case
- **Problem**
  - Fibonacci (recursive vs iterative)
- **Milestone**
  - Visualize recursion tree for `Fib(5)`

---

## Phase 4: Trees & Navigation (Week 4)

### Jan 19 (Mon): Binary Trees & DFS
- **Study**
  - Pre/In/Post-order traversals
- **Problem**
  - Invert Binary Tree
- **Milestone**
  - Recursive traversal of all nodes

### Jan 20 (Tue): Tree Depth
- **Problem**
  - Maximum Depth of Binary Tree
- **Milestone**
  - Understand value bubbling from leaves

### Jan 21 (Wed): Binary Search Trees
- **Study**
  - BST property
- **Problem**
  - Lowest Common Ancestor of BST
- **Milestone**
  - Iterative and recursive solutions

### Jan 22 (Thu): BFS (Level Order)
- **Study**
  - Queues and performance
- **Problem**
  - Level Order Traversal
- **Milestone**
  - Queue implementation in Go

### Jan 23 (Fri): Heaps
- **Study**
  - Priority queues
- **Problem**
  - Kth Largest Element in a Stream
- **Milestone**
  - Understand heap behavior

### Jan 24 (Sat): Tries
- **Study**
  - Prefix trees
- **Problem**
  - Implement Trie
- **Milestone**
  - Why Tries are fast for prefixes

### Jan 25 (Sun): Review
- **Tasks**
  - Re-solve Invert Binary Tree and Valid Parentheses
- **Deep Dive**
  - Go `sort.Search` source code

---

## Phase 5: Advanced & Final Push (Week 5)

### Jan 26 (Mon): Graphs
- **Study**
  - Adjacency lists
- **Problem**
  - Number of Islands
- **Milestone**
  - Correct visited-state handling

### Jan 27 (Tue): Dynamic Programming (Memoization)
- **Problem**
  - Climbing Stairs
- **Milestone**
  - Reduce O(2^N) to O(N)

### Jan 28 (Wed): DP (Tabulation)
- **Problem**
  - Coin Change / House Robber
- **Milestone**
  - Identify sub-problems

### Jan 29 (Thu): Intervals
- **Problem**
  - Merge Intervals
- **Milestone**
  - Custom sort in Go and JS

### Jan 30 (Fri): Mock Interview
- **Task**
  - 1 Easy + 1 Medium (45 mins)
- **Milestone**
  - No running code until end

### Jan 31 (Sat): Hard Problem Exposure
- **Task**
  - Merge K Sorted Lists / Trapping Rain Water
- **Goal**
  - Understand solution approach

### Feb 1 (Sun): Final Review
- **Task**
  - Create syntax cheat sheet (Go + JS)
- **Milestone**
  - Ready for intermediate-level grinding
