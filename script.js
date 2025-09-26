function sendData() {
    const input = document.getElementById("guessInput");
    const feedback = document.getElementById("feedback");
    const button = document.querySelector("button[onclick='sendData()']");
    const num = Number(input.value);

    // Validate input
    if (isNaN(num) || num < 0 || num > 99) {
        feedback.innerHTML = "Please enter a number between 0 and 99.";
        return;
    }

    button.disabled = true;

    fetch("http://localhost:8080/submit", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ num: num })
    })
    .then(res => res.json())
    .then(data => {
        feedback.innerHTML = `${data.message} (Target: ${data.target})`;
    })
    .catch(err => {
        feedback.innerHTML = "Server error. Please try again.";
        console.error("Error", err);
    })
    .finally(() => {
        input.value = "";
        button.disabled = false;
    });
}

document.getElementById("guessInput").addEventListener("keydown", function(event) {
    if (event.key === "Enter") {
        sendData();
    }
});