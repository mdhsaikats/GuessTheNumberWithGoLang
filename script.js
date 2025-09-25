function sendData(){
    const num = document.getElementById("guessInput").num
    console.log("value from the input:",num);

    fetch("http://localhost:7000/api/submit",{
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ guessInput: value})
    })
    .then(res => res.json())
    .then(data => {
        console.log("Response: ",data);
    })
    .catch(err => console.error("Error",err));
}