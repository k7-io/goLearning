package main

import(
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name "`
	Age int `json:"age"`
	Sex bool `json:"sex"`
	Birthday string `json:"birthday"`
	Company string `json:"company, omitempty"`
	Language []string `json:"language"`
}

func main()  {
	//create map
	m := make(map[string]interface{}, 6)
	m["name"] = "xiao wang"
	m["age"] = 24
	m["sex"] = true
	m["birthday"] = "1995-01-01"
	m["company"] = "01 quick study"
	m["language"] = []string{"GO", "PHP", "Python"}
	//encode json
	result, _ := json.Marshal(m)
	resultFormat, _ := json.MarshalIndent(m, "", "	")
	fmt.Println("result:", string(result))
	fmt.Println("resultFormat:", string(resultFormat))

	person := Person{"xiao wang", 24, true, "1995-01-01", "01 quick study", []string{"GO", "Python"}}
	result1, err1 := json.MarshalIndent(person, "", "	")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("result1:", string(result1))

	//decode json
	jsonStr := `{
	"name": "xiaowang",
	"age": 24,
	"sex": true,
	"birthday": "1995-01-01",
	"company": "01 quick study",
	"language":["GO", "Python"
	]
	}`
	m1 := make(map[string]interface{}, 6)
	err2 := json.Unmarshal([]byte(jsonStr), &m1)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("m1=", m1)
	for key, value := range m1 {
		switch data := value.(type){
		case string:
			fmt.Printf("map[%s] value type is string, value = %s\n", key, data)
		case float64:
			fmt.Printf("map[%s] value type is float64, value = %f\n", key, data)
		case bool:
			fmt.Printf("map[%s] value type is bool, value = %t\n", key, data)
		case []string:
			fmt.Printf("map[%s] value type is []string, value = %v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s] value type is string, value = %v\n", key, data)
		}
	}

	//struct unmarshal
	var person1 Person
	err3 := json.Unmarshal([]byte(jsonStr), &person1)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Printf("person = %+v", person1)
}

