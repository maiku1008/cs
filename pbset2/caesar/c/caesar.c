#include <cs50.h>
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[])
{
    if (argc != 2)
    {
        printf("Usage: ./caesar key\n");
        return 1;
    }
    string key = argv[1];
    for (int i = 0, n = strlen(key); i < n; i++)
    {
        if (!isdigit(key[i]))
        {
            printf("Usage: ./caesar key\n");
            return 1;
        }
    }
    int k = atoi(key);
    char c;

    string p = get_string("plaintext: ");
    printf("ciphertext: ");

    for (int i = 0, n = strlen(p); i < n; i++)
    {
        c = p[i];
        if (isalpha(c))
        {
            if (islower(c))
            {
                c = (c + k - 'a') % 26 + 'a';
            }
            if (isupper(c))
            {
                c = (c + k - 'A') % 26 + 'A';
            }
        }
        printf("%c", c);
    }
    printf("\n");
    return 0;
}
