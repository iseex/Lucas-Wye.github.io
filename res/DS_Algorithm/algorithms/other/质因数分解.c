/*
功能:输入一个正整数，按照从小到大的顺序输出它的所有质数的因子（如180的质数因子为2 2 3 3 5 ）
最后一个数后面也要有空格
*/

#include <iostream>
using namespace std;
int main(void)
{
    long input;
    //cin >> input;
    while (cin >> input)
    {
        while (input != 1)
        {
            for (int i = 2; i <= input; i++)
            {
                if (input % i == 0)
                {
                   input /= i;
                   cout << i << ' ';
                    break;
                }
            }
        }
         
    }
     
     
     
    return 0;
}
