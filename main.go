package main
import ("encoding/json";"fmt";"os";"strings")
func main() {
	if len(os.Args) < 2 { fmt.Fprintln(os.Stderr,"Usage: jsonpath <path>"); os.Exit(1) }
	path := strings.Split(strings.TrimPrefix(os.Args[1],"."), ".")
	var data any
	json.NewDecoder(os.Stdin).Decode(&data)
	current := data
	for _, key := range path {
		if m, ok := current.(map[string]any); ok {
			if v, exists := m[key]; exists { current = v } else { fmt.Fprintln(os.Stderr,"Key not found:", key); os.Exit(1) }
		} else { fmt.Fprintln(os.Stderr,"Not a map at:", key); os.Exit(1) }
	}
	b,_:=json.MarshalIndent(current, "", "  ")
	fmt.Println(string(b))
}
