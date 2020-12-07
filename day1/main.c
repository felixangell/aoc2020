#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

int main() {
	FILE* file = fopen("input", "r");
	fseek(file, 0, SEEK_END);
	uint64_t len = ftell(file);
	fseek(file, 0, SEEK_SET);

	char* buf = malloc(sizeof(*buf) * len);
	if (buf) {
		fread(buf, 1, len, file);
	}

	fclose(file);

	int values[256];

	int num_values = 0;
	for (char* line=strtok(buf, "\n");line;line=strtok(NULL, "\n")) {
		values[num_values++] = atoi(line);		
	}

	int result = -1;

	for (int i = 0; i < num_values; i++) {
		for (int j = 0; j < num_values; j++) {
			for (int k = 0; k < num_values; k++) {
				int x = values[i]; int y = values[j]; int z = values[k];
				if (x + y + z == 2020) {
					result = x * y * z;
					break;
				}
			}
		}
	}

	printf("total is %d\n", result);
	return 0;
}
