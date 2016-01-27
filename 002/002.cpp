#include <stdio.h>

int seq[40] = {};
int limit;
int sum = 0;

int main(void)
{
  for (int i = 0; i <= 40; i++)
  {
    if (i < 2) {
      seq[i] = i + 1;
    } else if (4000000 <= seq[i-2] + seq[i-1]) {
      limit = i;
      break;
    } else {
      seq[i] = seq[i-2] + seq[i-1];
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
