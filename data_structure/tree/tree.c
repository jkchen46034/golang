// g++ -o tree 
// ./tree

#include <iostream>
#include <cstddef>

typedef struct Node_T {
	Node_T *left;
	int val;
	Node_T *right;
} Node;

/*
    0
   / \
  1   2
 / \ / \
3  4 8  5
/\
7 6

*/


// 012348576
void bfs(Node *node) {
	Node *q[100];
	int qhead = 0;
  int	qend = 0;

	if (node != NULL) {
		q[qend++] = node;

	}
	while (qhead != qend) {
		// pop one out
		Node *n = q[qhead++];  
		std::cout << n->val; 
		// push children in
		if (n->left) 
			q[qend++] = n->left;
		if (n->right)
			q[qend++] = n->right;
	}
}

// 4
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

// 763418520
void postfix(Node* node) {
	if (node==NULL) {
		return;
	}
	postfix(node->left);
	postfix(node->right);
	std::cout << node->val;
}

// 013764285
void prefix(Node* node) {
	if (node == NULL) {
		return;
	}
	std::cout << node->val;
	prefix(node->left);
	prefix(node->right);
}

// 736140825 
void infix(Node *node) {
	if (node==0) {
		return ;
	}
	infix(node->left);
	std::cout << node->val;
	infix(node->right);
}

int main() {
	Node *node7 = new Node{NULL, 7, NULL};
  Node *node4 = new Node{NULL, 4, NULL};
  Node *node6 = new Node{NULL, 6, NULL};
  Node *node8 = new Node{NULL, 8, NULL};
  Node *node5 = new Node{NULL ,5, NULL};
  Node *node3 = new Node{node7,3, node6};
  Node *node2 = new Node{node8,2, node5};
  Node *node1 = new Node{node3,1, node4};
  Node *node0 = new Node{node1,0, node2};

	std::cout << "infix: ";  infix(node0); std::cout << std::endl;
	std::cout << "prefix: "; prefix(node0); std::cout << std::endl;
	std::cout << "postfix: "; postfix(node0); std::cout << std::endl;
	std::cout << "height: "<< height(node0) << std::endl;
	std::cout << "bfs: "; bfs(node0); std::cout << std::endl;
	return 1;
}
