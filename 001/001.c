#include <stdio.h>

int sum = 0;

int main(void)
{
  for (int i = 1; i < 1000; i++) {
    if (i % 3 == 0) {
      sum += i;
    } else if (i % 5 == 0) {
      sum += i;
    }
  }

  printf("%d\n", sum);

  return 0;
}
