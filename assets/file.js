window.addEventListener('load', () => {
    const req = new XMLHttpRequest()
    req.addEventListener("load", () => {
        alert(req.responseText)
    })
    req.open("GET", "http://localhost:8080/")
    req.send()
})