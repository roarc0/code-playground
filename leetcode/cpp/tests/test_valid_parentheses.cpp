#include <gtest/gtest.h>

#include "valid_parentheses.hpp"

TEST(ValidParentheses, MultiTest) {
  struct TestInputs {
    std::string expr;
    bool want;
  };
  std::vector<TestInputs> testInputs = {
      {"{[]}", true},
      {"(){}[]", true},
  };

  auto s = Solution{};
  for (const auto &t : testInputs) {
    EXPECT_EQ(t.want, s.isValid(t.expr));
  }
}