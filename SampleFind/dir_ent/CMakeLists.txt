cmake_minimum_required(VERSION 3.6)
project(dir_ent)

set(CMAKE_CXX_STANDARD 14)

set(SOURCE_FILES main.cpp ../DirEnt/library.cpp)
add_executable(dir_ent ${SOURCE_FILES})