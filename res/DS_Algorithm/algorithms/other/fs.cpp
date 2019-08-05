#include<stdio.h>
#include<queue>
#include<stack>
using namespace std;
#define MAX_X 100
#define MAX_Y 100

class node{
public:
    int x;
    int y;
};

int G[MAX_X][MAX_Y][2];//[0]数据,[1]标记


void bfs(int start_x,int start_y,int x_max,int y_max,int next[][2]){
    //初始化一个队列queue并将这个顶点放入队列
    queue<node> Q;
    node input;
    input.x = start_x;
    input.y = start_y;
    printf("%d -> ",G[input.x][input.y][0]);
    Q.push(input);

    node head;
    int i;
    int nx,ny;

    while(!Q.empty()){//queue不为空
        head = Q.front();//访问队头顶点
        G[head.x][head.y][1] = 1;//标记为已访问
        Q.pop();//出队
        for(i=0;i<4;i++){//遍历所有邻接顶点
            nx = head.x + next[i][0];
            ny = head.y + next[i][1];
            input.x = nx;
            input.y = ny;
            if(nx >= x_max || nx < 0 || ny >= y_max || ny < 0){
                continue;
            }
            if(G[nx][ny][1]==0){//未访问，入队
                Q.push(input);
                G[nx][ny][1] = 1;
                printf("%d -> ", G[nx][ny][0]);
            }
        }
            
    }

}

void dfs(int start_x,int start_y,int x_max,int y_max,int next[][2]){
    //初始化一个栈，并将这个顶点入栈
    stack<node> S;
    node input;
    input.x = start_x;
    input.y = start_y;
    printf("%d -> ",G[input.x][input.y][0]);
    S.push(input);

    node head;
    int i;
    int nx,ny;
    while(!S.empty()){//栈非空
        head = S.top();//访问栈顶顶点
        G[head.x][head.y][1] = 1;//标记为已访问
        S.pop();//将栈顶顶点出栈
        for(i=0;i<4;i++){//遍历所有邻接顶点
            nx = head.x + next[i][0];
            ny = head.y + next[i][1];
            input.x = nx;
            input.y = ny;
            if(nx >= x_max || nx < 0 || ny >= y_max || ny <0){
                continue;
            }
            if(G[nx][ny][1]==0){
                S.push(input);
                G[nx][ny][1] = 1;
                printf("%d -> ", G[nx][ny][0]);
            }
        }

    }
}

int main(void)
{
    int i,j;
    int count = 1;
    int x_max = 3;
    int y_max = 3;

    int next[4][2]={
        {0,1},//right
        {0,-1},//left
        {1,0},//up
        {-1,0}//down
    };    

    for(i=0;i<x_max;i++){
        for(j=0;j<y_max;j++){
            G[i][j][0] = count;
            G[i][j][1] = 0;
            count++;
            printf("%d ", G[i][j][0]);
        }
        printf("\n");
    }

    printf("bfs:\n");
    bfs(0,0,x_max,y_max,next);
    printf("\n");

    for(i=0;i<x_max;i++){
        for(j=0;j<y_max;j++){
            G[i][j][1] = 0;
        }
    }

    printf("dfs:\n");
    dfs(0,0,x_max,y_max,next);
    printf("\n");

    return 0;

}






