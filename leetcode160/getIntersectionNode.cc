#include <stdio.h>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
    public:
        ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
            if (headA == NULL || headB == NULL) {
                return NULL;
            }
            if (headA == headB) {
                return headA;
            }
            ListNode *pA = headA;
            ListNode *pB = headB;
            int Anum = 0;
            int Bnum = 0;

            while(pA->next != NULL) {
                pA = pA->next;
                Anum++;
            }
            while(pB->next != NULL) {
                pB = pB->next;
                Bnum++;
            }
            if (pA != pB) {
                return NULL;
            }
            pA = headA;
            pB = headB;
            if (Anum - Bnum > 0) {
                int i = 0;
                while (i < Anum - Bnum) {
                    pA = pA->next;
                    i++;
                }
            }
            if (Bnum - Anum > 0) {
                int i = 0;
                while (i < Bnum - Anum) {
                    pB = pB->next;
                    i++;
                }
            }

            while (pA->next != NULL && pB->next != NULL) {
                if (pA == pB) {
                    return pA;
                } else {
                    pA = pA->next;
                    pB = pB->next;
                }
            }
            if (pA == pB) {
                return pA;
            }
            return NULL;
        }
};
