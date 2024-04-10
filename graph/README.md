# Terminology

## Node

A point or vertex on the graph.

## Edge

A connection between two nodes.

## Cycle

A cycle is a list of nodes, where you can walk along edges to in the order of
the nodes and end up at the first node. Cycles must contain more than 2 nodes.

A, B, C, A is an example of a cycle
(if there are edges between A - B, B - C, C - A).

## Acyclic

A graph that is acyclic contains no cycles.

## Connected

A node in a connected graph can "get to" any other node in the graph.

## Directed

When there is a direction to the connections.

## Undirected

When there isn't a direction to the connections.

# Weighted

The edges have a weight associated with them.

# DAG

Directed, acyclic graph

# Representing a Graph

There are two main ways of storing a graph

## Adjacendy List

Stored as a list of nodes, where each node is a list of connections

## Adjacendy Matrix

Stored as a matrix of connections, where each space in the matrix represents
a connection from node x, to node y (or vice versa). This is much slower,
because in this configuration, it is also storing _lacks_ of connections.
