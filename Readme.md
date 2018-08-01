# gcs-copy

gcs-copy is a simple binary that copies files from one location in a gcs bucket to another
It facilitates providing a public static version of your chartmuseum server via direct public gcs bucket access

## Installing

You will need to generate a service account JSON key for your cluster and store that as a secret in Kubernetes
Then you can provision gcs-copy as a cronjob with something like this:

    apiVersion: batch/v1beta1
    kind: CronJob
    metadata:
      name: gcs-copy
    spec:
      schedule: "*/1 * * * *"
      jobTemplate:
        spec:
          template:
            spec:
              containers:
              - name: gcs-copy
                image: i0nw/gcs-copy:0.0.1
                command: ["gcs-copy"]
                env:
                - name: BUCKET_NAME
                  value: "jenkinsx-chartmuseum"
                - name: COPY_FROM
                  value: "charts/index-cache.yaml"
                - name: COPY_TO
                  value: "index.yaml"
                - name: GOOGLE_APPLICATION_CREDENTIALS
                  value: "/gcs-service-key/i0ntest.key.json"
                volumeMounts:
                - name: gcs-service-key
                  mountPath: /gcs-service-key
                  readOnly: true
              volumes:
              - name: gcs-service-key
                secret:
                  secretName: gcs-jenkinsx-chartmuseum
              restartPolicy: OnFailure


You will need to adjust your chartmuseum server settings to prefix the "charts" directory and use gcs as storage.
Now whenever you hit the api to upload a chart simply run `helm repo update` afterwards to force a recreation of the index-cache.yaml 
You now have a highly available (serverless) public endpoint to serve your helm charts from! 
