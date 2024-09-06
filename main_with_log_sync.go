package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os/signal"
// 	"runtime/pprof"
// 	"syscall"

// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"os"
// 	"time"
// )

// var token string

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// start := time.Now()

// 	buildRequest()
// 	// Set a custom header (optional)
// 	w.Header().Set("Content-Type", "text/json")
// 	buildPraRequest()

// 	// Write a status code
// 	w.WriteHeader(http.StatusOK) // 200 OK

// 	buildPraResponse()

// 	// Write the response body
// 	buildResponse()
// 	fmt.Fprintf(w, "%s", "hello")

// 	// fmt.Fprintf(w, "%s", buildResponse())
// 	// log.Printf("Finish - %13v", time.Since(start))

// }

// func main() {
// 	token = randomToken(1000)
// 	cpuProf, err := os.Create("profile.cpu")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cpuProf.Close()
// 	if err := pprof.StartCPUProfile(cpuProf); err != nil {
// 		log.Fatalf("start cpu profile got error, %s", err)
// 	}
// 	defer pprof.StopCPUProfile()

// 	http.HandleFunc("/", handler)

// 	stop := make(chan os.Signal, 1)
// 	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

// 	println("start server")

// 	go func() {
// 		err := http.ListenAndServe(":8080", nil)

// 		if err != nil {
// 			println("err starting server: ", err)
// 		}
// 	}()

// 	<-stop
// 	fmt.Println("\nShutting down the server...")

// }

// func buildPraRequest() {
// 	// logger.WriteLog("build prarequest", logger.LogInfo, token)
// 	// logger.WriteLog("build prarequest", logger.LogInfo, token)
// 	// logger.WriteLog("build prarequest", logger.LogInfo, token)
// 	// logger.WriteLog("build prarequest", logger.LogInfo, token)
// 	// logger.WriteLog("build prarequest", logger.LogInfo, token)
// 	// logger.WriteLog("build prarequest", logger.LogInfo, token)
// 	// logger.WriteLog("build prarequest", logger.LogInfo, token)

// 	time.Sleep(20 * time.Millisecond)
// }

// func buildPraResponse() {
// 	// logger.WriteLog("build praresponse", logger.LogInfo, token)
// 	// logger.WriteLog("build praresponse", logger.LogInfo, token)
// 	// logger.WriteLog("build praresponse", logger.LogInfo, token)
// 	// logger.WriteLog("build praresponse", logger.LogInfo, token)
// 	// logger.WriteLog("build praresponse", logger.LogInfo, token)
// 	// logger.WriteLog("build praresponse", logger.LogInfo, token)
// 	// logger.WriteLog("build praresponse", logger.LogInfo, token)

// 	time.Sleep(20 * time.Millisecond)
// }

// func buildRequest() {
// 	httpHeader := createDynamicDummyRequestHeader()

// 	getHeaderKeysAsBytes(httpHeader)

// 	// headerBytes := getHeaderKeysAsBytes(httpHeader)
// 	// logger.WriteLog("build request", logger.LogInfo, string(headerBytes))

// 	time.Sleep(40 * time.Millisecond)
// }

// func buildResponse() string {
// 	httpResp := createDummyHTTPResponse()
// 	resp, err := getAllDataFromResponse(httpResp)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// _, err = json.Marshal(resp)
// 	respJson, err := json.Marshal(resp)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// logger.WriteLog("build response", logger.LogInfo, string(respJson))

// 	time.Sleep(50 * time.Millisecond)
// 	return string(respJson)
// }

// func createDynamicDummyRequestHeader() http.Header {
// 	header := http.Header{}

// 	// Generate some random values for dummy headers
// 	randomToken := fmt.Sprintf("Bearer dummy-token-%s", token)
// 	randomUserAgent := fmt.Sprintf("Go-Client-%d", 100)
// 	randomCustomValue := "100"

// 	header.Set("Content-Type", "application/json")
// 	header.Set("Authorization", randomToken)
// 	header.Set("User-Agent", randomUserAgent)
// 	header.Set("Custom-Header", randomCustomValue)
// 	return header
// }

// func getHeaderKeysAsBytes(headers http.Header) []byte {
// 	var buffer bytes.Buffer

// 	// Iterate over header keys and add them to the buffer
// 	for key := range headers {

// 		buffer.WriteString(fmt.Sprintf("%s - %s", key, headers.Get(key)))
// 		buffer.WriteString("\n") // Optional: separate keys by newlines
// 	}

// 	return buffer.Bytes()
// }

// // Function to create a dummy HTTP response with random values
// func createDummyHTTPResponse() *http.Response {
// 	// Generate random status code and body
// 	statusCodes := []int{200, 201, 400, 404, 500}
// 	randomStatusCode := statusCodes[rand.Intn(len(statusCodes))]

// 	randomBody := fmt.Sprintf("Response Body with Random Value: %s", token)

// 	// Create a dummy response
// 	response := &http.Response{
// 		Status:     http.StatusText(randomStatusCode),
// 		StatusCode: randomStatusCode,
// 		Header:     http.Header{},
// 		Body:       nopCloser{bytes.NewBufferString(randomBody)},
// 	}

// 	// Add random headers
// 	response.Header.Set("Content-Type", "application/json")
// 	response.Header.Set("X-Random-Token", fmt.Sprintf("dummy-token-%s", randomToken(1)))
// 	response.Header.Set("X-Random-Value", "sgagasgag")

// 	return response
// }

// // Helper to make the response body implement io.ReadCloser
// type nopCloser struct {
// 	*bytes.Buffer
// }

// func (nopCloser) Close() error { return nil }

// type Resp struct {
// 	StatusCode int    `json:"status_code"`
// 	Body       string `json:"data"`
// }

// // Function to get all data from the HTTP response
// func getAllDataFromResponse(resp *http.Response) (Resp, error) {
// 	// Read the response body
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return Resp{}, err
// 	}

// 	// Close the response body to avoid resource leaks
// 	defer resp.Body.Close()

// 	return Resp{
// 		StatusCode: resp.StatusCode,
// 		Body:       string(body),
// 	}, nil
// }

// func randomToken(length int) string {
// 	// Characters to choose from: letters and digits
// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 	token := make([]byte, length)
// 	for i := range token {
// 		token[i] = charset[rand.Intn(len(charset))]
// 	}

// 	return string(token)
// }
