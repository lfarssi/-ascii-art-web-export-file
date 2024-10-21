package ascii

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func TraitmentData(w http.ResponseWriter, banner string, arg string) string {
	fileName := "./banners/" + banner + ".txt"

	// Open the ASCII art file
	file, err := os.Open(fileName)
	if err != nil {
		if err := fetchAndSave(banner); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Data successfully written")
			return TraitmentData(w, banner, arg)
		}
	}
	defer file.Close()

	var asciiArt []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		errorHandler(w, nil, http.StatusInternalServerError)
		return ""
	}

	var result string
	lines := strings.Split(arg, "\n")
	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}
		// Iterate over each row of the ASCII art (0 to 7, for the 8 rows)
		for i := 1; i <= 8; i++ {
			for _, r := range line {
				// Ensure the character is within the valid ASCII range
				if r < 32 || r > 126 {
					errorHandler(w, nil, http.StatusInternalServerError)
					return ""
				}
				index := 9*(int(r)-32) + i
				result += asciiArt[index]
			}
			result += "\n" // Add newline after finishing the current row of the line
		}
	}
	return result
}

func BannerExists(banner string) bool {
	// Check if the provided banner exists (this is a placeholder function)
	// For example, compare the banner string to a list of supported banners
	supportedBanners := []string{"standard", "shadow", "thinkertoy"}
	for _, b := range supportedBanners {
		if banner == b {
			return true
		}
	}
	return false
}

type MyErr struct {
	StatusCode int
	Error      string
}

func errorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 | Internal Server Error !", http.StatusInternalServerError)
		return

	}
	typeError := MyErr{
		StatusCode: statusCode,
		Error:      http.StatusText(statusCode),
	}
	if err := t.Execute(w, typeError); err != nil {
		http.Error(w, "500 | Internal Server Error !", http.StatusInternalServerError)
		return
	}
}
