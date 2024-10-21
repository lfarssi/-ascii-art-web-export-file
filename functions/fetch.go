package ascii

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// fetchAndSave fetches data from the API and saves it to a specified file.
func fetchAndSave(banner string) error {
	url := "https://raw.githubusercontent.com/01-edu/public/master/subjects/ascii-art/" + banner + ".txt"
	filePath := "./banners/" + banner + ".txt"

	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return err
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response status:", resp.StatusCode)
		return fmt.Errorf("received non-200 response status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// Create the directory if it doesn't exist
	err = os.MkdirAll("./banners", os.ModePerm)
	if err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	size, err := file.Write(body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded a file %s with size %d bytes\n", filePath, size)
	return nil
}
