#include "ransom_note.hpp"

#include <algorithm>
#include <string>
#include <vector>

bool Solution::canConstruct(const std::string &ransomNote,
                            const std::string &magazine) {
  auto used = std::vector<bool>(magazine.size(), false);
  ulong start = 0;

  return std::ranges::all_of(
      ransomNote.begin(), ransomNote.end(), [&](auto c) -> bool {
        for (auto i = start; i < magazine.length(); i++) {
          if (!used[i] && magazine[i] == c) {
            used[i] = true;
            if (start < i && std::all_of(used.begin() + start, used.begin() + i,
                                         [](auto v) { return v; })) {
              start = i;
            }
            return true;
          }
        }
        return false;
      });
}