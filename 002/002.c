#include <stdio.h>

int seq[40] = {1, 2};
int limit, tmp;
int sum = 2;

int main(void)
{
  for (int i = 2; i <= 40; i++)
  {
    tmp = seq[i-2] + seq[i-1];
    if (4000000 <= tmp) {
      limit = i;
      break;
    } else {
      seq[i] = tmp;
      if (tmp % 2 == 0) sum += tmp;
    }
  }

  printf("%d\n", sum);

  return 0;
}
