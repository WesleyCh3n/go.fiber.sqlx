login:
	curl --data "user=wesley&pass=chen" http://localhost:3000/login
restrict:
	curl localhost:3000/test -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjUwMjUwMzQ1LCJuYW1lIjoid2VzbGV5In0.otuRRtVTWv1UrgPqOGzi9my99xTZ3HzOX_S3zAPqhL0"
