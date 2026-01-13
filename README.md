# proxyenvoy
docker compose up --build

check restapi to grpc:
1) curl http://localhost:8080/health
2) curl -X POST http://localhost:8080/auth/login   -H "Content-Type: application/json"   -d '{"email":"a","password":"b"}'
3) curl -X POST http://localhost:8080/auth/refresh   -H "Content-Type: application/json"  
4) curl -X POST http://localhost:8080/auth/refresh   -H "Content-Type: application/json"   -H "Authorization: Bearer <from 2 points>"

#1 check health
#2 authenticate 
#3 refresh token without Header Authorization (jwt missing)
#4 refresh token with Header Authorization (ok)
