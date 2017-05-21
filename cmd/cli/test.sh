#!/bin/sh

go run main.go add-admin -email vera@gmail.com -password 112233 -name josh
# go run main.go login-admin -email vera@gmail.com -password 112233
# go run main.go add-clinic -name "great place" -address1 "4 Leng Kee Road" -jwt

go run main.go list-quads
go run main.go list-admins
go run main.go list-clinics
