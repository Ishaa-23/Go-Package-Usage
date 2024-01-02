package main

import (
    "context"
    "fmt"
    "os"
    EmployeeAPI "github.com/Ishaa-23/Go-Package"
)

func main() {
    id := int32(1) // int32 | 

    configuration := EmployeeAPI.NewConfiguration()
    apiClient := EmployeeAPI.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeAPI.ApiEmployeeIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAPI.ApiEmployeeIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEmployeeIdGet`: []Employee
    fmt.Fprintf(os.Stdout, "Response from `EmployeeAPI.ApiEmployeeIdGet`: %v\n", resp)
}