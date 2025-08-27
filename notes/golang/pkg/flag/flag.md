# Go `flag` Package Functions

## **Defining Flags**

1. **`flag.String(name string, defaultValue string, usage string) *string`**  
   - Defines a `string` flag with a default value and help text. Returns a pointer.

2. **`flag.Int(name string, defaultValue int, usage string) *int`**  
   - Defines an `int` flag with a default value and help text. Returns a pointer.

3. **`flag.Bool(name string, defaultValue bool, usage string) *bool`**  
   - Defines a `bool` flag with a default value and help text. Returns a pointer.

4. **`flag.Float64(name string, defaultValue float64, usage string) *float64`**  
   - Defines a `float64` flag with a default value and help text. Returns a pointer.

5. **`flag.StringVar(p *string, name string, defaultValue string, usage string)`**  
   - Defines a `string` flag but stores the value in a given variable instead of returning a pointer.

6. **`flag.IntVar(p *int, name string, defaultValue int, usage string)`**  
   - Defines an `int` flag and stores the value in the provided variable.

7. **`flag.BoolVar(p *bool, name string, defaultValue bool, usage string)`**  
   - Defines a `bool` flag and stores the value in the provided variable.

8. **`flag.Float64Var(p *float64, name string, defaultValue float64, usage string)`**  
   - Defines a `float64` flag and stores the value in the provided variable.

---

## **Parsing and Handling Flags**

9. **`flag.Parse()`**  

   - Parses all defined flags from command-line arguments.

10. **`flag.Args() []string`**  

   - Returns remaining non-flag command-line arguments after parsing.

11. **`flag.NArg() int`**  
   - Returns the number of remaining non-flag arguments.

12. **`flag.NFlag() int`**  
   - Returns the number of flags that were explicitly set by the user.

---

## **Accessing Flags**
13. **`flag.Lookup(name string) *flag.Flag`**  
   - Returns a pointer to the flag with the given name, or `nil` if not found.

14. **`flag.Set(name string, value string) error`**  
   - Sets the value of a flag dynamically (useful for modifying flag values programmatically).

15. **`flag.Visited() bool`**  
   - Checks whether a flag has been explicitly set by the user.

---

## **Customizing Usage**
16. **`flag.Usage`**  
   - A user-defined function that replaces the default usage message.

17. **`flag.PrintDefaults()`**  
   - Prints all defined flags and their default values.

---

## **Advanced Handling**
18. **`flag.CommandLine`**  
   - The default flag set used by `flag` package functions.

19. **`flag.Var(value Value, name string, usage string)`**  
   - Defines a custom flag using a `flag.Value` interface for complex types.

20. **`flag.NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet`**  
   - Creates a new independent flag set (useful when handling multiple sets of flags).
