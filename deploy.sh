gcloud run deploy go-template --source . \
    --set-env-vars "PSQL_DBNAME=$(printenv PSQL_DBNAME)" \
    --set-env-vars "PSQL_HOST=$(printenv PSQL_HOST)" \
    --set-env-vars "PSQL_USER=$(printenv PSQL_USER)" \
    --set-env-vars "PSQL_PASS=$(printenv PSQL_PASS)" \
    --set-env-vars "PSQL_SSLMODE=$(printenv PSQL_SSLMODE)" \
    --region asia-northeast1
