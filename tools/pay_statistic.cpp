#include<iostream>
#include<fstream>
using namespace std;

int main(int argc, char *argv[]){

  if(argc != 3){ 
    cerr << "Usage: ./pay_statistic outfile time" << endl;
    return 1;
  }

  // different kinds of payment 
  cout << "Please input the kind" << endl;
  cout << "[0]livelihood" << endl;
  cout << "[1]campus card" << endl;
  cout << "[2]breakfast" << endl;
  cout << "[3]lunch" << endl;
  cout << "[4]dinner" << endl;
  cout << "[5]snacks" << endl;
  cout << "[6]traffic" << endl;
  cout << "[7]game&play" << endl;

  // input
  int kind; 
  double money[8]={0.0};  
  double temp_money;
  while(cin >> kind){
    cin >> temp_money;
    money[kind] += temp_money;
    cout << "kind: " << kind << endl;
    cout << "payment: " << money[kind] << endl;
  }
  
  // output to file
  ofstream out;
  out.open(argv[1]);

  out << "########################"  << endl;
  out << "Report for " << argv[2] << endl;
  out << "[0]livelihood:  " << money[0] << endl;
  out << "[1]campus card: " << money[1] << endl;
  out << "[2]breakfast:   " << money[2] << endl;
  out << "[3]lunch:       " << money[3] << endl;
  out << "[4]dinner:      " << money[4] << endl;
  out << "[5]snacks:      " << money[5] << endl;
  out << "[6]traffic:     " << money[6] << endl;
  out << "[7]game&play:   " << money[7] << endl;
  out << "########################"  << endl;

  out.close();

  return 0;
   		
}
