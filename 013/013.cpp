#include <iostream>

int main() {
  int target[100][50] = {};
  int ans[55] = {};
  bool carry = false;

  FILE *fp;
  if((fp = fopen("input.txt", "r")) == NULL) return 1;

  for (size_t i=0; i<100; i++) {
    for (size_t j=0; j<50; j++) {
      fscanf(fp, "%d", &target[i][j]);
    }
  }

  for (int i=0; i<100; i++) {
    for (int j=49; j>=0; j--) {
      int tmp = ans[j+5] + target[i][j] + carry;
      if (tmp>=5) {
        ans[j+5] = tmp%10;
        carry = true;
      } else {
        ans[j+5] = tmp;
        carry = false;
      }
    }

    if (carry) {
      int digits_count = sizeof(target[0])/sizeof(int);
      while (carry) {
        int tmp = ans[sizeof(ans)/sizeof(int)-digits_count-1]+1;
        if (tmp>=10) {
          ans[sizeof(ans)/sizeof(int)-digits_count-1] = tmp%10;
          digits_count++;
        } else {
          ans[sizeof(ans)/sizeof(int)-digits_count-1] = tmp;
          carry = false;
        }
      }
    }
  }
  for (size_t i = 0; i < sizeof(ans)/sizeof(int); i++) std::cout << ans[i] << std::flush;
  std::cout << std::endl;
  return 0;
}
