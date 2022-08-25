var xml = new XMLHttpRequest();

// login
var data = JSON.stringify({
    query: `mutation{
      login(input: {
        user_email: "rikianfaisal@gmail.com"
        user_password: "12345"}) {
        user{
          user_id
          user_name
          user_image
          user_status
          created_date
          last_update
          products {
            UserId
            ProductId
            ProductName
            ProductPrice
          }
        }
        status
        message
      }
    }
  `
});

xml.open("POST", `http://localhost:8080/query`);

xml.setRequestHeader("Content-Type", "application/json")

xml.send(data)

xml.addEventListener("load", function() {
    document.querySelector(".user").innerHTML = this.responseText
    console.log(JSON.parse(this.responseText))
})
