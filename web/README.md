rgb(127, 76, 127) Not Found
rgb(95, 75, 115) No Content
rgb(165, 122, 122) Bad Request

{content: }
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
{content: }