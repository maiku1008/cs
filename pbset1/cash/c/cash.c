#include <cs50.h>
#include <math.h>
#include <stdio.h>

int main(void)
{
    float input;
    do
    {
        input = get_float("Change owed: ");
    } while (input < 0);

    int change = round(100 * input);
    int coins = 0;

    do
    {
        if (change >= 25)
        {
            change -= 25;
        }
        else if (change >= 10)
        {
            change -= 10;
        }
        else if (change >= 5)
        {
            change -= 5;
        }
        else if (change >= 1)
        {
            change -= 1;
        }
        coins++;
    } while (change > 0);
    printf("%i\n", coins);
}
