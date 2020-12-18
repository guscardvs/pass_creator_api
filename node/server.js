const express = require("express");
const { getPassPhrase } = require("./resources");

const app = express()
const port = 3000

app.use(express.json())

app.get("/", async (req, res) => {
    res.json(await getPassPhrase())
})

app.listen(port, ()=>{
    console.log("Express server running at :3000")
})