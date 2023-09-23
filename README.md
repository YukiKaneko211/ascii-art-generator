Thank you for taking my audit. 

To test the code, you can run manually everything or use unit test code I prepared for some cases:

TO RUN MANUALLY:

```
go run . "string"
```

TO RUN UNIT TEST: 

```
go test
```

NB1: I found that sometimes the test gives the output looks wrong but just run the test again may resolce the problem. I don't know why. (Let me know if you have any clue!)
NB2: This unit test includes most of the provided audit cases (you can check them in ascii-art_test.go) but a few cannot be placed in it. So you have to additionally run the following ones to test everything:

```
go run . 'hello There 1 to 2!'
go run . "[\]^_ 'a"
go run . '\!" #$%&'"'"'()*+,-./'
```

And don't forget to test the last random cases! 

Any feedbacks and comments are welcome!