# Job Viewer

This is a simple application to view job progress and results in a redis queue.

## Request

After having the application running, you can check results in a redis queue by making a GET reques:

```bash
$ curl <endpoint>/<redis address>/job:results/count
```

