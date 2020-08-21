# HOW TO RUN

1. Install golang
2. Download the vectory folder from [this link](https://drive.google.com/drive/folders/1bOFo7L9C8oLNLIdYi-CC3z9lG9hkEuip?usp=sharing), which contains the vector y construted from Planck data.
3. Build the program using 

```shell
go build -o cmbsolution cmd/main.go
```

4. Run the program, help message will be printed by passing -h flag

```shell
$ ./cmbsolution -h
Usage of ./cmbsolution:
  -afile string
        matrix A config file (default "config/a.csv")
  -lvl int
        level number (default 1)
  -m int
        value of m (default 9)
  -method string
        use which method,  you can choose iteration,std,and syl, and iteWoCorr(iteration method without residual correction) (default "syl")
  -n int
        value of n (default 4)
  -tfile string
        matrix T config file (default "config/t.csv")
  -th int
        the number of thread (default 100)
  -y string
        the vector y (default "config/vectory")
```

For example,

```shell
./cmbsolution -y vectory/vectory_10 -lvl 10 -method "syl"
```
