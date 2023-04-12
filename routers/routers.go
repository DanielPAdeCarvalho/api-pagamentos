package routers

import (
	"api-pagamentos/logar"
	"api-pagamentos/models"
	"api-pagamentos/query"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/beevik/ntp"
	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, log logar.Logfile) {
	c.IndentedJSON(http.StatusOK, "Servidor up")
}

// PostPagamento is a function that receives a POST request with a JSON body
func PostPagamento(c *gin.Context, log logar.Logfile, db *dynamodb.Client) {
	var pagamento models.Pagamento

	//o codigo esta indo no observatorio nacional pegar a data e hora
	datatemp, err := ntp.Time("gps.ntp.br")
	logar.Check(err, log)

	//Ajusta a hora para o horario de Fortaleza
	loc, err := time.LoadLocation("America/Fortaleza")
	logar.Check(err, log)
	pagamento.Data = datatemp.In(loc).Format("2006-01-02_15:04:05")

	err = c.BindJSON(&pagamento)
	logar.Check(err, log)
	// Bind JSON received to new Pagamento

	// Insert new Pagamento into database
	query.InsertPagamento(db, pagamento, log)
	c.IndentedJSON(http.StatusOK, "pagamento inserido com sucesso")
}
