# EVERMOS Backend Task

This repository contains all code that is submitted as part of Evermos Backend
Assessment.

## 01. Store API

> Running Locally Requires:
> * Go 1.17+
> * MySQL
> * GORM
> * Gin Framework

This is the solution to **Question 1 of Evermos Backend Engineer Task

### Running Locally

1. Clone this repository.
2. Go to `task-store` folder.
3. Configure your database in line with `.env` file and run the migrations in `migrations` folder and     run `go run main.go migrate`
4. `go run main.go`



### Problems
1. Error caused by no validation in database to avoid inventory become negative
2. No notification of error that prevent costomer to order product that currently has no inventory.
3. Concurenncy, if there are multiple customer order at the same time, some of customer order might be not recorded in databases.

### Solution
1. Set validation minimum 0 stock in database.
2. use many layer checking & transactional on order.
3. Set sync in "order" function to make multiple order executed after another if coming at the same time..



This is the solution to **Question 2 of Evermos Backend Engineer Task

## 02. Treasure Hunt Puzzle
I did two function to generate probabilities of treasure and generate layout of treasure map.
To run locally, simply go into `task-treasure` directory and run `go run main.go`. The list of coordinates and map layout will be generated.

