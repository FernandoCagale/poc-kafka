## Golang - Kafka [segmentio](https://github.com/segmentio/kafka-go)

## Go Mod

   * Download dependencies
     
```shell script
$ go mod download
```

   * Vendor local
   
```shell script
$ go mod vendor
```   

## Configurable environment variables LOCAL

   * Rename file `.envsample` to `.env`

```sh
$ cp .envsample .env
```

| Name                    	| Required                    	| Description                                                                                                                      	|
|-------------------------	|-----------------------------	|----------------------------------------------------------------------------------------------------------------------------------	|
| KAFKA_URL               	| TRUE                       	| Kafka URL                                                                                                                    	    |

## Example

   * Running `docker-compose up -d`   


   *  Producer

```sh   
$   go run producer/main.go --topic ecommerce --message info
```

   *  Consumer `group[order]`

```sh   
$   go run consumer/main.go --topic ecommerce --group order
$   message at topic:ecommerce partition:0 offset:2  = info
```

  *  Consumer `group[payment]`

```sh   
$   go run consumer/main.go --topic ecommerce --group payment
$   message at topic:ecommerce partition:0 offset:2  = info
```

   *  Consumer `group[notify]`

```sh   
$   go run consumer/main.go --topic ecommerce --group notify
$   message at topic:ecommerce partition:0 offset:2  = info
```
