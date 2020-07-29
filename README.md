#consumer-order-prediction

##To run this project, follow the steps below :
1. Clone the repository
2. Change GOPATH to pwd to this repo i.e. consumer-order-prediction or just clone the repo in usr/local/go/src
3. To install GIN and gPRC run the following commands
	i. go get -u "github.com/gin-gonic/gin"
	ii. go get -u "google.golang.org/grpc"
4. In terminal go to consumer-order-prediction/pkg/grpc and run “go run server.go”
5. In another terminal go to cmd/gin_grpc and run “go run main.go”
6. After this the API will run on localhost:5656

##To Fetch Different queries try :
1. “Localhost:5656/” for homepage
2. “Localhost:5656/popularrestaurant” to get popular restaurant
3. “Localhost:5656/order” to place a new order
4. “Localhost:5656/popularcuisine”  to check the most popular cuisine
