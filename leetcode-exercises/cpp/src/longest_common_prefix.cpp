#include "longest_common_prefix.hpp"

#include <ranges>

std::string Solution::longestCommonPrefix(std::vector<std::string> &strs) {
  if (strs.empty()) {
    return std::string();
  }

  uint len = 0;
  auto &first = strs[0];
  auto rest = strs | std::views::drop(1);
  while (true) {
    if (len >= first.length()) {
      break;
    }
    for (auto &s : rest) {
      if (len >= s.length() || first[len] != s[len]) {
        goto end; // :3 look how cute!
      }
    }
    len++;
  }
end:
  return strs[0].substr(0, len);
}
