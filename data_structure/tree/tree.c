// g++ -o tree 
// ./tree

#include <iostream>

typedef struct Node_T {
	Node_T *left;
	int val;
	Node_T *right;
} Node;

void bfs(Node* node) {
	Node* q[100];
	int qhead = 0;
  	int	qend = 0;

	if (node != NULL) {
		q[qend++] = node;

	}
	while (qhead != qend) {
		// pop one out
		Node* n = q[qhead++];  
		std::cout << n->val; 
		// push children in
		if (n->left) 
			q[qend++] = n->left;
		if (n->right)
			q[qend++] = n->right;
	}
}

int height(Node* node) {
	if (node == NULL) {
		return 0;
	}
	int left = height(node->left);
	int right = height(node->right);
	if (left>right) 
		return left+1;
	else 
		return right+1;
}

void postfix(Node* node) {
	if (node==NULL) {
		return;
	}
	postfix(node->left);
	postfix(node->right);
	std::cout << node->val;
}

void prefix(Node* node) {
	if (node == NULL) {
		return;
	}
	std::cout << node->val;
	prefix(node->left);
	prefix(node->right);
}

void infix(Node* node) {
	if (node==0) {
		return ;
	}
	infix(node->left);
	std::cout << node->val;
	infix(node->right);
}

Node* insert(Node* node, int val) { 
	if (node == 0) {
		return new Node{0, val, 0};
	}
	if (val < node->val) 
		node->left = insert(node->left, val);
	else 
		node->right = insert(node->right, val);
}

Node* buildBST(int* a, int n) {
	Node *node = 0;
	for (int i=0; i<n; i++) {
		node = insert(node, a[i]);
	}
	return node;
}

int main() {
	int a[] = {7,4,6,8,5,3,2,1,0};
	int n = sizeof(a)/sizeof(int);
	Node* node = buildBST(a, n);
	std::cout << "bfs of a bst: ";  bfs(node); std::cout << std::endl; // 748362510
	std::cout << "infix: ";  infix(node); std::cout << std::endl;  // 012345678 
	std::cout << "prefix: "; prefix(node); std::cout << std::endl; // 743210658
	std::cout << "postfix: "; postfix(node); std::cout << std::endl; // 012356487
	std::cout << "height: "<< height(node) << std::endl;  // 6
	return 1;
}
