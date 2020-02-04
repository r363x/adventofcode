#include <stdio.h>
#include <sys/stat.h>
#include <stdlib.h>
#include <ctype.h>
#include <stdbool.h>
#include <string.h>

bool clean_up_once(char* data);
void usage();

int main(int argc, char *argv[])
{

    // Handle usage and command line flags
    if (argc != 3)
        usage();
    if (strncmp(argv[1], "-f", 2))
        usage();

    // Read the input data filename
	char* filename = argv[2];

    // Declare variables for later use
	FILE* fd;
	struct stat file_info;
	char* data;

	// Open the input file for reading
	fd = fopen(filename, "r");
	
	// Error checking
	if (fd == NULL)
	{
		printf("Could not open file %s\n", filename);
		return 1;
	}

	// Read the file stats into file_info
	fstat(fileno(fd), &file_info);

	// Allocate storage for file data
	data = malloc(sizeof(char) * file_info.st_size + 1);

	// Auxiliary vars
	char c;
	int i = 0;

	// Read the whole file into "data"
	while ((c = fgetc(fd)) != EOF)
	{
		// Drop the newlines
		if (c == '\n')
			continue;
	
		// Put the character into "data"
		data[i++] = c;
	}

	// Terminate the string
	data[i] = '\0';

	// Aux vars
	bool mod = false;

	// Do as many iterations of cleanup as needed
	while (true)
	{
		if (!clean_up_once(data))
			break;
		mod = true;
	}

	// Shrink the size of "data" if needed
	if (mod)
		data = realloc(data, sizeof(char) * strlen(data) + 1);

	// Finally print the cleaned up data
	printf("Remaining units: %ld\n", strlen(data));

	// Return the allocated memory to the system
	free(data);

	// Close the file
	fclose(fd);

	return 0;
}

bool clean_up_once(char* data)
{
	// Auxiliary vars
	int i, j;
	bool mod = false;

	for(i = 1, j = 0; i < strlen(data)+1; i++, j++)
	{
		// If current and previous chars are the same (nocase)
		if (tolower(data[i]) == tolower(data[i-1]))
		{
			// If cases differ - this is it!
			if (data[i] != data[i-1])
			{
				i += 2;
				mod = true;
			}
		}
		data[j] = data[i-1];
	}
	data[j] = '\0';
	return mod;
}

void usage()
{
    printf("Usage\n"
           "--------\n"
           "-f     input data filename\n\n");
    exit(1);
}
