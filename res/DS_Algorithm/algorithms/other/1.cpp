/*
牛牛在地上捡到了一个手套，他带上手套发现眼前出现了很多个小人，当他打一下响指，这些小人的数量就会发生以下变化：
如果小人原本的数量是偶数那么数量就会变成一半，如果小人原本的数量是奇数那么数量就会加一。
现在牛牛想考考你，他要打多少次响指，才能让小人的数量变成1。 
*/ 
#include <bits/stdc++.h>
using namespace std;
char s[105];
int main() 
{
    int l, ans = 0, tops = 1, now;
    scanf("%s", s + 1);
	
    l = strlen(s + 1);
    			
	for(int i = 1; i <= l; i++) s[i] -= '0';
    
    for(int i;i<=l;i++)    
	{
		printf("%c ",s[i]);
	}
    
	while(tops < l || s[tops] != 1) 
    {
        ans++;
        if(s[l] % 2) 
        {
            now = l;
            s[now]++;
            while(s[now] == 10) 
            {
                s[now - 1]++;
                s[now] = 0;
                now--;
                if(now < tops)
                {
                    tops = now;
                }
            }
        } 
		else 
        {
            for(int i = tops; i <= l; i++) 
            {
                if(s[i] % 2)
                {
                    s[i + 1] += 10;
                }
                s[i] = s[i] / 2;
                if(i == tops && s[i] == 0)
                { 
                    tops++;
                }
            }
        }
    }
    printf("%d\n", ans);
    return 0;
}
