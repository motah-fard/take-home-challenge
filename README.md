Finding the shallowest duplicate value in the Tree!

We have to traverse all the nodes, imagine if the lowest node in the rightest place in the bottom is equal to the root, what will happen? the shallowest duplicate node will be the root! 

So the time complexity can not be better than O(n).

I used BFS here to solve the problem. 
I used a queue and a depth amp data structure one to traverse level by level and ont to use when I see a duplicate, and a pointer to keep the information about the shallowest duplicate node, and when we see a duplicate with a lesser depth we are going to update it. 

From the information above, the space complexity will be O(1).


In the question it is mentioned that we are dealing with a tree, not a binary tree and it is not mentioned that what type of ID (or value) each node carries.

So I made an interface for the type to be able to get whatever information that we want and assume that the nodes can have multiple children.

If you have any questions I will be glad to answer them, thanks thanks :D Have a great week!
