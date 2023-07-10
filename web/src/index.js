(async ($) => {
  console.log("hello world")
  const response = await fetch("/status");
  const jsonData = await response.json();
  console.log(jsonData)
  const echoResponse = await fetch("/echo",
    {
      method: "POST",
      body: JSON.stringify({ message: "hello" }),
      headers: {
        "Content-Type": "application/json",
      }
    }
    );
  console.log(echoResponse)
})()
