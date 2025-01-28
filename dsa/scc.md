# SCC - Strongly Connected Components
The **Strongly Connected Components (SCC)** algorithm identifies strongly connected components in a directed graph. A strongly connected component is a maximal subgraph where every node is reachable from every other node in the same component.

### Steps for SCC using **Kosaraju's Algorithm**
Kosaraju's algorithm is one of the most commonly used methods for finding SCCs, and it consists of the following steps:

1. **Perform a DFS to compute the finishing times**:
   - Traverse the graph in a Depth-First Search (DFS) manner and push nodes to a stack in the order of their "finishing times" (when all its descendants are visited).
   
2. **Transpose the graph**:
   - Reverse the direction of all edges in the graph.

3. **DFS on the transposed graph**:
   - Process nodes in the order they were added to the stack (from step 1).
   - Each DFS traversal on the transposed graph will give one strongly connected component.

---

### Kosaraju's Algorithm Implementation in C++

Here is the C++ implementation:

```cpp
#include <iostream>
#include <vector>
#include <stack>
using namespace std;

class Graph {
    int V; // Number of vertices
    vector<vector<int>> adj; // Adjacency list

    // Helper function for DFS
    void dfs(int v, vector<bool>& visited, stack<int>& finishStack) {
        visited[v] = true;
        for (int neighbor : adj[v]) {
            if (!visited[neighbor])
                dfs(neighbor, visited, finishStack);
        }
        finishStack.push(v); // Push to stack after finishing the node
    }

    // Helper function to perform DFS on the transposed graph
    void dfsTranspose(int v, vector<bool>& visited, vector<int> transposed[], vector<int>& component) {
        visited[v] = true;
        component.push_back(v);
        for (int neighbor : transposed[v]) {
            if (!visited[neighbor])
                dfsTranspose(neighbor, visited, transposed, component);
        }
    }

public:
    Graph(int V) : V(V) {
        adj.resize(V);
    }

    // Add edge to the graph
    void addEdge(int u, int v) {
        adj[u].push_back(v);
    }

    // Find and print all SCCs
    void findSCCs() {
        stack<int> finishStack; // Stack to store finish times
        vector<bool> visited(V, false);

        // Step 1: Fill the stack with finishing times
        for (int i = 0; i < V; ++i) {
            if (!visited[i])
                dfs(i, visited, finishStack);
        }

        // Step 2: Transpose the graph
        vector<int> transposed[V];
        for (int u = 0; u < V; ++u) {
            for (int v : adj[u]) {
                transposed[v].push_back(u);
            }
        }

        // Step 3: Perform DFS on the transposed graph
        fill(visited.begin(), visited.end(), false); // Reset visited array
        while (!finishStack.empty()) {
            int v = finishStack.top();
            finishStack.pop();

            if (!visited[v]) {
                vector<int> component;
                dfsTranspose(v, visited, transposed, component);

                // Print the SCC
                cout << "SCC: ";
                for (int node : component) {
                    cout << node << " ";
                }
                cout << endl;
            }
        }
    }
};

int main() {
    // Example graph
    Graph g(5);
    g.addEdge(0, 2);
    g.addEdge(2, 1);
    g.addEdge(1, 0);
    g.addEdge(0, 3);
    g.addEdge(3, 4);

    cout << "Strongly Connected Components are:\n";
    g.findSCCs();

    return 0;
}
```

---

### How It Works:
1. **Graph Representation**:
   - The graph is represented using an adjacency list.
   
2. **DFS Traversal (Step 1)**:
   - Compute the order of "finishing times" of nodes and store them in a stack.
   
3. **Graph Transposition (Step 2)**:
   - Reverse all edges in the graph.

4. **DFS on Transposed Graph (Step 3)**:
   - Process nodes in the order of their finishing times from the stack.
   - Each DFS traversal gives a strongly connected component.

---

### Example Input:
Graph with edges:
```
0 → 2, 2 → 1, 1 → 0, 0 → 3, 3 → 4
```

### Output:
The SCCs are:
```
SCC: 4 
SCC: 3 
SCC: 0 1 2
```

---

Let me know if you need further explanations or modifications to the code! 🚀