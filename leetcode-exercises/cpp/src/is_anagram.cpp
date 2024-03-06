#include "is_anagram.hpp"

#include <algorithm>
#include <unordered_map>

bool Solution::isAnagram(const std::string &s, const std::string &t) {
  if (s.length() != t.length()) {
    return false;
  }
  auto m = std::unordered_map<char, int>();
  for (ulong i = 0; i < s.length(); i++) {
    m[s[i]] += 1;
    m[t[i]] -= 1;
  }
  return std::ranges::all_of(m.begin(), m.end(),
                             [](const auto &e) { return e.second == 0; });
};
