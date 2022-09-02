package docs

var Yaml = `info:
  title: Docs
  version: 1.0.0
consumes:
  - application/json
produces:
  - application/json
swagger: '2.0'
paths:
  /api/v1/clusters/:
    get:
      operationId: ListClusters
      summary: List available clusters
      responses:
        200:
          description: OK
          schema:
            type: array
            items:
              $ref: '#/definitions/Cluster'
      security:
        - keystone: []
    post:
      operationId: CreateCluster
      summary: Create a cluster
      responses:
        200:
          description: OK
          schema:
            $ref: '#/definitions/Cluster'
      parameters:
        - name: body
          in: body
          required: true
          schema:
            $ref: '#/definitions/Cluster'
      security:
        - keystone: []

definitions:
  Cluster:
    type: object
    properties:
      name:
        description: name of the cluster
        type: string`

var Page = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta
            name="description"
            content="SwaggerUI"
    />
    <title>SwaggerUI</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@4.5.0/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@4.5.0/swagger-ui-bundle.js" crossorigin></script>
<script>
    window.onload = () => {
        window.ui = SwaggerUIBundle({
            url: './swagger.yaml',
            dom_id: '#swagger-ui',
        });
    };
</script>
</body>
</html>`
