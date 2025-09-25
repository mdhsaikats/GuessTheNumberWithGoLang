function sendData(){
    const num = Number(document.getElementById("guessInput").value); // <-- Convert to number
    console.log("value from the input:", num);

    fetch("http://localhost:8080/submit",{
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ num: num }) // <-- Send as number
    })
    .then(res => res.json())
    .then(data => {
        console.log("Response: ", data);
        document.getElementById("feedback").innerHTML = data.message || JSON.stringify(data);
    })
    .catch(err => console.error("Error", err));
}