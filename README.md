# Cobra setup steps 
- go get -u github.com/spf13/cobra/cobra

### help
 go run cmd/main.go --help
 go run cmd/main.go echo -h
  go run cmd/main.go help echo
 go run cmd/main.go help echo times

golang-cobra-example $ go run cmd/main.go echo times -t 3 test test2
Echo: test test2
Echo: test test2
Echo: test test2




### Persistent (global) flags vs (local) flags
- persistent are applied to current command and all sub commands
- local only applied to current command

### Example
name of command will be echo
>
`echoCmd = &cobra.Command{
 		Use: "echo [strings to echo]",
 		Short: "prints given strings to stdout",
 		Args: cobra.MinimumNArgs(1),
 		Run: func(cmd *cobra.Command, args []string) {
 			fmt.Println("Echo: " + strings.Join(args, " "))
 		},
 	}`

	