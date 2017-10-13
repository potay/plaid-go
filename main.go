package main

import (
	"fmt"
	"time"

	"github.com/potay/plaid-go/plaid"
)

// main contains example usage of all functions
func main() {
	// Create a Plaid client
	plaid := plaid.NewPlaid("59a4e1edbdc6a44914be2a8b", "95339295f3c3482a974d987a532703", "0f688d6dbe0ede437bf179fa694791", "sandbox", false, time.Duration(600)*time.Second)

	// GET /institutions
	res, err := plaid.Institutions.Get(50, 1)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(res["institutions"])

	// // GET /institutions/5301a93ac140de84910000e0
	// inst, err := plaid.GetInstitution(plaid.Sandbox, "5301a93ac140de84910000e0")
	// if err != nil {
	// 	fmt.Println("Institution Error:", err)
	// }
	// fmt.Println(inst.Name, "has mfa:", strings.Join(inst.MFA, ", "))

	// // GET /categories
	// categories, err := plaid.GetCategories(plaid.Sandbox)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("First category:", categories[0])

	// // GET /categories/13001001
	// category, err := plaid.GetCategory(plaid.Sandbox, "13001001")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("category", category.ID, "is", strings.Join(category.Hierarchy, ", "))

	// // POST /auth
	// postRes, mfaRes, err :=
	// 	client.AuthAddUser("user_good", "pass_good", "", "citi", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// } else if mfaRes != nil {
	// 	switch mfaRes.Type {
	// 	case "device":
	// 		fmt.Println("--Device MFA--")
	// 		fmt.Println("Message:", mfaRes.Device.Message)
	// 	case "list":
	// 		fmt.Println("--List MFA--")
	// 		fmt.Println("Mask:", mfaRes.List[0].Mask, "\nType:", mfaRes.List[0].Type)
	// 	case "questions":
	// 		fmt.Println("--Questions MFA--")
	// 		fmt.Println("Question:", mfaRes.Questions[0].Question)
	// 	case "selections":
	// 		fmt.Println("--Selections MFA--")
	// 		fmt.Println("Question:", mfaRes.Selections[1].Question)
	// 		fmt.Println("Answers:", mfaRes.Selections[1].Answers)
	// 	}

	// 	postRes2, mfaRes2, err := client.AuthStepSendMethod(mfaRes.AccessToken, "type", "email")
	// 	if err != nil {
	// 		fmt.Println("Error submitting send_method", err)
	// 	}
	// 	fmt.Printf("%+v\n", mfaRes2)
	// 	fmt.Printf("%+v\n", postRes2)

	// 	postRes2, mfaRes2, err = client.AuthStep(mfaRes.AccessToken, "tomato")
	// 	if err != nil {
	// 		fmt.Println("Error submitting mfa", err)
	// 	}
	// 	fmt.Printf("%+v\n", mfaRes2)
	// 	fmt.Printf("%+v\n", postRes2)
	// } else {
	// 	fmt.Println(postRes.Accounts)
	// 	fmt.Println("Auth Get")
	// 	fmt.Println(client.AuthGet("test_citi"))

	// 	fmt.Println("Auth DELETE")
	// 	fmt.Println(client.AuthDelete("test_citi"))
	// }

	// // POST /connect
	// postRes, mfaRes, err = client.ConnectAddUser("plaid_test", "plaid_good", "", "citi", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// } else if mfaRes != nil {
	// 	switch mfaRes.Type {
	// 	case "device":
	// 		fmt.Println("--Device MFA--")
	// 		fmt.Println("Message:", mfaRes.Device.Message)
	// 	case "list":
	// 		fmt.Println("--List MFA--")
	// 		fmt.Println("Mask:", mfaRes.List[0].Mask, "\nType:", mfaRes.List[0].Type)
	// 	case "questions":
	// 		fmt.Println("--Questions MFA--")
	// 		fmt.Println("Question:", mfaRes.Questions[0].Question)
	// 	case "selections":
	// 		fmt.Println("--Selections MFA--")
	// 		fmt.Println("Question:", mfaRes.Selections[1].Question)
	// 		fmt.Println("Answers:", mfaRes.Selections[1].Answers)
	// 	}

	// 	postRes2, mfaRes2, err := client.ConnectStepSendMethod(mfaRes.AccessToken, "type", "email")
	// 	if err != nil {
	// 		fmt.Println("Error submitting send_method", err)
	// 	}
	// 	fmt.Printf("%+v\n", mfaRes2)
	// 	fmt.Printf("%+v\n", postRes2)

	// 	postRes2, mfaRes2, err = client.ConnectStep(mfaRes.AccessToken, "1234")
	// 	if err != nil {
	// 		fmt.Println("Error submitting mfa", err)
	// 	}
	// 	fmt.Printf("%+v\n", mfaRes2)
	// 	fmt.Printf("%+v\n", postRes2)
	// } else {
	// 	fmt.Println(postRes.Accounts)
	// 	fmt.Println("Connect GET")
	// 	connectRes, _, _ := client.ConnectGet("test_citi", &plaid.ConnectGetOptions{true, "", "", ""})
	// 	fmt.Println(len(connectRes.Transactions))
	// 	fmt.Println(connectRes.Transactions)

	// 	fmt.Println("Connect DELETE")
	// 	fmt.Println(client.ConnectDelete("test_citi"))
	// }

	// // POST /balance
	// fmt.Println("Balance")
	// postRes, err = client.Balance("test_citi")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", postRes)

	// // POST /upgrade
	// fmt.Println("Upgrade")
	// postRes, mfaRes, err = client.Upgrade("test_bofa", "connect", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", mfaRes)
	// fmt.Printf("%+v\n", postRes)

	// // POST exchange_token
	// fmt.Println("ExchangeToken")
	// postRes, err = client.ExchangeToken("test,chase,connected")
	// fmt.Printf("%+v\n", postRes)
}
