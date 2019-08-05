#include <stdio.h>
#define STUD_MAX 30001
int num[STUD_MAX],father[STUD_MAX],conut[STUD_MAX];

int find_father(int x)
{
	if(x != father[x])
		father[x] = find_father(father[x]);
	return father[x];
}

int main(void)
{
	int stud_num,club_num;
	int i,j;
	scanf("%d%d",&stud_num,&club_num);

	for(i=0;i<=stud_num;i++)
	{
		father[i] = i;
	}
	
	
	for(i=1;i<=club_num;i++)
	{
		int temp,last_father;
		scanf("%d",&temp);
		for(j=1;j<=temp;j++)
		{
			int stud;
			scanf("%d",&stud);
			if(num[stud]==0)
			{				
				if(j==1)
				{
					last_father = father[stud];
				}
				num[stud] = stud;
				father[stud] = last_father; 
			}
			else
			{
				if(j==1)
				{
					last_father = find_father(stud);
					father[last_father] = last_father;
				}
				else
				{
					int t = find_father(stud);
					father[t] = last_father;
				}
			}
		}
	}

	int result = 0,index;
	for(i=1;i<=stud_num;i++)
	{
		index = find_father(i);
		conut[index]++;
		if(result < conut[index]) 
		{
			result = conut[index];
		}
	}

	printf("%d\n",result);
	return 0;
}
