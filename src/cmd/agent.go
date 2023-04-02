/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type LifeCycle struct {
	Operating      bool
	UpTime         int
	DownTime       int
	Timeout        int
	Random         bool
	RandomUpTime   int
	RandomDownTime int
}

type State struct {
	Status string
	Cycle  LifeCycle
}

// We only need one state object globally
// Status will always eventually reconcile
// Random will always start defaulted to false
var globalState = State{
	Status: "down",
	Cycle: LifeCycle{
		Operating:      false,
		UpTime:         10,
		DownTime:       30,
		Timeout:        480,
		Random:         false,
		RandomUpTime:   10,
		RandomDownTime: 30,
	},
}

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Starts the agent REST server for remote interaction",
	Long:  `Starts the agent REST server - responsible for both the action of moving the desk up and down as well as randomizing the timing when required.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the agent server")
		router := gin.Default()
		router.GET("/state", getState)
		router.GET("/status", getStatus)
		router.PUT("/status", putStatus)
		router.GET("/lifecycle", getLifecycle)
		router.PUT("/lifecycle", putLifecycle)

		router.Run("localhost:8080")
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getState(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, globalState)
}

func getStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, globalState.Status)
}

func putStatus(c *gin.Context) {
	var newState State

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newState); err != nil {
		return
	}
	newDirection := strings.ToLower(newState.Status)
	if newDirection == "down" || newDirection == "up" {
		changeStatus(newDirection)
	} else {
		fmt.Println("Invalid state change request - only 'up' and 'down' permitted")
		c.IndentedJSON(http.StatusBadRequest, newState)
	}

	c.IndentedJSON(http.StatusNoContent, globalState.Status)
}

func getLifecycle(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, globalState.Cycle)
}

func putLifecycle(c *gin.Context) {
	var newCycle LifeCycle

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newCycle); err != nil {
		return
	}

	// Modify/setup lifecycle here
	changeLifecycle(newCycle)

	c.IndentedJSON(http.StatusNoContent, globalState.Cycle)
}

func changeStatus(direction string) {
	// this should be an idempotent operation - but we're going to check and log anyways
	// this may aid in the troubleshooting of incorrect or non-movement
	fmt.Printf("Attempting to move direction: %s - Current Direction: %s\n", direction, globalState.Status)

	// TODO: return an error?

	// execute some logic here to actuate the RPI relays

	// Update global state
	globalState.Status = direction
}

func changeLifecycle(newLC LifeCycle) {

	if globalState.Cycle.Operating {
		// If already operating - send data to the routine
		fmt.Println("Currently in operation")
		if !newLC.Operating {
			fmt.Println("Shutting down routine")
			// TODO: shut down the goroutine via channel?

			globalState.Cycle.Operating = false
		}
	} else {
		if newLC.Operating {
			fmt.Println("Starting a new routine")
			// TODO: Need to start new routine
			globalState.Cycle.Operating = true
		} else {
			// routine already running - do we need to do anything?
			fmt.Println("Routine already running")
		}
	}

	if newLC.Random && !globalState.Cycle.Random {
		fmt.Println("Transitioning to Random lifecycle")
		globalState.Cycle.Random = true
	}

	globalState.Cycle.UpTime = newLC.UpTime
	globalState.Cycle.DownTime = newLC.DownTime

}

func lifecycle() {
	// This will be the continuous loop for automating up/down intervals

}
