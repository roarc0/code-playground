#include <gtest/gtest.h>

#include <vector>

#include "longest_common_prefix.hpp"

TEST(LongestCommonPrefix, MultiTest) {
  struct TestInputs {
    std::vector<std::string> strVec;
    std::string want;
  };
  std::vector<TestInputs> testInputs = {
      {{"flower", "flow", "flight"}, "fl"},
  };

  auto s = Solution{};
  for (auto &t : testInputs) {
    EXPECT_EQ(t.want, s.longestCommonPrefix(t.strVec));
  }
}