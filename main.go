package main
import (
   "io/ioutil"
   "bytes"
   "log"
   "net/http"
   "encoding/json"
   "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  // MIDDLEWARE
  app.Use(func(c *fiber.Ctx) error {
      fmt.Println("First handler")
      return c.Next()
  })

  app.Get("/user/:name", func(c *fiber.Ctx) error {
      // c.Set("Connection", "keep-alive")
      // c.Set("Content-Type", "text/html; charset=utf-8")
      return c.SendString("Hello "+ c.Params("name"))
  })

  app.Get("/dump", func(c *fiber.Ctx) error {
      resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
      if err != nil {
          log.Fatalln(err)
      }
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
          log.Fatalln(err)
      }
      return c.SendString(string(body))
  })

  app.Get("/post", func(c *fiber.Ctx) error {
      //Encode the data
      postBody, _ := json.Marshal(map[string]string{
          "name":  "Toby",
          "email": "Toby@example.com",
      })
      responseBody := bytes.NewBuffer(postBody)
      //Leverage Go's HTTP Post function to make request
      resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
      //Handle Error
      if err != nil {
          log.Fatalf("An Error Occured %v", err)
      }
      defer resp.Body.Close()
      //Read the response body
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
          log.Fatalln(err)
      }
      return c.SendString(string(body))
  })

  app.Listen(":8080")
}
