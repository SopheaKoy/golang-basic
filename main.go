package main

import (
	"fmt"
	"log"
	"start/greetings"
	t "start/syntax"
)

// constant fixed values
const PI               = 3.14
const BASE_API_URL     = "https://api.escuelajs.co/api/v1/products"
const TELEGRAM_TOKEN   = ""
const TELEGRAM_CHAT_ID = ""
const FORMAT_DATETIME  = "20060102 15:04:05"

func main() {

    data := make([]string, 10)
    fmt.Print(data)

    s := make([]int, 3)
    s[0] = 12
    s[2] = 3
    
    fmt.Println("My name is Koy Sophea..!")
    t.L()
    t.Info()

    // varible declaration
    var a string
    var b int
    var c bool

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    // multiple varible 
    var one, two, three, four  int = 1, 2, 3, 4

    fmt.Print(one,   "\n")
    fmt.Print(two,   "\n")
    fmt.Print(three, "\n")
    fmt.Print(four,  "\n")

    fmt.Printf("values of one is %v and data type is %T", one, one)

    // var i   = 15.5
    // var txt = "Hello World!"

    // fmt.Printf("%v\n", i)
    // fmt.Printf("%#v\n", i)
    // fmt.Printf("%v%%\n", i)
    // fmt.Printf("%T\n", i)

    // fmt.Printf("%v\n", txt)
    // fmt.Printf("%#v\n", txt)
    // fmt.Printf("%T\n", txt)

    // var b1 bool = true // typed declaration with initial value
    // var b2 = true // untyped declaration with initial value
    // var b3 bool // typed declaration without initial value
    // b4 := true // untyped declaration with initial value

    // fmt.Println(b1) // Returns true
    // fmt.Println(b2) // Returns true
    // fmt.Println(b3) // Returns false
    // fmt.Println(b4) // Returns true


    // declaration in a block
    // var (
    //     productName  string
    //     productType  string
    //     productPrice float64
    //     productQTY   int
    // )

    is_disable := true
    fmt.Print(is_disable)
    is_open := false
    fmt.Print(is_open)

    // fmt.Print(productName)
    // fmt.Print(productType)
    // fmt.Print(productPrice)
    // fmt.Print(productQTY)

    log.SetPrefix("===== greetings: ")
    log.SetFlags(0)

    var name = "Sophea"
    age := 20          

    fmt.Println("Hello, world!")
    fmt.Printf("My name is %s and %d years old.\n", name, age)
    msg := greetings.Hello(name)
    fmt.Println(msg)

    // return with error
    msgOne, err:= greetings.Greeting(name)
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(msgOne)

    log.Println(name)
    
    fmt.Println("===================================")
    fmt.Println("=========== GOLANG ARRAY ==========")
    fmt.Println("===================================")
    
    // persons := []string{
    //     "Sophea",
    //     "Dann",
    //     "Dannra",
    // }
    
    // decal
    // var datas = []int{
    //     1,
    //     2,
    //     3,
    // }

    // fmt.Println(persons)
    // fmt.Println(datas)
    // fmt.Println("===================================")

    // num := 22
    // if num < 10 {
    //     fmt.Println("Good morning.")
    // } else if num < 20 {
    //     fmt.Println("Good day.")
    // } else {
    //     fmt.Println("Good evening.")
    // }

    // switch case signal
    // var day int = 4
    // switch day {
    //     case 1:
    //         fmt.Println("Monday")
    //     case 2:
    //         fmt.Println("Tuesday")
    //     case 3:
    //         fmt.Println("Wednesday")
    //     case 4:
    //         fmt.Println("Thursday")
    //     case 5:
    //         fmt.Println("Friday")
    //     case 6:
    //         fmt.Println("Saturday")
    //     case 7:
    //         fmt.Println("Sunday")
    //     default:
    //         fmt.Println("Not a weekday")
    // }

    // Multiple switch
    // The Multi-case switch Statement
    // It is possible to have multiple values for each case in the switch statement:

    // Syntax
    // switch expression {
    // case x,y:
    // // code block if expression is evaluated to x or y
    // case v,w:
    // // code block if expression is evaluated to v or w
    // case z:
    // ...
    // default:
    // // code block if expression is not found in any cases
    // }

    // Loop

    // Go for Loop
    // Loops are handy if you want to run the same code over and over again, each num with a different value.

    // Each execution of a loop is called an iteration.

    // The for loop can take up to three statements:

    // Syntax
    // for statement1; statement2; statement3 {
    //      code to be executed for each iteration
    // }

    // Golang with keyword continue
    // for i:=0; i < 5; i++ {
    //     if i == 3 {
    //         continue
    //     }
    //     fmt.Println(i)
    // }

    // using with rang
    // for i := range 10 {
    //     fmt.Println(i)
    // }

    // arr := []int{10, 20, 30}
    // for i, v := range arr {
    //     fmt.Printf("Index %d: Value %d\n", i+1, v)
    // }

    // who this book is for
    // now := time.Now().UTC()

    // fmt.Println(now.Format(FORMAT_DATETIME))
    
    // // 1. Fetch API
    // response, err := http.Get(BASE_API_URL)
    // if err != nil {
    //     fmt.Println("Request error:", err)
    //     return
    // }
    // defer response.Body.Close()

    // // 2. Read response body
    // body, err := io.ReadAll(response.Body)
    // if err != nil {
    //     fmt.Println("Read error:", err)
    //     return
    // }

    // // 3. Decode JSON into variable (interface{} for any JSON)
    // var data interface{}
    // if err := json.Unmarshal(body, &data); err != nil {
    //     fmt.Println("JSON unmarshal error:", err)
    //     return
    // }

    // // 4. Marshal JSON with indentation (pretty print)
    // pretty_json, err := json.MarshalIndent(data, "", "  ")
    // if err != nil {
    //     fmt.Println("JSON marshal error:", err)
    //     return
    // }

    // // 5. Print pretty JSON
    // fmt.Println(string(pretty_json[0]))
}
