#include "length_of_last_word.hpp"

int Solution::lengthOfLastWord(const std::string &s) {
  if (s.length() == 0) {
    return 0;
  }
  auto sv = std::string_view(s);

  auto end = sv.find_last_not_of(" ");
  if (end == std::string_view::npos) {
    return 0;
  } else if (end != s.length()) {
    sv = sv.substr(0, end + 1);
  }

  auto start = sv.find_last_of(" ");
  return end + 1 - (start != std::string_view::npos ? (start + 1) : 0);
}
