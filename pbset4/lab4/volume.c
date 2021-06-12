// Modifies the volume of an audio file

#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

// Number of bytes in .wav header
const int HEADER_SIZE = 44;

int main(int argc, char *argv[])
{
    // Check command-line arguments
    if (argc != 4)
    {
        printf("Usage: ./volume input.wav output.wav factor\n");
        return 1;
    }

    // Open files and determine scaling factor
    FILE *input = fopen(argv[1], "r");
    if (input == NULL)
    {
        printf("Could not open file.\n");
        return 1;
    }

    FILE *output = fopen(argv[2], "w");
    if (output == NULL)
    {
        printf("Could not open file.\n");
        return 1;
    }

    float factor = atof(argv[3]);

    // Create array to hold header
    uint8_t header[HEADER_SIZE];

    // Read the data from the input's header to our header
    fread(header, HEADER_SIZE, 1, input);
    fwrite(header, HEADER_SIZE, 1, output);

    // buffer to hold 2 bytes containing each audio sample
    int16_t buffer;

    // While we read each audio sample, fread evaluated as true
    // when we have an error in reading, or reach the end of file, the loop stops
    while (fread(&buffer, sizeof(int16_t), 1, input))
    {
        // modify the buffer
        buffer *= factor;
        // write the modified buffer
        fwrite(&buffer, sizeof(int16_t), 1, output);
    }

    // Close files
    fclose(input);
    fclose(output);
}
