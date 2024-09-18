#include<stdio.h>
#include<stdbool.h>


// Naive way 
int countSetBit(int u){
 int count = 0;
 int number = u;
 for (; u > 0; u >>=1)
 {
   count += u&1;
 }
 printf("Number Of Set bits in \t%d  are\t %d  \n",number, count);
 return 0;
}



//Kernighan way of doing it 
// this appraoch would eliminate unecessary steps 
// for power two and so 
int countSetBits2(int u){
    int number = u;
    int count = 0;
    for(;u >0; u&=u-1){
        count++;
    }
    printf("Using Kernighan Method Number Of Set bits in \t%d  are\t %d  \n",number, count);

    return 0;
}

//Sign extending from a variable bit-width
int singExtending(){
    unsigned int b;
    int x;
    int y;
    int const m = 1U <<(b-1) ;//mask can be precomputed if b is fixed
    //x = x &   
    return 0;
}


int isPowOfTwo(unsigned int u){
    
    bool isPowerOfTwo = (u&(u-1)) == 0;
    // in the above 0 is classified as power of two wrongly 
   
    bool isPowerOfTwo2 =  u && !(u & (u - 1));

    printf("%d is Power of Two Evaluated to %s\n", u,isPowerOfTwo?"True":"False");
    return 0;


}
int abs(int u){
    /**
     *  If u is +ve number then mask will be 0 u xor mask will be resulted to u
     *  If u is -ve number then mask will be -1 then (u-1) xor -1 will be resulted to abs(u)
     * 
     */
    unsigned int r;
    int const mask = u >> (sizeof(int) * __CHAR_BIT__-1);
    r = (u-mask) ^ mask;
    return 0;
}
int sing (int u, int v){
    int sing1;
    int sing2;

    sing1 = u >> (sizeof(int) * __CHAR_BIT__-1);
    sing2 = v >> (sizeof(int) * __CHAR_BIT__-1);

    printf("Sing of U\t%d\n", sing1);
    printf("Sing of V\t%d\n", sing2);
    printf("Unsinged Int Size of V\t%lu\n", sizeof(unsigned int));

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

    countSetBit(11);
    countSetBit(16);

    /**
     *  this approch is more efficient from number of 
     *  operation perspective  but we are going to do 32 operation in the 
     *  worest case if the number is int and int is 4 byte 
     * 
     * **/
    countSetBits2(11);
    countSetBits2(16);
    return 0;
}
