#include "is_subsequence.hpp"

bool Solution::isSubsequence(const std::string &s, const std::string &t) {
  if (s.empty()) {
    return true;
  }
  uint len = 0;
  for (auto c : t) {
    if (s[len] == c) {
      len++;
      if (len >= s.length()) {
        return true;
      }
    }
  }
  return false;
}
