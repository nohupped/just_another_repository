#include <iostream>
#include "find_dbs.h"
#include "leveldb/db.h"
#include <sys/stat.h>
#include <vector>

//using namespace std;
int main(int argc, char** argv) {
      if (argc < 2) {
        std::cout << "Please give a path as argument" << std::endl;
        return 1;
    }
    struct stat info;
    if( stat( argv[1], &info ) != 0 ){
        std::cout << "cannot access " << argv[1] << std::endl;
        return 1;
    }
    std::vector<std::string> dirs = getDirs(argv[1]);
    for (std::vector<std::string>::const_iterator i = dirs.begin(); i != dirs.end(); ++i){
        std::cout << *i << "\n";
    }
    leveldb::DB* db;
    leveldb::Options options;
    options.create_if_missing = true;
    // todo shit
    leveldb::Status status = leveldb::DB::Open(options, "/tmp/somedir", &db);
    assert(status.ok());
}
