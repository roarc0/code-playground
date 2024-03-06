#include <gtest/gtest.h>

#include <string>

#include "is_subsequence.hpp"

TEST(IsSubsequence, MultiTest) {
  struct TestInputs {
    std::string s;
    std::string t;
    bool want;
  };
  std::vector<TestInputs> testInputs = {
      {"", "ahbgdc", true},
      {"abc", "ahbgdc", true},
      {"axc", "ahbgdc", false},
  };

  auto s = Solution{};
  for (const auto &t : testInputs) {
    EXPECT_EQ(t.want, s.isSubsequence(t.s, t.t));
  }
}