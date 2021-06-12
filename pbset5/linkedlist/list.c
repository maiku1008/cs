#include <stdio.h>
#include <stdlib.h>

typedef struct node
{
    int number;
    struct node *next;
}
node;

int main(void)
{
    // Initial list variable which is set to NULL so we avoid garbage memory
    node *list = NULL;

    // n is a temporary variable, and we allocate it to the size of one node
    node *n = malloc(sizeof(node));
    if (n == NULL)
    {
        return 1;
    }
    // Set the values for first node
    (*n).number = 1;
    (*n).next = NULL;
    // Set the root node
    list = n;

    // Use n again, and allocate one space the size of a node
    n = malloc(sizeof(node));
    if (n == NULL)
    {
        // if we cannot for any reason, get rid of the allocation and return
        free(list);
        return 1;
    }
    // set the values
    n->number = 2;
    n->next = NULL;
    // point list's next to n
    list->next = n;

    // ditto like before
    n = malloc(sizeof(node));
    if (n == NULL)
    {
        printf("returning at third malloc\n");
        free(list->next);
        free(list);
        return 1;
    }
    n->number = 3;
    n->next = NULL;
    list->next->next = n;

    // Traverse the linked list, starting from the root node, and print the values for each node
    for (node *tmp = list; tmp != NULL; tmp = tmp->next)
    {
        printf("%i\n", tmp->number);
    }

    // Free the list cause we're done
    while (list != NULL)
    {
        node *tmp = list->next;
        free(list);
        list = tmp;
    }
}
