# ASCII Art Generator

## Description
This is a simple command-line tool written in Go. It takes a string input entered by the user then returns its graphical representation using the ASCII art characters from the specified banner file.

## Usage
1. **Clone Repository**: Clone this repository to your local machine.
2. **Install go**: Ensure you have Go installed on your machine.
3. **Run Program**: Execute the program by providing the flag, input string, and banner file as command-line arguments. 
    ```bash
    go run . --output=<filename.txt> something standard
    ```

    Example :

    ```bash
    go run . --output=<output.txt> hello standard
    ```
    This will return the graphical representation of the input string "hello". The results will be written to the specified output.txt file.
    ```bash
    cat -e output.txt
    ```

    ```bash
     _              _   _          
    | |            | | | |         
    | |__     ___  | | | |   ___   
    |  _ \   / _ \ | | | |  / _ \  
    | | | | |  __/ | | | | | (_) | 
    |_| |_|  \___| |_| |_|  \___/  
                                
                                                          
    ```   
Additionally, the program can run with a single string argument even when the user doesn't specify the flag.
```bash
    go run . hello 
```
```bash
     _              _   _          
    | |            | | | |         
    | |__     ___  | | | |   ___   
    |  _ \   / _ \ | | | |  / _ \  
    | | | | |  __/ | | | | | (_) | 
    |_| |_|  \___| |_| |_|  \___/  
                                
                                                          
```   
In this case, the result will be written to the default output file then automatically displayed on the terminal.

### Note
The main_test.go file tests the whole program while ascii_art_test.go tests the functions. To be able to run the main_test.go file successfully you have to first run the main.go file, with "hello" as your input string and "output.txt" as the value for the flag. This will help generate the output file which will be checked by the main_test.go file.

## Collaborators

* **[Tomlee Abila](https://github.com/Tomlee-abila/)**
* **[samuel omulo](https://github.com/somulo1/)**
* **[Rabin Otieno](https://github.com/Rabinnnn/)**