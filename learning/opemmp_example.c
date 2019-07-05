#include<stdio.h>
#include<omp.h>

int main(void){
  const int n = 10;
  int arr[n];
#pragma omp parallel for     
    for(int i=0;i<n;i++)
    {
      arr[i] = i;
      printf("%d\n",i);
    }
  
  int j;
  for(j=0;j<n;j++){
    //printf("%d\n",arr[j]);
  }
  return 0;  
}	
