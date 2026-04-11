
const express = require('express')
const app = express()
const port = 8000

app.use(express.json()); 
app.use(express.urlencoded({extended: true}));

app.get('/', (req, res) => {
  res.status(200).send("Welcome")
})


// async
// await

app.get("/get", async (req, res) => {
  try {
    const response = await fetch("https://api.github.com/users/devprajwalgotnochill");
    
    const data = await response.json();

    res.json(data); // sends JSON response
  } catch (err) {
    res.status(500).json({ error: "Failed to fetch commits" });
  }
});


app.post('/post', (req, res) => {

    let myJson = req.body; //json
	
	res.status(200).send(myJson);
})

// encoded url form
app.post('/postform', (req, res) => {

    res.status(200).send(JSON.stringify(req.body));
})
  

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})