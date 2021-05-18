#include <cs50.h>
#include <ctype.h>
#include <stdio.h>
#include <string.h>
#include <math.h>

int main(void)
{
    string text = get_string("Text: ");

    float letters = 0;
    float words = 1;
    float sentences = 0;
    for (int i = 0, n = strlen(text); i < n; i ++)
    {
        char c = tolower(text[i]);
        if ('a' <= c && c <= 'z')
        {
            letters++;
        }
        if (c == ' ')
        {
            words++;
        }
        if (c == '.' || c == '!' || c == '?')
        {
            sentences++;
        }
    }
    float average_letters = (letters / words) * 100;
    float average_sentences = (sentences / words) * 100;
    int index = round(0.0588 * average_letters - 0.296 * average_sentences - 15.8);

    if (index < 1)
    {
        printf("Before Grade 1\n");
        return 0;
    }
    if (index > 16)
    {
        printf("Grade 16+\n");
        return 0;
    }
    printf("Grade %i\n", index);
    return 0;
}

