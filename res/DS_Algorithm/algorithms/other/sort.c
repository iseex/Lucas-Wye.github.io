#include<stdio.h>
#define MAX 100
void select_sort(int *a,int len){
	int i,j,index,temp;
	for(i=0;i<len;i++){
		index = i;
		for(j=i+1;j<len;j++){
			if(a[index]>a[j]){
				index = j;
			}
		}
		if(index!=i){
			temp = a[i];
			a[i] = a[index];
			a[index] = temp; 
		}

	}
	for(i=0;i<len;i++){
		printf("%d ", a[i]);
	}

}

void bubble_sort(int *a,int len){
	int i,j,temp;
	for(i=0;i<len;i++){
		for(j=0;j<len-1-i;j++){
			if(a[j]>a[j+1]){
				temp = a[j];
				a[j] = a[j+1];
				a[j+1] = temp;
			}
		}
	}
	for(i=0;i<len;i++){
		printf("%d ", a[i]);
	}

}

int main(void){
	// int a[MAX];
	int a[12] = {2,1,3,6,5,44,20,41,33,50,10,9};
	select_sort(a,12);
	printf("\n");
	bubble_sort(a,12);
	printf("\n");


}