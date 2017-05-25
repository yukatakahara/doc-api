## User stories
As an engineer I want to create a Admin so she can setup clinics.

```
go build
./cli list-quads
./cli list-admins
./cli list-clinics
./cli add-admin -email vera@gmail.com -password 112233 -name josh
./cli login-admin -email vera@gmail.com -password 112233
./cli add-clinic -name "great place" -address1 "4 Leng Kee Road" -jwt 111
./cli delete-clinic -id 111 -jwt 111
```
