"""
Implementation of Node.
"""
import uuid


class Node:
    """
    Implementation of a node in a trie/tree.

    Stores all the children of the node, the words that start with this word, if this is a leaf node and value of the
    character in the current node.

    Since we do no limit the characters that can be stored in each node(anything in unicode is allowed), we will be
    using a set to store the children of each node. Using an array is not practical in this case, due to the fact that
    0x10FFFF is too large of an array at every node, especially since the connection between every unique code character
    are pretty unlikely.

    Each node stores its children as a set, meaning the look ups on average will take O(1) to find a child in a node.

    Furthermore, the children are stored in an ordered set so we can iterate through them.
    """

    def __init__(self, value):
        """
        The initializer, stores the value of the node passed in and initializes other members.
        """
        self.value = value
        self.children = []
        self.children_index = dict()
        self.leaf = False
        self.unique_id = self.__unique_identifier()

    def add_child(self, node):
        """
        Takes a node as a child, check if it is already a child, if not, adds it, if yes, throws an error
        :param node:
        :return:
        """
        if node.value not in self.children_index:
            self.children.append(node)
            self.children_index[node.value] = len(self.children) - 1
        else:
            raise ValueError(
                "Node::add_child attempting to insert duplicate child: '{0}'".format(
                    node.value
                )
            )

    def __unique_identifier(self):
        """
        Returns a uuid that can be used to uniquely identify a node outside of its value
        :return:
        """
        return uuid.uuid1()

    def __eq__(self, other):
        """
        Compares the value and returns true/false. Children are not compared due to the fact it would slow down the
        comparison of the node too much and it there is no real need for it to make the object hashable.
        :param other:
        :return:
        """
        return self.value == other.value

    def __hash__(self):
        """
        Returns the value since it is unique for our purposes.
        :return:
        """
        return ord(self.value)
