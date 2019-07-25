package greetings

import "time"

// Print good day
func GoodDay(name string) string {
    return "Good Day " + name
}

// Print good night
func GoodNight(name string) string {
    return "Good Night " + name
}

// Print current time, with RFC822 time format
func GetTime() string {
    localTime := time.Now()
    return localTime.Format(time.RFC822)
}

// If morning, return true
func IsAM() bool {
    localTime := time.Now()
    return localTime.Hour() <= 12
}

// If afternoon, return true
func IsAfternoon() bool {
    localTime := time.Now()
    return localTime.Hour() <= 18
}

// If evening, return true
func IsEvening() bool {
    localTime := time.Now()
    return localTime.Hour() <= 22
}
