#include <stdio.h>
#include <stdlib.h>
#include "stack.h"
/* stack implementation */

/* allocate a new node */
node *node_new()
{
    node *n = (node *)malloc(sizeof(node));
    return n;
}

void node_set_key(node *n, int key)
{
    n->key = key;
}

void node_free(node *n)
{
    free(n);
}

stack *stack_new()
{
    stack *s = (stack *)malloc(sizeof(stack));
    s->head = node_new();
    s->tail = node_new();
    s->head->next = s->tail;
    s->length = 0;
    return s;
};
/* push a value to the stack */
void stack_push(stack *s, int key)
{
    node *n = node_new();
    node_set_key(n, key);
    n->next = s->head->next;
    s->head->next = n;
    s->length = s->length + 1;
}
/* return the langth of the stack */
int stack_get_length(stack *s)
{
    return s->length;
}
/* pop the first value of a stack */
int stack_pop(stack *s)
{
    if (s->length > 0)
    {
        node *popped_node = s->head->next;
        s->head->next = popped_node->next;
        s->length = s->length - 1;
        int result = popped_node->key;
        node_free(popped_node);
        return result;
    }
    return 0;
}
/** free a stack */
void stack_free(stack *s)
{
    while (stack_get_length(s) > 0)
    {
        stack_pop(s);
    }
    node_free(s->tail);
    node_free(s->head);
    free(s);
}

