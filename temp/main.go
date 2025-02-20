package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
	"log"
)

// GPIO Pin definitions
const (
	triggerPin = 17 // GPIO pin for HC-SR04 trigger
	echoPin    = 27 // GPIO pin for HC-SR04 echo
	relayPin   = 22 // GPIO pin for relay control
)

// Threshold values (adjust as needed)
const (
	standingHeightThreshold = 80.0             // Example threshold in cm
	checkInterval           = 5 * time.Second  // Interval between height checks
	sittingTimeLimit        = 30 * time.Minute // Time before activating relay
)

func main() {
	// Initialize GPIO
	if err := rpio.Open(); err != nil {
		log.Fatalf("Failed to open GPIO: %v", err)
	}
	defer rpio.Close()

	// Configure pins
	trigger := rpio.Pin(triggerPin)
	echo := rpio.Pin(echoPin)
	relay := rpio.Pin(relayPin)

	trigger.Output()
	echo.Input()
	relay.Output()

	var sittingStartTime time.Time
	timerRunning := false

	fmt.Println("Standing Desk Monitor Started")

	for {
		distance := measureDistance(trigger, echo)

		fmt.Printf("Measured Distance: %.2f cm\n", distance)

		if distance < standingHeightThreshold {
			// Desk is at sitting height
			if !timerRunning {
				sittingStartTime = time.Now()
				timerRunning = true
				fmt.Println("Desk in sitting position. Timer started.")
			} else if time.Since(sittingStartTime) >= sittingTimeLimit {
				// Timer elapsed, activate relay
				fmt.Println("30 minutes elapsed. Moving desk to standing height.")
				activateRelay(relay)
				timerRunning = false
			}
		} else {
			// Desk is at standing height, reset timer
			if timerRunning {
				fmt.Println("Desk moved to standing height. Timer reset.")
				timerRunning = false
			}
		}

		time.Sleep(checkInterval)
	}
}

// measureDistance calculates distance using HC-SR04
func measureDistance(trigger, echo rpio.Pin) float64 {
	trigger.Low()
	time.Sleep(2 * time.Microsecond)
	trigger.High()
	time.Sleep(10 * time.Microsecond)
	trigger.Low()

	startTime := time.Now()
	for echo.Read() == rpio.Low {
		startTime = time.Now()
	}

	endTime := startTime
	for echo.Read() == rpio.High {
		endTime = time.Now()
	}

	duration := endTime.Sub(startTime).Seconds()
	distance := duration * 34300 / 2 // Speed of sound: 343 m/s

	return distance
}

// activateRelay triggers the relay to move the desk
func activateRelay(relay rpio.Pin) {
	relay.High()
	time.Sleep(2 * time.Second) // Adjust duration as needed
	relay.Low()
	fmt.Println("Desk moved to standing height.")
}
