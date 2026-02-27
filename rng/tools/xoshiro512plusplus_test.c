#include <stdio.h>
#include "xoshiro512plusplus.c"

int main(){
          s[0] = 12744715263588028796ull;
          s[1] = 16192141852193020578ull;
          s[2] = 16161435109270938784ull;
          s[3] = 9687954255755471688ull; 
		  s[4] = 13813258639892344913ull;
		  s[5] = 7328218089997940832ull;
		  s[6] = 13272147554285316318ull;
		  s[7] = 1545831295407692026ull; // generated with seed 987654321 for splitmix
      printf("xoshiro512plusplusdata := [...]uint64{");
      for (int i=0; i<10; i++){
         for (int j=0; j < 5; j++) printf("%llu,",next());
         printf("\n");
                 if (i % 2) jump();
                 else long_jump();
      }
      printf("%llu}\n",next());
}