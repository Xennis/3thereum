# contractfinder

Configuration

* Create a env file: `cp .env.template .env`
* Configure the variables in the created `.env` file
* Load the variables, e.g. `export $(cat .env | xargs)`


Run the programm

```shell
go run main.go
```

### Notes

Generate contract:

```shell
cd contracts/<name>
abigen --abi token.abi --pkg <name> --type Token --out token.go
```
