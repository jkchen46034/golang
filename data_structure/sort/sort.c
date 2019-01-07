#include <iostream>

int n=0;
void print (int*a, int n);
void merge(int *a, int l,  int m, int r);
int findpivot(int *a, int l, int r);

int findpivot(int *a, int l, int r) {
	int pivot = a[r-1];
	int pindex = l;
	for (int i=l ; i<r-1; i++) {
		if (a[i] < pivot) {
			int temp = a[i];
			a[i] = a[pindex];
			a[pindex] = temp;
			pindex ++;
			print(a, n);
		}
	}	
	int temp = a[pindex];
	a[pindex] = a[r-1];
	a[r-1] = temp;
	print(a,n);
  return pindex;
}

void quicksort(int*a, int l, int r) {
	if (r-l<=1) {
		return;
	}
	int pindex = findpivot(a, l, r);
	quicksort(a, l,  pindex);
	quicksort(a, pindex+1, r);
}

void mergesort(int*a, int l, int r) {
	if (r-l<=1) {
		return;
	}
	int middle = (l+r)/2;
	mergesort(a, l, middle);
	mergesort(a, middle, r);
  merge(a, l, middle, r);
}

void merge(int *a, int l,  int m, int r) {
	for (int i=m; i<r; i++) {
		for (int j=i-1; j>=l; j--) {
			if (a[j] > a[j+1]) {
				int temp = a[j];
				a[j] = a[j+1];
				a[j+1] = temp;
				print(a,n);
			} else  {
				break;
			}
		}
	}
}

void selectionsort(int*a, int n) {
	for (int i=0; i<n-1; i++) {
		for (int j=i; j<n; j++) {
			if (a[i] > a[j]) {
				int temp =  a[j];
				a[j] = a[i];
				a[i] = temp;
				print(a,n);
			}
		}
	
	}
}

void insertionsort(int* a, int n) {
	for (int i=1; i<n; i++) {
		for (int j=i-1; j>=0; j--) {
			if (a[j] > a[j+1]) {
				int temp = a[j];
				a[j] = a[j+1];
				a[j+1] = temp;
				print(a,n);
			}
		}
	}
}

void bubblesort(int* a, int n) {
	for (int i=0; i<n-1; i++)	{
		for (int j=0; j<n-1-i; j++) {
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
		std::cout << a[i] <<" ";
	std::cout << std::endl;
}

int main() {
	int a[]={5, 1, 4, 8, 2};
	int n = sizeof(a)/sizeof(int);
	std::cout << "bubble sort:\n"; print(a,n); bubblesort(a, n); 
	int b[]={5, 1, 4, 8, 2};
	n = sizeof(b)/sizeof(int);
	std::cout << "insertion sort:\n"; print(b, n); insertionsort(b, n); 
	int c[]={5, 1, 4, 8, 2};
	n = sizeof(c)/sizeof(int);
	std::cout << "selection sort:\n"; print(c, n); selectionsort(c, n); 
	int d[]={15,13,1,7,5,9,6,4};
	n = sizeof(d)/sizeof(int);
	::n = n;
	std::cout << "merge sort:\n"; print(d, n); mergesort(d, 0, n); 
	int e[]={15,13,1,7,5,9,6,4};
	n = sizeof(e)/sizeof(int);
	::n = n;
	std::cout << "quick sort:\n"; print(e, n); quicksort(e, 0, n); 
	
	return 1;
}
/*
$ g++ sort.c
./a.out
bubble sort:
5 1 4 8 2 
1 5 4 8 2 
1 4 5 8 2 
1 4 5 2 8 
1 4 2 5 8 
1 2 4 5 8 
insertion sort:
5 1 4 8 2 
1 5 4 8 2 
1 4 5 8 2 
1 4 5 2 8 
1 4 2 5 8 
1 2 4 5 8 
selection sort:
5 1 4 8 2 
1 5 4 8 2 
1 4 5 8 2 
1 2 5 8 4 
1 2 4 8 5 
1 2 4 5 8 
merge sort:
15 13 1 7 5 9 6 4 
13 15 1 7 5 9 6 4 
13 1 15 7 5 9 6 4 
1 13 15 7 5 9 6 4 
1 13 7 15 5 9 6 4 
1 7 13 15 5 9 6 4 
1 7 13 15 5 9 4 6 
1 7 13 15 5 4 9 6 
1 7 13 15 4 5 9 6 
1 7 13 15 4 5 6 9 
1 7 13 4 15 5 6 9 
1 7 4 13 15 5 6 9 
1 4 7 13 15 5 6 9 
1 4 7 13 5 15 6 9 
1 4 7 5 13 15 6 9 
1 4 5 7 13 15 6 9 
1 4 5 7 13 6 15 9 
1 4 5 7 6 13 15 9 
1 4 5 6 7 13 15 9 
1 4 5 6 7 13 9 15 
1 4 5 6 7 9 13 15 
quick sort:
15 13 1 7 5 9 6 4 
1 13 15 7 5 9 6 4 
1 4 15 7 5 9 6 13 
1 4 7 15 5 9 6 13 
1 4 7 5 15 9 6 13 
1 4 7 5 9 15 6 13 
1 4 7 5 9 6 15 13 
1 4 7 5 9 6 13 15 
1 4 5 7 9 6 13 15 
1 4 5 6 9 7 13 15 
1 4 5 6 7 9 13 15 
*/
