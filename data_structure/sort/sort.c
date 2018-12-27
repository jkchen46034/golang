#include <iostream>

void print (int*a, int n);

void insertionsort(int* a, int n) {
	for (int i=1; i<n; i++) {
		for (int j=i-1; j>=0; j--) {
				if (a[j]>a[j+1]) {
					int temp = a[j];
					a[j] = a[j+1];
				  a[j+1] = temp;
					print(a,n);
				}
		}
	}
}

void bubblesort(int* a, int n) {
	for (int i=0; i<n-1; i++) {
		for (int j=0; j<n-i-1; j++) {
			if (a[j]>a[j+1]) {
				 int temp = a[j];
				 a[j] = a[j+1];
				 a[j+1] = temp;
				 print(a,n);
			}
	 	}
	}
}

void print (int *a, int n) {
	for (int i=0; i<n; i++)
		std::cout << a[i];
	std::cout << std::endl;
}

int main() {
	int a[]={5, 1, 4, 2, 8};
	int n = sizeof(a)/sizeof(int);
	std::cout << "bubble sort: "; print(a,n); bubblesort(a, n); 
	int b[]={5, 1, 4, 2, 8};
	n = sizeof(b)/sizeof(int);
	std::cout << "insertion sort: "; print(b, n); insertionsort(b, n); 
	
	return 1;
}
