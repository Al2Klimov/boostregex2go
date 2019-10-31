cmake_minimum_required(VERSION 2.8.8)

project(boostregex CXX)

set(CMAKE_CXX_STANDARD 11)

set(BOOST_MIN_VERSION "1.53.0")
find_package(Boost ${BOOST_MIN_VERSION} COMPONENTS regex system REQUIRED)
include_directories(${Boost_INCLUDE_DIRS})

add_library(boostregex regex.cpp)

file(WRITE "${CMAKE_SOURCE_DIR}/../ld.go" "

package boostregex2go

/*
#cgo LDFLAGS: -L${CMAKE_BINARY_DIR} -L${Boost_LIBRARY_DIRS} -lboostregex -lstdc++ -lboost_regex -lboost_system
*/
import \"C\"

")
