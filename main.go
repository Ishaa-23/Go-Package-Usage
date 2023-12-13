package main

import (
    "context"
    "fmt"
   
    openapiclient "github.com/Ishaa-23/Go-SDK"
)

func main() {
    id := int32(2) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmployeeAPI.ApiEmployeeIdGet(context.Background(), id).Execute()
    if err != nil {

        fmt.Println(r)
    } else
	{
       // response from `ApiEmployeeIdGet`: []Employee
    fmt.Printf("Response from `EmployeeAPI.ApiEmployeeIdGet`: %v\n", resp)
	}
    
}