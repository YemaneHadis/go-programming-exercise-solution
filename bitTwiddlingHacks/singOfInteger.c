#include<stdio.h>
#include<stdbool.h>

int sing (int u, int v){
    int sing1;
    int sing2;

    sing1 = u >> (sizeof(int) * __CHAR_BIT__-1);
    sing2 = v >> (sizeof(int) * __CHAR_BIT__-1);

    printf("Sing of U\t%d\n", sing1);
    printf("Sing of V\t%d\n", sing2);

    /* 
    - HOW WE ARE GETTING THE SING OF AN INTEGER 
        1. Binary for 12 in 8 bit is 00001100 and rest left side bits filled with 0
        2. size of int (4) * __CHAR_BIT__ (8 in most computer) -1 = 31
        4. so if you shift 31 bit to the right then you will get - 32 - zeros which is 0
        5. for -11 binary is 1111 1111 1111 1111 1111 1111 1111 010
        and with singed -ve integer right shift 
        Bits are shifted to the right.
        The leftmost bits are filled with the sign bit.
        The rightmost bits are discarded.
    
     */


    return 0;
}

int checkSameSign(int u, int v){
    bool f = ((u^v)<0);
    printf("%d and %d have the different sing evaluated to %s\n",u,v,f?"True":"False");
    return 0;

    /*
        if they have different sing the left most bit will be different and this will be 
        evaluated to 1 with xor so the decimal representation of the number will be less than zero
    */
}
int main(int argc, char const *argv[])
{
    sing(12,-11);
    checkSameSign(12,11);
    checkSameSign(12,-12);
    return 0;
}
