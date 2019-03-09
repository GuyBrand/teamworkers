echo
echo \ "use internal get product"

curl -X POST \
  http://localhost:9001/local \
  -d '{"status":false}
'
./post1a.sh
