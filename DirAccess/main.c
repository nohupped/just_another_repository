#include <stdio.h>
#include <dirent.h>
int
main (void)
{
    DIR* dir_p;
    struct dirent* dir_ent;

    dir_p = opendir ("/");

    if (dir_p != NULL) {
        // The readdir() function returns a pointer to a dirent structure representing the next
        // directory entry in the directory stream pointed to by dirp.
        // It returns NULL on reaching the end of the directory stream or if an error occurred.
        while ((dir_ent = readdir (dir_p)) != NULL) {
            // printf("%s", dir_ent->d_name);
            // printf("%d", dir_ent->d_type);
            if (dir_ent->d_type == DT_DIR) {
                printf("%s is a directory", dir_ent->d_name);
            } else {
                printf("%s is not a directory", dir_ent->d_name);
            }

            printf("\n");
        }
            (void) closedir(dir_p);

    }
    else
        perror ("Couldn't open the directory");

    return 0;
}
