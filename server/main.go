package main

import (
	"encoding/json"
	"fmt"
	"grpc-with-goroutine/proto"
	"log"

	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedServerServer
	listUser []*proto.User
}

func main() {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	s := server{}
	s.LoadData(data)

	srv := grpc.NewServer()
	proto.RegisterServerServer(srv, &s)
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
	fmt.Println("Server started")

}

func (s *server) LoadData(data []byte) {
	if err := json.Unmarshal(data, &s.listUser); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}
}

func (s *server) GetAllFriends(request *proto.UserQuery, stream proto.Server_GetAllFriendsServer) error {
	id := request.GetId()
	c := make(chan *proto.User)
	n := 0
	for _, user := range s.listUser {
		if user.Id == id {
			n = len(user.Friends)
			for _, friendId := range user.Friends {
				go s.FindMyFriends(friendId, c)
			}
			break
		}
	}
	j := 0
	for v := range c {
		if err := stream.Send(v); err != nil {
			return err
		}
		if j == n-1 {
			close(c)
		}
		j++
	}
	return nil
}

func (s *server) FindMyFriends(userId string, c chan *proto.User) {
	// time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	for _, people := range s.listUser {
		if people.Id == userId {
			c <- people
			return
		}
	}
}

var data = []byte(`[
	  {
		"id": "618014731e4eb97a57129641",
		"roles": [
		  "disabled",
		  "trial"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Lancaster Travis",
		  "email": "HensleyLane@example.com",
		  "created_at": "2021-06-25T23:30:13.851Z"
		},
		"friends": [
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "618014736f77c474b08c0be2",
		"roles": [
		  "trial",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Ward Warren",
		  "email": "GretchenAlvarado@example.com",
		  "created_at": "2020-11-24T00:14:29.243Z"
		},
		"friends": []
	  },
	  {
		"id": "6180147327dc8706fd6d15fe",
		"roles": [
		  "disabled",
		  "premium"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Dena Dodson",
		  "email": "HoldenCole@example.com",
		  "created_at": "2020-07-29T13:27:51.766Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "61801473736dd3eef0a70ce6",
		"roles": [
		  "premium",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Dionne Gilmore",
		  "email": "TraciePalmer@example.com",
		  "created_at": "2020-10-23T05:12:45.796Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "618014738e6a8a83b7d1cd9b",
		"roles": [
		  "trial"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Merritt Warner",
		  "email": "TabathaMcdowell@example.com",
		  "created_at": "2020-10-26T07:07:03.776Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "6180147333c2288c44e12aed",
		"roles": [
		  "premium",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Burch Bullock",
		  "email": "IrisKim@example.com",
		  "created_at": "2020-12-09T13:22:45.350Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "6180147357294a20458a3179",
		"roles": [
		  "premium"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Janell Richmond",
		  "email": "CamposFreeman@example.com",
		  "created_at": "2021-07-11T03:15:42.355Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "6180147343d47bea1e6c9944",
		"roles": [
		  "disabled"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Jana Charles",
		  "email": "IdaCarrillo@example.com",
		  "created_at": "2020-08-09T02:02:58.389Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "61801473fd9e357c2ad8cd5f",
		"roles": [
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Kris Abbott",
		  "email": "CorineWilcox@example.com",
		  "created_at": "2020-01-20T11:59:23.856Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "618014730dfc762b769c395c",
		"roles": [],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Sawyer English",
		  "email": "GoffReese@example.com",
		  "created_at": "2020-06-14T10:01:58.117Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "618014731c8c66b3d51e3049",
		"roles": [],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Dejesus Elliott",
		  "email": "JacquelynMccarthy@example.com",
		  "created_at": "2020-05-27T13:48:36.399Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "618014730efbbaba484b0c70",
		"roles": [
		  "disabled",
		  "premium",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Atkinson Osborne",
		  "email": "AmieLester@example.com",
		  "created_at": "2021-01-01T04:48:49.767Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "618014735100328988e2dfd3",
		"roles": [
		  "disabled",
		  "trial",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Castro Atkins",
		  "email": "LaceyBecker@example.com",
		  "created_at": "2020-08-30T09:34:18.618Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "618014739efecf996dbfe458",
		"roles": [
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Stewart Espinoza",
		  "email": "MaynardCase@example.com",
		  "created_at": "2021-03-26T19:00:59.793Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "61801473b517382627d6584a",
		"roles": [
		  "premium",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Neal Gallagher",
		  "email": "NatashaSteele@example.com",
		  "created_at": "2021-04-23T02:38:22.184Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "61801473bac0b24d45264667",
		"roles": [
		  "trial",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Susan Sullivan",
		  "email": "MarylouAllen@example.com",
		  "created_at": "2021-07-31T21:20:04.770Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "618014738af6255414704d29",
		"roles": [
		  "trial",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Alyssa Duncan",
		  "email": "JanelleBush@example.com",
		  "created_at": "2020-10-15T06:03:08.635Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "6180147364d0c9e1c8a09d33",
		"roles": [
		  "premium"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Kelly Moran",
		  "email": "HamptonHolcomb@example.com",
		  "created_at": "2020-07-01T09:17:00.981Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29"
		]
	  },
	  {
		"id": "61801473888e5c5be29e5b1c",
		"roles": [
		  "disabled",
		  "premium",
		  "trial",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Morton Landry",
		  "email": "DianneErickson@example.com",
		  "created_at": "2021-03-05T05:21:05.699Z"
		},
		"friends": [
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "6180147326ac32645e27303e",
		"roles": [
		  "disabled",
		  "premium",
		  "trial",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Frazier Carter",
		  "email": "DeboraTerry@example.com",
		  "created_at": "2020-06-06T12:28:31.148Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "61801473c61f0381bddbce7c",
		"roles": [
		  "disabled",
		  "premium",
		  "trial",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Riggs Shields",
		  "email": "JaniceMurphy@example.com",
		  "created_at": "2021-06-08T12:42:01.506Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "6180147305795138d9f16b37",
		"roles": [
		  "disabled",
		  "premium",
		  "trial",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Barker Calhoun",
		  "email": "CareyDouglas@example.com",
		  "created_at": "2021-03-27T10:52:12.974Z"
		},
		"friends": []
	  },
	  {
		"id": "61801473e2736d6dd99fee1c",
		"roles": [
		  "disabled",
		  "premium",
		  "trial",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Jocelyn Tyler",
		  "email": "McfaddenWeaver@example.com",
		  "created_at": "2021-06-24T21:31:49.642Z"
		},
		"friends": []
	  },
	  {
		"id": "61801473fea6832491cbaba1",
		"roles": [
		  "disabled",
		  "premium",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Collier Blevins",
		  "email": "EttaBlanchard@example.com",
		  "created_at": "2020-04-29T17:16:18.415Z"
		},
		"friends": [
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  },
	  {
		"id": "61801473131f1faa223ce48b",
		"roles": [
		  "disabled",
		  "premium",
		  "suspended"
		],
		"avatar": {
		  "thumbnail": "https://placekitten.com/g/300/400",
		  "original": "https://placekitten.com/g/1200/1300"
		},
		"identity": {
		  "name": "Williams Mcintosh",
		  "email": "ToddGardner@example.com",
		  "created_at": "2020-01-13T21:25:04.666Z"
		},
		"friends": [
		  "618014731e4eb97a57129641",
		  "618014736f77c474b08c0be2",
		  "6180147327dc8706fd6d15fe",
		  "61801473736dd3eef0a70ce6",
		  "618014738e6a8a83b7d1cd9b",
		  "6180147333c2288c44e12aed",
		  "6180147357294a20458a3179",
		  "6180147343d47bea1e6c9944",
		  "61801473fd9e357c2ad8cd5f",
		  "618014730dfc762b769c395c",
		  "618014731c8c66b3d51e3049",
		  "618014730efbbaba484b0c70",
		  "618014735100328988e2dfd3",
		  "618014739efecf996dbfe458",
		  "61801473b517382627d6584a",
		  "61801473bac0b24d45264667",
		  "618014738af6255414704d29",
		  "6180147364d0c9e1c8a09d33"
		]
	  }
	]`)
