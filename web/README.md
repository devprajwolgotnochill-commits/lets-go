
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
# The root folder should contain only one main.go file 
# other file should be sotred in seprate folders
#IMP NOTE ABT HTTP
```
GET — Requests the server to send back the resource specified by the URI.

HEAD — Works like GET, but the server only returns the response headers and
not the actual body content. This is useful for checking metadata without
downloading the full resource.

POST — Sends data in the request body to the server, which then processes
it according to the target resource.

PUT — Uploads data to the specified URI and stores it as the resource there.
If the resource already exists, it is replaced; otherwise, a new one is created.

DELETE — Requests the server to delete the resource located at the given URI.

TRACE — Asks the server to send back the same request it received, allowing the client to see any changes made by intermediate servers.

OPTIONS — Requests the server to provide the list of HTTP methods that are supported for the target resource.

CONNECT — Instructs the server to establish a direct connection with the client, commonly used for creating secure HTTPS tunnels.
PATCH — Sends partial updates to modify the existing resource identified by the URI.
```
