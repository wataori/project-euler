#include <stdio.h>

int seq[40] = {1, 2};
int limit, tmp;
int sum = 0;

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
    }
  }

  for (int j = 0; j < limit; j++ )
  {
    if (seq[j] % 2 == 0) {
      sum += seq[j];
    }
  }

  printf("%d\n", sum);

  return 0;
}
