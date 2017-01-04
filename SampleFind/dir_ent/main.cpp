#include <iostream>
#include "../DirEnt/library.h"



int main(int argc, char** argv) {
    if (argc < 2) {
        std::cout << "Please give a path as argument" << std::endl;
        return 1;
    }
    getDirs(argv[1]);
    return 0;
}