package main

type Xrip struct{}

const ReqUrl = "http://s1.ripple.com:51234/"

func (Xrip) NewWallet() {
	param := Xparam{
		Method: "wallet_propose",
		Params: []interface{}{
			WalletProposeParam{
				Passphrase: "snoPBrXGMeMyMHUVTgbuqAfg1SUTb",
			},
		},
	}
	req := new(Req)
	req.Post(ReqUrl, param)
}
