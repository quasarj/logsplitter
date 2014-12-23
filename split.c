#include <stdio.h>
#include <string.h>

#define BUFFSIZE 4096

// replace any / with _
void fix_filename(char* filename) {
    char* pos;

    while (pos = strchr(filename, '/')) {
        *pos = '_';
    }
}

int main(void) {
    printf("Logsplitter in c, v0.2\n");

    FILE* input_file;
    FILE* output_file = NULL;
    char buf[BUFFSIZE];
    int pos;
    char date[10];
    char filename[50];

    input_file = fopen("WoWCombatLog.txt", "r");
    if (!input_file) {
        printf("Error opening input file :(\n");
        return 1;
    }

    // read lines from the file
    while (fgets(buf, BUFFSIZE, input_file) != NULL) {
        pos = strchr(buf, ' ') - buf;

        if (strncmp(date, buf, pos)) {
            if (output_file)
                fclose(output_file);

            strncpy(date, buf, pos);
            printf("New date: %s\n", date);

            // generate the new filename
            sprintf(filename, "wow_log__%s.txt", date);
            fix_filename(filename);
            
            output_file = fopen(filename, "w");
        }

        fputs(buf, output_file);
    }
    
    fclose(output_file);
    fclose(input_file);

    return 0;
}
