package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Response estructura del response
type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

//CreateDefaultResponse creamos la estructura que respondera por defecto
func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		writer:      w,
		contentType: "application/json",
	}
}

//SendNotFound envia el estado de notFound
func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

//NotFound escribe el estado y el mensaje del encabezado
func (res *Response) NotFound() {
	res.Status = http.StatusNotFound
	res.Message = "Resource not found"
}

//SendUnprocessableEntity envia el estado UnprocessableEntity
func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

//UnprocessableEntity escribe el estado y el mensaje del encabezado
func (res *Response) UnprocessableEntity() {
	res.Status = http.StatusUnprocessableEntity
	res.Message = "Unprocessable Entity"
}

//SendNoContent envia el estado NoContent
func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NoContent()
	response.Send()
}

//NoContent escribe el estado y el mensaje del encabezado
func (res *Response) NoContent() {
	res.Status = http.StatusNoContent
	res.Message = "No Content"
}

//SendData Envia los datos que se encuentra en "data" y los escribe en "w"
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

//Send envia el encabezado
func (res *Response) Send() {
	res.writer.Header().Set("Content-Type", res.contentType)
	output, _ := json.Marshal(&res)
	fmt.Fprintf(res.writer, string(output))
}
