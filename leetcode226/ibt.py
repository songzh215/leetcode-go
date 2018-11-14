# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

import Queue

class Solution(object):
    def invertTree(self, root):
        q = Queue.Queue()
        if root == None:
            return root
        root.left, root.right = root.right, root.left
        q.put(root.left)
        q.put(root.right)
        while not q.empty():
            n = q.get()
            if n == None:
                continue
            n.left, n.right = n.right, n.left
            q.put(n.left)
            q.put(n.right)
        return root
