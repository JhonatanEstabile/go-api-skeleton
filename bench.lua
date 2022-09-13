wrk.path  = "/user"
wrk.method = "POST"
wrk.body   = '{"name":"test", "email":"test@test.com"}'
wrk.headers["Content-Type"] = "application/json"