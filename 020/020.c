#include "stdio.h"

int main(void) {
  long result = 1;
  int ans = 0;
  for (int i=20; i>0; i--) result *= i;

  while (result>0) {
    int digit = result%10;
    result = result/10;
    ans += digit;
  }

  printf("%d\n", ans);
};
