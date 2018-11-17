#include <stdio.h>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {

private:
    TreeNode* res;

public:
    int find(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == NULL) {
            return 0;
        }
        int count = 0;
        if (root->val == p->val || root->val == q->val) {
            count++;
        }
        if (find(root->left,p,q) >= 1) {
            count++;
        }
        if (find(root->right,p,q) >= 1) {
            count++;
        }
        if (count == 2) {
            res = root;
        }
        return count;
    }

    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == NULL) {
            return root;
        }
        if (find(root, p, q) == 0) {
            return res;
        }
        lowestCommonAncestor(root->left, p, q);
        lowestCommonAncestor(root->right, p, q);
        return res;
    }
};

int main(int argc, char** args) {
    TreeNode n1(3);     
    TreeNode n2(5);     
    TreeNode n3(1);     
    TreeNode n4(6);     
    TreeNode n5(2);     
    TreeNode n6(0);     
    TreeNode n7(8);     
    TreeNode n8(7);     
    TreeNode n9(4);
    n1.left = &n2;
    n1.right = &n3;
    n2.left = &n4;
    n2.right = &n5;
    n3.left = &n6;
    n3.right = &n7;
    n5.left = &n8;
    n5.right = &n9;
    Solution s;
    TreeNode * res = s.lowestCommonAncestor(&n1, &n2, &n3);
    printf("case 1 res is %d \n", res->val);
    TreeNode * res1 = s.lowestCommonAncestor(&n1, &n2, &n9);
    printf("case 2 res is %d \n", res1->val);
    return 0;
}
