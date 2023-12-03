"""
Implementation of a trie, AKA prefix tree/digital tree. Includes the implementation of Trie Node and Trie itself.
Supports insertion of any

This implementation uses a traditional tree, not the space optimized radix tree. It only supports insert and search,
does not have an implementation of delete.

More details at:
    https://en.wikipedia.org/wiki/Trie
"""
from src.DataStructures.Node import Node


class Trie:
    """
    Implementation of the trie.

    Stores the root of the trie. Contains implementation for insert and search, no implementation of delete.
    """

    def __init__(self):
        """
        The initializer.
        """
        self.root = Node("-1")

    def insert(self, to_insert):
        """
        Inserts the given to_insert into the trie if it is not already there.
        :param to_insert:
        :return:
        """
        current = self.root

        for c in to_insert:
            node = Node(c)
            if c not in current.children_index:
                current.children.append(node)
                current.children_index[c] = len(current.children) - 1
            current = current.children[current.children_index[c]]

        current.leaf = True

    def query(self, prefix):
        """
        Searches through the trie to find the words that start with this prefix.

        Returns the results, all the words that start with the prefix.
        :param prefix:
        :return:
        """
        if len(prefix) < 1:
            return set()
        current = self.root

        for c in prefix:
            if c in current.children_index:
                current = current.children[current.children_index[c]]
            else:
                return set()

        return self.dfs(prefix, current)

    def dfs(self, prefix, node):
        """
        Traverse the tree from the node passed in to grab all the words that should be suggested.
        :param prefix:
        :param node:
        :return:
        """
        words = []
        stack = [node]
        cvt = {}
        word = prefix

        while len(stack) > 0:
            if len(node.children) == 0:
                if node.leaf:
                    words.append(word)

                stack.pop()
                word = word[:-1]
                if len(stack) > 0:
                    node = stack[-1]
            else:
                if node.unique_id not in cvt:
                    if node.leaf:
                        words.append(word)
                    cvt[node.unique_id] = 0
                    stack.append(node.children[0])
                    node = node.children[0]
                    if node.value != -1:
                        word += node.value
                elif cvt[node.unique_id] < len(node.children) - 1:
                    cvt[node.unique_id] += 1
                    stack.append(node.children[cvt[node.unique_id]])
                    node = node.children[cvt[node.unique_id]]
                    if node.value != -1:
                        word += node.value
                else:
                    stack.pop()
                    word = word[:-1]
                    if len(stack) > 0:
                        node = stack[-1]

        return words
