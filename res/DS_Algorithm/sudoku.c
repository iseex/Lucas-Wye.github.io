/*
输入描述:
输入9行，每行为空格隔开的9个数字，为0的地方就是需要填充的。

输出描述:
输出九行，每行九个空格隔开的数字，为解出的答案。

算法利用回溯的思想：
（1）从第一个空白处开始，找到其候选解
排除同行、同列以及同一小九宫格的所有出现过的数字，剩下未出现的数字都是候选解的第一个值填入数独。
（2）对第二个空白执行第一步（前面所填入的数字对此空白处有影响）
当出现某个空白的候选解个数为0时，就开始回溯，找到第一个候选解多于一个的，将其在使用的候选解设为不可取（本程序取值为-1），
找到其下一个候选解，继续上面的步骤！
（3）直到所有空白处填满，运算完成，输出结果！

*/

#include <stdio.h>
#include <stdlib.h>

typedef struct node
{
	int col;
	int row;
	// value存储0-9每个值是否在列表
	int value[10];
}Node;

int findvalue(int sudoku[9][9], Node * node);

void Print(int sudoku[9][9]){
	int i,j;
	//打印数独
	printf("-------------------------------------\n");
	for(i=0; i<9; i++)
	{
		printf("| ");
		for(j=0; j<9; j++){
			printf("%2d ", sudoku[i][j]);
			if((j+1)%3 == 0){
				printf(" | ");
			}
		}
		printf("\n");
		if((i+1)%3 == 0){
			printf("-------------------------------------\n");
		}
	}
}

int main(void)
{
	int sudoku[9][9] = 
  {
		{9, 0, 0, 0, 0, 0, 1, 8, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 3, 0, 0, 0},
		{0, 4, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 2, 0},
		{0, 3, 6, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 6, 0, 0, 0, 0, 3},
		{0, 5, 0, 0, 0, 0, 4, 0, 0}
	};
	printf("Input:\n");
	Print(sudoku);
	printf("\n");

	int i, j, num_of_empty = 0;
	int index, temp = 0;
	
	//计算所给数独中待填入的空白数
	for(i=0; i<9; i++)
		for(j=0; j<9; j++)
			if(sudoku[i][j]==0)
				num_of_empty++;

	//为回溯栈分配空间
	Node * node_stack = (Node *)malloc
	(sizeof(struct node) * num_of_empty);

	int k = 0;
	//回溯法求解数独
	while(num_of_empty) {
		for(i=0; i<9; i++)
			for(j=0; j<9; j++)
				if(sudoku[i][j]==0) {
					//初始化栈中存储候选值的数组
					for(index=0; index<10; index++)
						(node_stack + k)->value[index] = 0;

					(node_stack + k)->col = i;
					(node_stack + k)->row = j;
					sudoku[i][j] = findvalue(sudoku, node_stack + k);
					// 回溯
					if(sudoku[i][j]==-1) {
						sudoku[i][j] = 0;
						k--;
						while((node_stack + k)->value[0]==1) {
							sudoku[(node_stack + k)->col][(node_stack + k)->row] = 0;
							num_of_empty++;
							k--;
						}
						(node_stack + k)->value[0]--;
						i = (node_stack + k)->col;
						j = (node_stack + k)->row;
						for(index=1; index<10; index++)
							if((node_stack + k)->value[index]==0) {
								(node_stack + k)->value[index] = -1;
								break;
							}
						for(index=1; index<10; index++)
							if((node_stack + k)->value[index]==0) {
								sudoku[i][j] = index;
								break;
							}
					}
					k++;
				}
		//当栈空，说明数独错误，无解
		if(k==0) {
			printf("此数独无解！\n");
			free(node_stack);
			return 0;
		}
		num_of_empty--;
	}

	free(node_stack);
	printf("Output:\n");
	Print(sudoku);
	printf("\n");
	return 0;
}

int findvalue(int sudoku[9][9], Node * node) {
	int m, n;
	int i = node->col, j = node->row;

	// 检查行和列缺少的值
	int t1,t2,t3;
	for(m=0; m<9; m++) {
		// 行
		t1 = sudoku[i][m];
		if(node->value[t1]==0)
			node->value[t1] = t1;
		// 列
		t2 = sudoku[m][j];
		if(node->value[t2]==0)
			node->value[t2] = t2;
	}
	// 检查九宫格缺少的值
	for(m=0; m<3; m++) {
		for(n=0; n<3; n++) {
			t3 = sudoku[i/3*3+m][j/3*3+n];
			if(node->value[t3]==0)
				node->value[t3] = t3;
		}
	}

	// 计算个数
	for(m=1; m<10; m++)
		if(node->value[m]==0)
			node->value[0]++;

	// 找到第一个为0的数
	for(m=1; m<10; m++)
		if(node->value[m]==0)
			break;

	if(node->value[0]==0)
		return -1;
	else
		return m;
}

