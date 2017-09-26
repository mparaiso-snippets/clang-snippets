#include <stdio.h>
#include <stdlib.h>

/**
 * plus grand commun diviseur
 */
int gcd(int u, int v)
{
    if (u <= 0)
    {
        return v;
    }
    if (u < v)
    {
        return gcd(v - u, u);
    }
    return gcd(u - v, v);
}
void main()
{
    int x, y;
    while (scanf("%d %d", &x, &y) != EOF)
        if (x > 0 && y > 0)
        {
            printf("%d %d %d\n", x, y, gcd(x, y));
        }
    exit(EXIT_SUCCESS);
}