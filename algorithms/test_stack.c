#include <stdio.h>
#include "stack.h"

int main(int argc, char **archv)
{
    stack *s = stack_new();
    int array_length = 5;
    int array[] = {5, 10, -5, 3, 8};
    for (int i = 0; i < array_length; i++)
    {
        stack_push(s, array[i]);
    }

    while (stack_get_length(s) > 0)
    {
        int v = stack_pop(s);
        printf("%d\n", v);
    }

    stack_free(s);
    return 0;
}