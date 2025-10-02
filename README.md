# Concurrency Patterns in Golang

## Worker pool pattern in golang
The worker pool pattern as the name suggests contains a worker pool , and the tasks are filled in a channel which acts as a container. the workers pick the task and complete them . the process goes on till all tasks are not completed .

## Pipeline Concurrency Pattern in golang
The pipeline concurrency pattern works to manage concurrency in a situation where there are tasks taking place in stages . here we do not get to define no. of goroutines shooted at time of execution . 
