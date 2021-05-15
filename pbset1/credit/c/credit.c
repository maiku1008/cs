#include <cs50.h>
#include <stdio.h>

int addDigits(int);
bool isMastercard(int, int, int);
bool isVisa(int, int, int);
bool isAmex(int, int, int);

int main(void)
{
    // Get the number from the user
    long number = get_long("Number: ");

    int sum = 0;
    int length = 0;
    int first = 0;
    int second;

    do
    {
        length++;

        int digit = number % 10;
        second = first;
        first = digit;
        if (length % 2 == 0)
        {
            digit *= 2;
            digit = addDigits(digit);
        }
        sum += digit;
    }
    while (number /= 10);

    string result = "INVALID";
    if (sum % 10 == 0)
    {
        if (isAmex(first, second, length))
        {
            result = "AMEX";
        }
        if (isVisa(first, second, length))
        {
            result = "VISA";
        }
        if (isMastercard(first, second, length))
        {
            result = "MASTERCARD";
        }
    }
    printf("%s\n", result);
}

// Add the digits of an int
// e.g. 12 = 1 + 2 = 3
int addDigits(int number)
{
    int sum = 0;
    do
    {
        sum += number % 10;
    }
    while (number /= 10);
    return sum;
}

// Determine whether the card is a mastercard
bool isMastercard(int first, int second, int length)
{
    return first == 5 && (second == 1 || second == 2 || second == 3 || second == 4 || second == 5) && length == 16;
}

// Determine whether the card is a visa
bool isVisa(int first, int second, int length)
{
    return first == 4 && (length == 13 || length == 16);
}

// Determine whether the card is an american express
bool isAmex(int first, int second, int length)
{
    return first == 3 && (second == 4 || second == 7) && length == 15;
}
