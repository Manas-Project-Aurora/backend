An API backend for Project Aurora written in Go. 

The backend consists of three microservices:
- Admin panel
- Authenthication service
- Website for commoners
---
***Running the backend***
1. Clone the repo:
```bash
git clone git@github.com:Manas-Project-Aurora/gavna.git && cd gavna
```
2.  Choose which service you want to run on a particular server:
```bash
~ls
admin auth internal site LICENSE README.md
```
3. Suffer
4. Have a dbconfig.yaml prepared in desired service's directory for connecting to PostgresQL DB

Example yaml:
```yaml
  host: "localhost"
  port: 5432
  user: "your_user"
  password: "your_password"
  name: "your_database"
  sslmode: "disable"
```

**Running locally:**
1. Install go 1.22.2
2. Change directory to desired service's directory
```bash
cd gavna/service_name
```
3. Install dependencies:
```bash
go mod tidy && go mod download
```
4. Run the service:
```bash
go run cmd/main.go -p 8080 -d dbconfig.yaml
```
`-p` flag stands for the port on which the service will run. Default: `8080`
`-d` flag stands for the path to the yaml. Default: `dbconfig.yaml`

**Running in docker**
1. Build Docker container (from the root):
```bash
 sudo docker build -t myservice -f service_name/Dockerfile . --build-arg YAML=service_name/dbconfig.yaml
```
`--build-arg YAML` flag stands for the path to the yaml. Default: `service_name/dbconfig.yaml`
2. Run container:
```bash
sudo docker run -it --rm -e PORT=8080 -p 8080:8080 myservice
```
`-e PORT` flag stands for the port to expose inside the container. Default: `8080`
`-p` flag stands for the port to expose outside of the container. Syntax:
`-p OUTER:INNER`
3. Inside container:
```bash
./server -p 8080
```

---
***Contributions***
Don't
