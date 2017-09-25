typedef struct node
{
    int key;
    struct node *next;
} node;

node *node_new();

void node_set_key(node *n, int key);

int node_get_key(node *n);

void node_free(node *n);

typedef struct
{
    node *head;
    node *tail;
    int length;
} stack;

stack *stack_new();

void stack_push(stack *s, int key);

int stack_get_length(stack *s);

int stack_pop(stack *s);

void stack_free(stack *s);