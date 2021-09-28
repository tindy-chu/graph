package cmd

import (
	"app/model"
	"app/util"
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

var possibleCmd = &cobra.Command{
	Use:   "possible",
	Short: "Print the possible paths between args",
	Args:  validateArgs,
	Run:   possible,
}

var shortestCmd = &cobra.Command{
	Use:   "shortest",
	Short: "Print the shortest paths between args",
	Args:  validateArgs,
	Run:   shortest,
}

var nodes = model.TNodes{}

func init() {
	nodes.Get()
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("Required start point and end point")
	}

	if len(args) > 2 {
		return errors.New("Too many args (expected 2 args)")
	}

	if start := nodes[args[0]]; len(start) == 0 {
		return errors.New("Invalid start point")
	}

	if end := nodes[args[1]]; len(end) == 0 {
		return errors.New("Invalid end point")
	}

	return nil
}

type tPath []string
type tPaths []string

func possible(cmd *cobra.Command, args []string) {
	paths := tPaths{}
	findAllPath(args[0], args[1], &tPath{args[0]}, &paths)

	resp := tResp{}
	resp["paths"] = paths
	resp["total"] = len(paths)

	response(&resp)
}

func shortest(cmd *cobra.Command, args []string) {
	paths := tPaths{}
	findAllPath(args[0], args[1], &tPath{args[0]}, &paths)

	var shortestPath string
	for _, current := range paths {
		if shortestPath == "" || len(shortestPath) >= len(current) {
			shortestPath = current
		}
	}

	resp := tResp{}
	resp["path"] = shortestPath

	response(&resp)
}

func findAllPath(currentKey string, endKey string, path *tPath, paths *tPaths) {
	for _, node := range nodes[currentKey] {
		currentPath := append(*path, node)

		if node == endKey {
			*paths = append(*paths, strings.Join(currentPath, ""))
			continue
		}

		if util.Includes(*path, node) {
			continue
		}

		findAllPath(node, endKey, &currentPath, paths)
	}
}
