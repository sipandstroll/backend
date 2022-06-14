# Sip & Stroll

If you want to test as authorized in Postman.
1. Build mobile app: either paste +40741198606 with 111 111 as code or your number and code sms received.
2. Copy JWT from android studio terminal. It is printed from the code for the moment
3. Add in postman Authorization as header and value: "Bearer ${JWT}"

Spor !

# DB Connection

Run the Cloud SQL Auth proxy to connect to DB. Learn more here: https://cloud.google.com/sql/docs/postgres/connect-instance-auth-proxy
```
./cloud_sql_proxy -instances=INSTANCE_CONNECTION_NAME=tcp:5432
```

# Deploy

Deployments are made automatically to production using the GitHub Actions pipeline when a new commit is made on the 'main' branch.
