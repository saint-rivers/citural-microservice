package database

import "time"

func ContainerError(message string) map[string]string {
	return map[string]string{
		"message":   message,
		"timestamp": time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
}

func ContainerSuccess(containerId string, status string) map[string]string {
	return map[string]string{
		"containerId": containerId,
		"status":      status,
		"timestamp":   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
}
