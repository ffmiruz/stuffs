// https://commaok.xyz/post/lookup_tables/

#include <stdio.h>
#include <stdlib.h>

// https://lemire.me/blog/2021/05/28/computing-the-number-of-digits-of-an-integer-quickly
 int int_log2(unsigned int x) { return 31 - __builtin_clz(x|1); }

// https://www.geeksforgeeks.org/program-to-compute-log-n/
unsigned int Log2n(unsigned int n)
{
    return (n > 1) ? 1 + Log2n(n / 2) : 0;
}

int fast_digit_count(unsigned int x) {
  static unsigned long long table[] = {
      4294967296,  8589934582,  8589934582,  8589934582,  12884901788,
      12884901788, 12884901788, 17179868184, 17179868184, 17179868184,
      21474826480, 21474826480, 21474826480, 21474826480, 25769703776,
      25769703776, 25769703776, 30063771072, 30063771072, 30063771072,
      34349738368, 34349738368, 34349738368, 34349738368, 38554705664,
      38554705664, 38554705664, 41949672960, 41949672960, 41949672960,
      42949672960, 42949672960};
  return (x + table[int_log2(x)]) >> 32;
}

int slow_digit_count(unsigned int x)
{
  static unsigned int table[] = {9, 99, 999, 9999, 99999,
    999999, 9999999, 99999999, 999999999};
    int y = (9 * int_log2(x)) >> 5;
    y += x > table[y];
    return y + 1;
}

int main() {
    int x;
    
    x = fast_digit_count(2147483647);
    printf("fast: %d\n", x);
    x = slow_digit_count(9);
    printf("slow: %d\n", x);
    return 0;
}

// https://www.geeksforgeeks.org/program-to-compute-log-a-to-any-base-b-logb-a/
// int log_base(int base, int n)
// {
//     return (n > base - 1)
//                ? 1 + logxx(n / base, base): 0;
// }