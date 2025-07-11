# Ascii-Art Generator

Simple Go program to generate ascii-art from the input text.

## Used Technologies

- Go (libraries below):
    - fmt
	- os
	- strings
	- testing

## Installation & How to Use

1. Clone repository
2. Run `ascii-art.go` with the string you want to convert to ascii-art.

    Ex:
    ```
    go run ascii-art.go "hello There 1 to 2!"
     _              _   _                 _______   _                                            _                           _  
    | |            | | | |               |__   __| | |                                 _        | |                  ____   | | 
    | |__     ___  | | | |   ___            | |    | |__     ___   _ __    ___        / |       | |_    ___         |___ \  | | 
    |  _ \   / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \       | |       | __|  / _ \          __) | | | 
    | | | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/       | |       \ |_  | (_) |        / __/  |_| 
    |_| |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|       |_|        \__|  \___/        |_____| (_) 
    ```
    If the art looks broken, make sure your terminal width has enough space to print the art.

## Spec Detail

This program convert and print the input string into letter ascii-art formatted in the `standard.txt`.