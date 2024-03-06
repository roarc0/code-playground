#include <gtest/gtest.h>

#include <string>
#include <vector>

#include "ransom_note.hpp"

TEST(RansomNote, MultiTest) {
  struct TestInputs {
    std::string ransomNote;
    std::string magazine;
    bool want;
  };
  std::vector<TestInputs> testInputs = {
      {"a", "b", false},
      {"aa", "ab", false},
      {"aab", "baa", true},

  };

  auto s = Solution{};
  for (const auto &t : testInputs) {
    EXPECT_EQ(t.want, s.canConstruct(t.ransomNote, t.magazine));
  }
}