package controller

import (
	"context"
	"fatture75/model"
	"net/http"

	fattureincloudapi "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
)

type FattureInCloudApi struct {
	authContext context.Context
	client      *fattureincloudapi.APIClient
	companyId   int32
}

func NewFattureInCloudApi(token string, companyId int) *FattureInCloudApi {
	configuration := fattureincloudapi.NewConfiguration()

	fa := FattureInCloudApi{
		authContext: context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, token),
		client:      fattureincloudapi.NewAPIClient(configuration),
		companyId:   int32(companyId),
	}

	return &fa

}

func (fa *FattureInCloudApi) CreateDocument(
	doc *model.Document) (*http.Response, error) {

	createIssuedDocumentRequest := *fattureincloud.NewCreateIssuedDocumentRequest().SetData(doc.IssuedDocument)

	_, r, err := fa.client.IssuedDocumentsApi.CreateIssuedDocument(
		fa.authContext, fa.companyId,
	).CreateIssuedDocumentRequest(createIssuedDocumentRequest).Execute()

	return r, err

	//json.NewEncoder(os.Stdout).Encode(resp)
}
