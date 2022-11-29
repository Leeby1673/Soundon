const url = "http://127.0.0.1:8080"

var inputForm = document.getElementById("inputForm")

    inputForm.addEventListener("submit", (e)=>{
    e.preventDefault();
    const formData = new FormData(inputForm);
    var account = document.getElementById("floatingInput userAccount").value;
    var password = document.getElementById("floatingPassword userPassword").value;

    if (account == "" || password == "") {
        document.getElementById("prompt").style.display = "flex";
    } else {
        // Object.fromEntries() method 將鍵值組合變成 object物件 (類golang struct)
        
        // FormData.entries() method returns an iterator which iterates through all key/value pairs contained in the FormData. 
        // The key of each pair is a string object, and the value is either a string or a Blob.
        const formDataObj = Object.fromEntries(formData.entries());
        console.log(JSON.stringify(formDataObj))

        fetch(url,{
            method:"POST",
            body: JSON.stringify(formDataObj),
        }).then(
            response => response.json()
        ).then(
            data => {console.log(data)}
        ).catch(
            error => console.error(error)
        )
    }

})

document.getElementById("demogogo").onclick = testJS;
function testJS() {
    document.getElementById("demo").innerHTML = "success";
}