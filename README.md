# actTree

actTree is a Golang package which implements a double-linked tree data structure.

Nodes in actTree are double-linked together. Therefore it is 
efficient to travel between immediate nodes back and forth, 
making it a good fit for taking on "timeline review" kind of 
features. 

actTree takes a small footprint both on code and at runtime. 
The code was initially a part of a board game for holding the 
player moves and reference diagrams, and then extracted and 
used by many projects for it is almost the least code to have 
when people need a tree structure. 
