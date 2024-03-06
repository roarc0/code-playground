#include "is_isomorphic.hpp"

#include <unordered_map>

template <typename Container>
bool valueExistsInContainer(const Container &c,
                            const typename Container::mapped_type &value) {
  for (const auto &pair : c) {
    if (pair.second == value) {
      return true;
    }
  }
  return false;
}

bool Solution::isIsomorphic(std::string s, std::string t) {
  if (s.length() != t.length()) {
    return false;
  }
  std::unordered_map<char, char> m;
  for (uint i = 0; i < s.length(); i++) {
    if (!m.contains(s[i])) {
      if (valueExistsInContainer(m, t[i])) {
        return false;
      }
      m[s[i]] = t[i];
    } else if (m[s[i]] != t[i]) {
      return false;
    }
    if (t[i] != s[i]) {
      t[i] = s[i];
    }
  }
  return s == t;
}
