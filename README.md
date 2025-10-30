#  Bangor Burger CLI App

A simple Command-Line Interface (CLI) application written in Go (Golang) for simulating a burger shop ordering system.  
This project demonstrates fundamental Go concepts such as structs, interfaces, methods, error handling, concurrency, and modular design.


## Features

- Display menu list (with item ID, name, and price)  
- Add menu items to the cart (orders)  
- Checkout and generate invoice  
- Simulated payment system using concurrency (goroutines & WaitGroup)  
- View shopping history with detailed order breakdown  
- Error handling & input validation with graceful recovery  
- Clean modular code with reusable structs and interfaces 

## Technology Used
- [Go](https://go.dev/) — main programming languages
- [pgx](https://github.com/jackc/pgx) —   PostgreSQL driver and toolkit for Go
- [godotenv](https://github.com/joho/godotenv) — to read files `.env`