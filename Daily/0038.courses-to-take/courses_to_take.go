package main

import "fmt"

var mapOfSlicesOfStrings = map[string][]string{
	"CSC300": {"CSC100", "CSC200"},
	"CSC200": {"CSC100"},
	"CSC100": {},
}

func main() {
	fmt.Printf("For the courses %v the recommended order is:\n\t-> %v", mapOfSlicesOfStrings, courses_to_take(mapOfSlicesOfStrings))
}

// courses_to_take returns the recommended order to take the courses
// each course is represented by a string
// each course has a list of courses that are required to take it
// Approach:
// We have already been given a prerequisite graph adjacency list.
// For each vertex in the graph we do a depth first search following the outward edges.
// If we find a loop we immediately return nil otherwise we add each vertex to a stack.
// Time complexity: O(V + E) where V is the number of vertices and E is the number of edges
func courses_to_take(courses map[string][]string) []string {
	courseOrder := make([]string, 0, len(courses)) // the order of courses to take
	setVisiting := make(map[string]bool)           // keeps track of the vertices we have visited to avoid cycles
	setVisited := make(map[string]bool)            // the vertices/courses which requirements have already been checked

	var visit func(string) bool // the recursive function to visit the vertices
	visit = func(course string) bool {
		if _, ok := courses[course]; course != "" && !ok { // if the course is not in the graph
			return false
		}
		if setVisiting[course] { // loop detected
			return false
		}
		if setVisited[course] { // already verified
			return true
		}
		setVisiting[course] = true // add to loop checker
		for _, req := range courses[course] {
			if !visit(req) { // bubble up the "false" state in recursion
				return false
			}
		}
		setVisiting[course], setVisited[course] = false, true // passed loop check so remove it and mark as visited
		courseOrder = append(courseOrder, course)             // add to the course order
		return true
	}

	// loop through the courses in the given graph adjacency list
	for course := range courses {
		if !visit(course) {
			return nil
		}
	}

	return courseOrder
}
