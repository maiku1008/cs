#include <cs50.h>
#include <stdio.h>

void printHashes(int hashes);
void printSpaces(int hashes);

int main(void)
{
    // Prompt the user for the desired height
    int height = 0;
    do
    {
        height = get_int("Height: ");
    }
    while (height < 1 || height > 8);

    // Draw the pyramids
    for (int i = 1; i <= height; i++)
    {
        int spaces = height - i;
        int hashes = height - spaces;

        printSpaces(spaces);
        printHashes(hashes);
        // printSpaces(2);
        // printHashes(hashes);
        printf("\n");
    }
}

// Print the hashes
void printHashes(int hashes)
{
    for (int j = 0; j < hashes; j++)
    {
        printf("#");
    }
}

// Print the spaces
void printSpaces(int spaces)
{
    for (int j = 0; j < spaces; j++)
    {
        printf(" ");
    }
}
