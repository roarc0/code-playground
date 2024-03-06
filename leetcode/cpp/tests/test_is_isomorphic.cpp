#include <gtest/gtest.h>

#include <string>

#include "is_isomorphic.hpp"

TEST(IsIsomorphic, MultiTest) {
  struct TestInputs {
    std::string s;
    std::string t;
    bool want;
  };
  std::vector<TestInputs> testInputs = {
      {"badc", "baba", false},
      {"foo", "bar", false},
      {"paper", "title", true},
      {"egg", "add", true},
  };

  auto s = Solution{};
  for (const auto &t : testInputs) {
    EXPECT_EQ(t.want, s.isIsomorphic(t.s, t.t));
  }
}