#include<stdio.h>
#include<string.h>
int main(void)
{
    int x1,x2,k1,k2;
    int v1[451],v2[451];
    scanf("%d%d%d%d",&x1,&k1,&x2,&k2);
    
    int len_1,len_2;
    int i,j;
    i=1;
    while(x1!=0)
    {
        v1[i] = x1 % 10;
        i++;
        x1 = x1 /10;        
    }
    len_1 = i - 1;
    i = 1;
    while(x2!=0)
    {
        v2[i] = x2 % 10;
        i++;
        x2 = x2 /10;        
    }
    // printf("%d %d\n",v1[4],v2[4] );

    
    len_2 = i - 1;
    
    // printf("The first\n");
    // for(i=1;i<=len_1;i++)
    // {
    //     printf("%d", v1[i]);
    // }
    // printf("\n");
    // printf("The second\n");
    // for(i=1;i<=len_2;i++)
    // {
    //     printf("%d", v2 [i]);
    // }
    // printf("\n");
    // printf("len_1 = %d ,k1 = %d\n", len_1,k1);
    // printf("len_2 = %d ,k2 = %d\n", len_2,k2);

    if(len_1*k1 > len_2 * k2)
    {
        printf("Greater\n");
    }
    else if(len_1*k1 < len_2 * k2)
    {
        printf("Less\n");
    }
    else
    {
        for(i=1;i<=k1*len_1;i++)
        {
            if(0);
           
            else
            {
                v1[i] = v1[(i%(len_1)!=0)?(i%(len_1)):len_1];
                v2[i] = v2[(i%(len_2)!=0)?(i%(len_2)):len_2];
            }
            
            if(v1[i] > v2[i])
            {
                printf("Greater,i = %d,v1[i] = %d, v2[i] = %d\n",i,v1[i],v2[i]);
                return 0;
            }
            else if(v1[i] < v2[i])
            {
                printf("Less,i = %d,v1[i] = %d, v2[i] = %d\n",i,v1[i],v2[i]);                
                return 0;
            }                            
        }
        printf("Equal\n");
        
        
    }
    return 0;
    
    
    
}