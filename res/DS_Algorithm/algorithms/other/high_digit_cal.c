#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#define MAX 10001

int multiply(int *sum,int *a,int *b,int len_a,int len_b)
{
	int i,j,len_sum = 0;
	memset(sum,0,sizeof(sum));
	
	for(i=1;i<=len_a;i++) /*用数组模拟运算*/
	{       
		for(j=1,len_sum=i-1;j<=len_b;j++)
		{
			sum[++len_sum] += a[i] * b[j];
		}
	}
	
	for(i=1;i<=len_sum;i++)/*进位处理*/
	{
		if(sum[i] >= 10)
		{
			if (sum[len_sum] >= 10)
			{
				len_sum ++;
			}
			sum[i+1] += sum[i] / 10;
			sum[i] %= 10;
		}
	}

	return len_sum ;
}

int add_function(int *sum,int *a,int *b,int len_a,int len_b)
{
	int i,j,len_sum = 0;
	int len_min,len_max;
	len_min = (len_a>len_b)?len_b:len_a;
	len_max = (len_a>len_b)?len_a:len_b;
	
	for(i=1;i<=len_min;i++)
	{
		sum[++len_sum] = a[i] + b[i];
	}
	if(len_min==len_a)
	{
		for(i=len_min+1;i<=len_max;i++)
		{
			sum[++len_sum] = b[i];
		}	
	}
	else
	{
		for(i=len_min+1;i<=len_max;i++)
		{
			sum[++len_sum] = a[i];
		}		
	}

	for(i=1;i<=len_sum;i++)/*进位处理*/
	{
		if(sum[i]>= 10)
		{
			if (sum[len_sum] >= 10)
			{
				len_sum ++;
			}
			sum[i+1] += sum[i] / 10;
			sum[i] %= 10;
		}
	}

	return len_sum;
}

int sub_function(int *sum,int *a,int *b,int len_a,int len_b)
{
	//尚存在不足，例如对负数的处理，对a>b的处理
	int i,j,len_sum = 0;
	int len_min,len_max;
	len_min = (len_a>len_b)?len_b:len_a;
	len_max = (len_a>len_b)?len_a:len_b;

	for(i=1;i<=len_min;i++)
	{
		sum[++len_sum] = a[i] - b[i];
	}
	if(len_min==len_a)
	{
		for(i=len_min+1;i<=len_max;i++)
		{
			sum[++len_sum] = -b[i];
		}		
	}
	else
	{
		for(i=len_min+1;i<=len_max;i++)
		{
			sum[++len_sum] = a[i];
		}		
	}

	for(i=1;i<=len_sum;i++)/*进位处理*/
	{
		if(sum[i] < 0)
		{
			// if (sum[len_sum] < 0)
			// {
			// 	len_sum ++;
			// }
			sum[i+1] -= 1;
			sum[i] += 10;
		}
	}

	return len_sum;

	
}

int main(void)
{
	int a[MAX]={0};
	int b[MAX]={0};
	int sum[MAX*2]={0};
	int add[MAX+1]={0};
	int sub[MAX+1]={0};

	int len_a=0;
	int len_b=0;
	int len_sum=0;
	int len_add=0;
	int len_sub=0;

	int i,j;

	//for input
	char sa[MAX],sb[MAX];    
	scanf("%s %s",sa,sb);
	len_a = strlen(sa);
	len_b = strlen(sb);

	printf("第一个数字的反向是：");
	//数组的低地址存放低位数据	
	for(i=1,j=len_a-1;i<=len_a;i++,j--)
	{
		a[i] = sa[j] - '0';
		printf("%d",a[i]);
	}
	printf("\n");
	
	printf("第二个数字的反向是：");
	for(i=1,j=len_b-1;i<=len_b;i++,j--)
	{
		b[i] = sb[j] - '0';
		printf("%d",b[i]);
	}
	printf("\n");
	
	
	len_sum = multiply(sum,a,b,len_a,len_b);
	printf("乘法计算结果为：");
	for(i=len_sum;i>0;i--)
	{
		printf("%d",sum[i]);
	}
	printf("\n");
	
	len_add = add_function(add,a,b,len_a,len_b);
	printf("加法计算结果为：");
	for(i=len_add;i>0;i--)
	{
		printf("%d",add[i]);
	}
	printf("\n");

	len_sub = sub_function(sub,a,b,len_a,len_b);
	printf("减法计算结果为：");
	for(i=len_sub;i>0;i--)
	{
		printf("%d",sub[i]);
	}
	printf("\n");




	
	return 0;
}
