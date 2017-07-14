#include "find_dbs.h"
#include <dirent.h>
#include <iostream>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <vector>
#include <string>

//using namespace std;
std::vector<std::string> dirs;
std::vector<std::basic_string<char> > getDirs(char* path) {
    
    DIR *dir_p = NULL;
    dirent *dir_ent = NULL;
    dir_p = opendir(path);
 
    while ((dir_ent = readdir(dir_p)) != NULL) {
        char* absolute_path = get_absolute_path(path, dir_ent->d_name);
        //get_absolute_path(path, dir_ent->d_name);
        if ((absolute_path != NULL) && (dir_ent->d_type ==  DT_DIR)) {
            //printf("%s\n", absolute_path);

            dirs.push_back(absolute_path);
            getDirs(absolute_path);
            free(absolute_path);
        }
        
    }
    closedir(dir_p);
    return dirs;

}

char* get_absolute_path(char* path, char* dname) {
    if (!(strcmp(dname, ".")) || !(strcmp(dname, "..")) ){
        return NULL;
    }

    if ((strcmp(path, "/"))) {
        char *newpath = (char*) malloc(strlen(path)+2+strlen(dname));
        strcpy(newpath, path);
        strcat(newpath, "/");
        strcat(newpath, dname);
        return newpath;
    } else {
        char *newpath = (char*) malloc(strlen(path)+1+strlen(dname));
        strcpy(newpath, path);
        strcat(newpath, dname);
        return newpath;
    }

}