gcloud run deploy go-template --source . --update-env-vars "DSN=$(printenv DSN)" --region asia-northeast1
