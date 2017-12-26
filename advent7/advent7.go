package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	weight	int
	parent	string
}

var nodeMap map[string]Node

func parseOneNode(name string, weight int) {
	if _, ok := nodeMap[name]; !ok {
		var node Node
		node.weight = weight
		nodeMap[name] = node
		//fmt.Println("Adding", name, node.weight, "parent:", node.parent)
	}
}

func linkNode(name, parentNodeName string) {

	if node, ok := nodeMap[name]; ok {
		// already is in a map		
		node.parent = parentNodeName
		nodeMap[name] = node
		//fmt.Println("Linking existing", name, node.weight, "parent:", node.parent)
	} else {
		node.parent = parentNodeName
		nodeMap[name] = node
		//fmt.Println("Linking Added", name, "to", parentNodeName, node.weight, "parent:", node.parent)
	}

}

func findRoot() string {
	for k, v := range nodeMap {
		if v.parent == "" {
			//fmt.Println(k, v.parent)
			return k
		}
	}
	return ""
}

func processLine(line string) {
	split := strings.Split(line, " ")

	
	if len(split) == 2 {
		// if already is in map, update weight
		// else create
		name := split[0]
		weight, _ := strconv.Atoi(split[1][1:len(split[1])-1])

		parseOneNode(name, weight)

	} else {
		// if already is in map, update weight
		// else create
		parentNodeName := split[0]
		weight, _ := strconv.Atoi(split[1][1:len(split[1])-1])

		parseOneNode(parentNodeName, weight)

		for i := 3; i < len(split); i++ {
			// if already is in map, update weight
			// else create
			name := strings.Replace(split[i], ",", "", -1)
			linkNode(name, parentNodeName)
		}
	}

}

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	nodeMap = make(map[string]Node)
	for fscanner.Scan() {
		processLine(fscanner.Text())
	}

	//fmt.Println(nodeMap)
	fmt.Println("1st Answer:", findRoot())
}
