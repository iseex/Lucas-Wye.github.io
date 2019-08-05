#include<stdio.h>
#include<string.h>

int Max_same(char *input,int s,int len)
{
    int len1 = s;
    int len2 = len - len1;
    int i,j;
    int count = 0;

	//动态规划    
    int c[50][50]={0};
    for(i=1;i<=len1;i++)
    {
    	for(j=1;j<=len2;j++)
    	{
    		if(input[i-1] == input[len1+j-1])
    		{
    			c[i][j] = c[i-1][j-1]+1;
    		}
            else
            {
                c[i][j] = (c[i-1][j]>c[i][j-1])?c[i-1][j]:c[i][j-1];                
            }
    	}
    }
    count = c[len1][len2];


    return count;    
}
int main(void)
{
    char input[51];
    strcpy(input,"aababbababbabbbbabbabb");
//    scanf("%s",input);
    int len = strlen(input);
    int i;
    int same = 0;    
    for(i=0;i<(len-1);i++)
    {
        // printf("%c",input[i] );
        int temp = Max_same(input,i,len);        
        printf("i = %d, temp = %d\n",i,temp );
        if(same < temp)
            same = temp;
    }
    printf("%d",2*same);
    return 0;
}
