# I tried to solve this without NetworkX, but I used it in the end
import networkx as nx

# Create a graph
G = nx.Graph()
with open('./Day25/input/25.txt', 'r') as file:
    lines = file.readlines()
    for line in lines:
        line = line.replace(':', '').strip()
        pieces = line.split(' ')
        key = pieces[0]
        values = pieces[1:]
        for value in values:
             G.add_edge(key, value)

# Shortest path from one node to another function
def shortest_path(graph, start, end):
    try:
        path = nx.shortest_path(graph, source=start, target=end)
        return path
    except nx.NetworkXNoPath:
        return "No path found"

# Using another BFS strategy, I determined these two nodes were likely
# on the opposite sides of the divide
start = 'qtv'
end = 'crc'

# Get the shortest path three times, then eliminate those edges
for i in range(3):
    result = shortest_path(G, start, end)
    for i in range(len(result) - 2):
        node1, node2 = result[i], result[i + 1]
        if G.has_edge(node1, node2):
            G.remove_edge(node1, node2)

# The nodes are now in two separate groups
            
# Count the number of reachable nodes (excluding unreachable nodes with infinite path length)
# From the start
shortest_paths = nx.shortest_path_length(G, source=start)
reachable_nodes1 = sum(1 for length in shortest_paths.values() if length != float('inf'))
# print(f"Number of reachable nodes from {start}: {reachable_nodes1}")

# From the end
shortest_paths = nx.shortest_path_length(G, source=end)
reachable_nodes2 = sum(1 for length in shortest_paths.values() if length != float('inf'))
# print(f"Number of reachable nodes from {end}: {reachable_nodes2}")

print(reachable_nodes1 * reachable_nodes2)