# Autocomplete
Autocomplete is a simple experiment in building a autocomplete querying system.

## Process
Before starting the project, I had heard about Tries and wanted to experiment with them. So, this project was created. 

### Iteration 1
The first iteration of the project uses a Trie to store the words and query the stored words that start with a particular prefix.

### Iteration 2/Final Iteration
The second iteration of the program uses a Prefix Hash Tree. After doing some research, I found that both a Trie and Radix Tree
will be too slow for our purposes when the there is enough data stored in the structure. [Prexify](https://prefixy.github.io) has 
explanation of what a prefix hash tree is and why it should be used over a Trie and Radix Tree. Using a prefix hash tree, we 
can fetch a list of words that start with the given prefix in O(1), something that is not possible with a Trie or a Radix Tree. 

In this case, we are using more memory to reduce the time in which a fetch a list of completions for a given prefix, but 
since autocomplete suggestion need to be shown very quickly, it is important to prioritize speed. 

So, the second data structure is a dictionary where the key is the prefix and the value is a set of string that will be the
completion strings.

## Time comparisons for the 2 data structures used

Times were obtained for both insertion and querying both data structures with varying number of words inserted into the data 
structures. On each run, the words stored in both the data structures were the same and the prefixes being queried were also
the same. Both data structure were tested 100 times for the following number of words: 1000, 10000, 100000 and 1000000.

Prefix Hash Tree = PHT , Insertion Time = IT, Querying Time = QT

The following results were obtained:

| # of Words | IT/Word (Trie) | IT/Word (PHT) | QT/Prefix (Trie) | QT/Prefix (PHT) |
|------------|----------------|---------------|------------------|-----------------|
| 1000       | 4.037010e-05   | 5.404002e-06  | 1.653383e-04     | 9.963251e-07    |
| 10000      | 6.574455e-05   | 1.569968e-05  | 1.510254e-03     | 3.737118e-06    |
| 100000     | 7.421327e-05   | 1.872740e-05  | 1.180164e-02     | 3.888364e-05    |
| 1000000    | 5.565568e-05   | 1.375930e-05  | 7.429206e-02     | 8.035530e-04    |


As we can see from the results above, PHT is 2 order of magnitude faster than a trie at querying when 1000000 words are 
stored in the PHT, and when querying less words, PHT is even 3 order of magnitude faster than a Trie. So PHT seems to be 
a better struture for querying results quickly even though the results show an increase in query time per prefix. 

### Thought Process
- ~~Iteration 2: Radix Tree~~ Will not be done due to the fact the runtime for this structure is not fast enough. Replaced by
Prefix Hash Tree.  
- Find other, faster data structures for fetching, even if insertion is slower, has been found, will be using Prefix Hash Tree
because of the explanation at the following link: https://prefixy.github.io

