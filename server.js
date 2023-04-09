const express = require('express.js');
const path = require('path');

let initial_path = path.join(__dirname, "public");

let app = express();
app.use(express.static(initial_path));

app.get('/',(req,res)=>{
    res.sendFile(path.join(initial_path,"index.html"));
});

app.listen(3000,() => {
    console.log('listing on port 3000........');
});