#include <cs50.h>
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

bool isValidKey(string k);

int main(int argc, char *argv[])
{
    if (argc != 2)
    {
        printf("Usage: ./substitution key\n");
        return 1;
    }
    string key = argv[1];
    if (!isValidKey(key))
    {
        return 1;
    }

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
                c = tolower(key[c - 'a']);
            }
            if (isupper(c))
            {
                c = toupper(key[c - 'A']);
            }
        }
        printf("%c", c);
    }
    printf("\n");
    return 0;
}

// Validate the key
bool isValidKey(string k)
{
    // Ensure the key is exactly 26 characters long
    int n = strlen(k);
    if (n != 26)
    {
        printf("Key must contain 26 characters.");
        return false;
    }

    char c;
    for (int i = 0; i < n; i++)
    {
        c = k[i];
        // Ensure we have an alphabetic character
        if (!isalpha(c))
        {
            printf("Key must contain 26 alphabetical characters.");
            return false;
        }
        // Ensure there are no duplicates by iterating through the key
        for (int j = 0; j < n; j++)
        {
            if (i != j && c == k[j])
            {
                printf("Key must not contain any duplicates.");
                return false;
            }
        }
    }
    return true;
}
