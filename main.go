package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Ishaa-23/TestSDK"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StudentAPI.ApiStudentGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StudentAPI.ApiStudentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStudentGet`: []Student
    fmt.Fprintf(os.Stdout, "Response from `StudentAPI.ApiStudentGet`: %v\n", resp)
}