package query

import (
	"api-pagamentos/logar"
	"api-pagamentos/models"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func InsertPagamento(client *dynamodb.Client, Pagamento models.Pagamento, log logar.Logfile) {
	Pagamento.Seq = ReturnSeq(client, log)

	pagamento, err := attributevalue.MarshalMap(Pagamento)
	logar.Check(err, log)

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Pagamentos"),
		Item:      pagamento,
	}
	fmt.Println(pagamento)
	_, err = client.PutItem(context.Background(), input)
	logar.Check(err, log)
}

func ReturnSeq(client *dynamodb.Client, log logar.Logfile) int {
	proj := expression.NamesList(expression.Name("Seq"))
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	logar.Check(err, log)

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("CaixaSeq"),
	}

	page, err := client.Scan(context.Background(), input)
	logar.Check(err, log)

	if len(page.Items) == 0 {
		return 0
	}

	var seq models.CaixaSeq
	err = attributevalue.UnmarshalMap(page.Items[0], &seq)
	logar.Check(err, log)

	return seq.Seq
}
