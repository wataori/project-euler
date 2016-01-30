#include <iostream>

long ggd(long a, long b)
{
  if (a == b) return a;

  if (a > b) {
    return ggd(b, a - b);
  } else {
    return ggd(b - a, a);
  }
}

long lcm(long a, long b)
{
  return a * b / ggd(a, b);;
}

long loop(long *arr, long length)
{
  if (length == 0) return arr[0];

  long new_arr[20] = {};

  for(int i = 0; i < length; i++)
  {
    new_arr[i] = lcm(arr[i + 1], arr[i]);
  }

  return loop(new_arr, length - 1);
}

int main()
{
  int count = 20;
  long input[20] = {};

  for(int j = 0; j < count; j++)
  {
    input[j] = j + 1;
  }

  std::cout << loop(input, count - 1) << std::endl;

  return 0;
}
