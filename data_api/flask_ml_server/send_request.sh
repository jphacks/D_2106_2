req=$(cat ex_req.txt)
res=$(eval curl -X POST -H "Content-Type: Application/json" -d \'$req\' http://localhost:8080/api/clustering)
echo $res > ex_res.txt