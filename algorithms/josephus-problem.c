#include <stdio.h>
#include <stdlib.h>
#define N 9
#define M 5
struct node
{
    int key;
    struct node *next;
};

int main()
{
    int i = 0;
    struct node *t, *x;
    t = (struct node *)malloc(sizeof *t);
    t->key = 1;

    // keep a reference to the first element
    x = t;
    // build the list
    for (i = 2; i <= N; i++)
    {
        t->next = (struct node *)malloc(sizeof *t);
        t = t->next;
        t->key = i;
    }
    t->next = x;
    while (t != t->next)
    {
        for (i = 1; i < M; i++)
        {
            t = t->next;
        }
        printf("%d ", t->next->key);
        x = t->next;
        t->next = t->next->next;
        free(x);
    }
    printf("%d\n", t->key);
    return 0;
}
