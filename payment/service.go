package payment

import (
	"course-bwastartup-backend/user"
	"os"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(t Transaction, u user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(t Transaction, u user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("SERVER_KEY")
	midclient.ClientKey = os.Getenv("CLIENT_KEY")
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(t.Id),
			GrossAmt: int64(t.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: u.Name,
			Email: u.Email,
		},
	}

	// log.Println("GetToken:")
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, err
}
