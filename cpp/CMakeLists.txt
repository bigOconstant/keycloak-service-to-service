# CMakeLists.txt
cmake_minimum_required(VERSION 3.0)
project(test)

set(OPENSSL_USE_STATIC_LIBS TRUE)
find_package(OpenSSL REQUIRED)

include(/vcpkg/scripts/buildsystems/vcpkg.cmake)



find_package(nlohmann_json REQUIRED)

include(FetchContent)
FetchContent_Declare(cpr GIT_REPOSITORY https://github.com/whoshuu/cpr.git GIT_TAG c8d33915dbd88ad6c92b258869b03aba06587ff9) # the commit hash for 1.5.0
FetchContent_MakeAvailable(cpr)

add_executable(main main.cpp )

target_include_directories(main PRIVATE )


target_link_libraries(main PRIVATE OpenSSL::Crypto
 nlohmann_json::nlohmann_json  cpr::cpr)