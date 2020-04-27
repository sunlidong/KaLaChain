package main

import (
	"KaLaChain/alg"
	"KaLaChain/auth"
	"KaLaChain/bccsip"
	"KaLaChain/center"
	"KaLaChain/chaincode"
	"KaLaChain/data"
	"KaLaChain/fcw"
	"KaLaChain/gossip"
	"KaLaChain/kyc"
	"KaLaChain/msp"
	"KaLaChain/orderer"
	"KaLaChain/peer"
	"KaLaChain/pnf"
	"KaLaChain/privacy"
	"KaLaChain/token"
	"fmt"
	g "github.com/gin-gonic/gin"
)

func main() {
	r := g.Default()
	r.GET("/ping", func(c *g.Context) {
		c.JSON(200, g.H{
			"message": "pong",
		})
	})

	qie := []int{1, 2, 3, 4, 5}
	res := search(qie, 4)

	fmt.Println("res=>", res)
	//r.Run()
}

//
func server() {

	alg.Alg_main() // 排序算法
	auth.Auth_main()
	bccsip.Bsscip_main()
	center.Center_main()
	chaincode.Chaincode_main()
	data.Data_main()
	fcw.Fcw_main()
	gossip.Gossip_main()
	kyc.Kyc_main()
	msp.Msp_main()
	orderer.Orderer_main()
	peer.Peer_main()
	pnf.Pnf_main()
	privacy.Privacy_main()
	token.Token_main()

}

//

func search(nums []int, target int) int {
	msp_map := map[int]int{}

	for k, v := range nums {
		msp_map[v] = k
	}

	if msp_map[target] > 0 {
		return msp_map[target]
	} else {
		return int(-1)

	}

}
