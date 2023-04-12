package query

import (
	"api-pagamentos/logar"
	"api-pagamentos/models"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func InsertPagamento(client *dynamodb.Client, Pagamento models.Pagamento, log logar.Logfile) {
	pagamento, err := attributevalue.MarshalMap(Pagamento)
	logar.Check(err, log)

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Pagamentos"),
		Item:      pagamento,
	}

	_, err = client.PutItem(context.Background(), input)
	logar.Check(err, log)
}
