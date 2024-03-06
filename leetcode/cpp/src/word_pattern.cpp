#include "word_pattern.hpp"

#include <ranges>
#include <unordered_map>

bool Solution::wordPattern(const std::string &pattern, const std::string &s) {
  std::unordered_map<char, std::string> m;
  std::unordered_map<std::string, char> m2;

  uint i = 0;
  auto splitted =
      s | std::ranges::views::split(' ') |
      std::ranges::views::transform([](auto &&rng) {
        return std::string(&*rng.begin(), std::ranges::distance(rng));
      });

  for (const auto w : splitted) {
    if (!m.contains(pattern[i])) {
      if (m2.contains(w)) {
        return false;
      }
      m2[w] = pattern[i];
      m[pattern[i]] = w;
    } else if (m[pattern[i]] != w) {
      return false;
    }
    i++;
    if (i > s.length()) {
      return false;
    }
  }
  return i >= pattern.length();
}