
#MiddleWare eg
```
    func BookValidationMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            var b Book
            json.NewDecoder(r.Body).Decode(&b)

            if b.IsEmpty() {
                http.Error(w, "Payload cannot be empty", http.StatusBadRequest)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
```
#Handlers
```
    func getBookHandler(w http.ResponseWriter, r *http.Request) {
        // 1. Get your data (even if it's your current fake data)
        var b Book = GetFakeData() 

        // 2. Run the check
        if b.IsEmpty() {
            // 3. If empty, send a 404 or 204 No Content status
            w.WriteHeader(http.StatusNotFound)
            w.Write([]byte("No book data found"))
            return 
        }

        // 4. If NOT empty, proceed to give the API response
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(b)
    }
```

#PRETTYFIRE
```
func prettify(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJSON.String(), nil
}

```
