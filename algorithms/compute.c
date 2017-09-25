#include <stdio.h>
#include "stack.h"

int main(int argc, char **argv)
{
    char a;
    stack *operand_stack = stack_new();
    while (scanf("%c", &a) != EOF)
    {
        
    }
    stack_free(operand_stack);
    return 0;
}
