#include "find_first_of.hpp"

int Solution::strStr(const std::string &haystack, const std::string &needle) {
  auto pos = haystack.find(needle);
  return pos == std::string::npos ? -1 : pos;
}