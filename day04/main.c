#include <stdio.h>

void split(int *nn, int n) {
  for (int i = 0; i < 6; i++) {
    nn[5-i] = n % 10;
    n /= 10;
  }
}

int main(void) {
  int a = 0, b = 0;
  int nn[6];
  for (int n = 359282; n <= 820401; n++) {
    split(nn, n);
    int va = 0, vb = 0;
    for (int i = 0; i < 6; ) {
      int j = i;
      int c = nn[j];
      for ( ; j < 6 && c == nn[j]; j++);
      if (j < 6 && c > nn[j]) goto next;
      int diff = j - i;
      va |= diff >= 2;
      vb |= diff == 2;
      i = j;
    }
    a += va;
    b += vb;
next: ;
  }

  printf("a) %d\n", a);
  printf("b) %d\n", b);
}
