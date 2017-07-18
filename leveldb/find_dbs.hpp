#ifndef LEVELDB_FIND_DBS_HPP
#define LEVELDB_FIND_DBS_HPP
#include <vector>
#include <string>
std::vector<std::basic_string<char> > getDirs(char* path);
char* get_absolute_path(char* path, char* dname);

#endif //LEVELDB_FIND_DBS_H
