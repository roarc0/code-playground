#include "valid_parentheses.hpp"

#include <stack>

auto oppositeParenthesis(char c) {
  switch (c) {
  case '(':
    return ')';
  case '[':
    return ']';
  case '{':
    return '}';
  default:
    return c;
  }
}

bool isOpenParenthesis(char c) { return c == '(' || c == '[' || c == '{'; }

bool isCloseParenthesis(char c) { return c == ')' || c == ']' || c == '}'; }

bool Solution::isValid(const std::string &s) {
  std::stack<char> st;
  for (auto c : s) {
    if (!st.empty() && isCloseParenthesis(c) &&
        oppositeParenthesis(st.top()) == c) {
      st.pop();
    } else if (isOpenParenthesis(c)) {
      st.push(c);
    } else {
      return false;
    }
  }
  return st.empty();
}
