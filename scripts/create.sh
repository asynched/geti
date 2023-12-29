curl http://localhost:3000/links \
  -H 'Content-Type: application/json' \
  -d '{"redirectUrl":"https://google.com","slug":"urubu_do_pix"}' | jq