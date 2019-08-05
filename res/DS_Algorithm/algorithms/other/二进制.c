int dif=m^n;//先将二者做异或运算，得到结果；
int cnt=0;
while(dif!=0){
    dif=dif&(dif-1);//计算1的个数
    cnt++;
}   