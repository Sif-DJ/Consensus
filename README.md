# Consensus
Mandatory Activity 4
the ```Main.txt``` is an example of how it could look when runing the program
and the ``` runner5050, runner5051, runner5052``` are the individual clients output


to run this program open a terminal and navigate to the Consensus folder.
then run ```go run .\runner.go```
then accept the security popup that comes
the output will now be in the terminal you ran it from


to add additional client open the runner.go file and add more 
```go 
go client.StartNode("port1", "port2", false)
```
where port1 is the client port and port2 is the port of the client it can pass the token to
the token start in the client that has a true in boolean variable slot