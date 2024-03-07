// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Lookout v2 API",
    "version": "2.0.0"
  },
  "paths": {
    "/api/v1/jobGroups": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "groupJobs",
        "parameters": [
          {
            "name": "groupJobsRequest",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "filters",
                "order",
                "groupedField",
                "aggregates"
              ],
              "properties": {
                "activeJobSets": {
                  "description": "Only include jobs in active job sets",
                  "type": "boolean"
                },
                "aggregates": {
                  "description": "Additional fields to compute aggregates on",
                  "type": "array",
                  "items": {
                    "type": "string",
                    "x-nullable": false
                  },
                  "x-nullable": true
                },
                "filters": {
                  "description": "Filters to apply to jobs before grouping.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/filter"
                  },
                  "x-nullable": true
                },
                "groupedField": {
                  "type": "object",
                  "required": [
                    "field"
                  ],
                  "properties": {
                    "field": {
                      "description": "Field or annotation key to group by",
                      "type": "string",
                      "x-nullable": false
                    },
                    "isAnnotation": {
                      "type": "boolean",
                      "x-nullable": false
                    }
                  }
                },
                "order": {
                  "description": "Ordering to apply to job groups.",
                  "x-nullable": true,
                  "$ref": "#/definitions/order"
                },
                "skip": {
                  "description": "First elements to ignore from the full set of results. Used for pagination.",
                  "type": "integer"
                },
                "take": {
                  "description": "Number of job groups to fetch.",
                  "type": "integer"
                }
              }
            }
          },
          {
            "$ref": "#/parameters/backend"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns job groups from API",
            "schema": {
              "type": "object",
              "required": [
                "groups"
              ],
              "properties": {
                "groups": {
                  "description": "List of Job groups",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/group"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/v1/jobRunError": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "getJobRunError",
        "parameters": [
          {
            "name": "getJobRunErrorRequest",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "runId"
              ],
              "properties": {
                "runId": {
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Returns error for specific job run (if present)",
            "schema": {
              "type": "object",
              "properties": {
                "errorString": {
                  "description": "Error for individual job run",
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/v1/jobSpec": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "getJobSpec",
        "parameters": [
          {
            "name": "getJobSpecRequest",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "jobId"
              ],
              "properties": {
                "jobId": {
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Returns raw Job spec",
            "schema": {
              "type": "object",
              "properties": {
                "job": {
                  "description": "Job Spec object",
                  "type": "object",
                  "x-nullable": false
                }
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/v1/jobs": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "getJobs",
        "parameters": [
          {
            "name": "getJobsRequest",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "filters",
                "order"
              ],
              "properties": {
                "activeJobSets": {
                  "description": "Only include jobs in active job sets",
                  "type": "boolean"
                },
                "filters": {
                  "description": "Filters to apply to jobs.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/filter"
                  },
                  "x-nullable": true
                },
                "order": {
                  "description": "Ordering to apply to jobs.",
                  "x-nullable": true,
                  "$ref": "#/definitions/order"
                },
                "skip": {
                  "description": "First elements to ignore from the full set of results. Used for pagination.",
                  "type": "integer"
                },
                "take": {
                  "description": "Number of jobs to fetch.",
                  "type": "integer"
                }
              }
            }
          },
          {
            "$ref": "#/parameters/backend"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns jobs from API",
            "schema": {
              "type": "object",
              "properties": {
                "jobs": {
                  "description": "List of jobs found",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/job"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "filter": {
      "type": "object",
      "required": [
        "field",
        "value",
        "match"
      ],
      "properties": {
        "field": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "isAnnotation": {
          "type": "boolean",
          "x-nullable": false
        },
        "match": {
          "type": "string",
          "enum": [
            "exact",
            "anyOf",
            "startsWith",
            "contains",
            "greaterThan",
            "lessThan",
            "greaterThanOrEqualTo",
            "lessThanOrEqualTo",
            "exists"
          ],
          "x-nullable": false
        },
        "value": {
          "type": "object"
        }
      }
    },
    "group": {
      "type": "object",
      "required": [
        "name",
        "count",
        "aggregates"
      ],
      "properties": {
        "aggregates": {
          "type": "object",
          "additionalProperties": {
            "type": "object"
          },
          "x-nullable": false
        },
        "count": {
          "type": "integer",
          "x-nullable": false
        },
        "name": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "job": {
      "type": "object",
      "required": [
        "jobId",
        "queue",
        "owner",
        "jobSet",
        "cpu",
        "memory",
        "ephemeralStorage",
        "gpu",
        "priority",
        "submitted",
        "state",
        "lastTransitionTime",
        "duplicate",
        "annotations",
        "runs"
      ],
      "properties": {
        "annotations": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "x-nullable": false
        },
        "cancelReason": {
          "type": "string",
          "x-nullable": true
        },
        "cancelled": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "cpu": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "duplicate": {
          "type": "boolean",
          "x-nullable": false
        },
        "ephemeralStorage": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "gpu": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "jobId": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "jobSet": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "lastActiveRunId": {
          "type": "string",
          "x-nullable": true
        },
        "lastTransitionTime": {
          "type": "string",
          "format": "date-time",
          "minLength": 1,
          "x-nullable": false
        },
        "memory": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "namespace": {
          "type": "string",
          "x-nullable": true
        },
        "owner": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "priority": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "priorityClass": {
          "type": "string",
          "x-nullable": true
        },
        "queue": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "runs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/run"
          },
          "x-nullable": false
        },
        "state": {
          "type": "string",
          "enum": [
            "QUEUED",
            "PENDING",
            "RUNNING",
            "SUCCEEDED",
            "FAILED",
            "CANCELLED",
            "PREEMPTED",
            "LEASED"
          ],
          "x-nullable": false
        },
        "submitted": {
          "type": "string",
          "format": "date-time",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "order": {
      "type": "object",
      "required": [
        "field",
        "direction"
      ],
      "properties": {
        "direction": {
          "type": "string",
          "enum": [
            "ASC",
            "DESC"
          ],
          "x-nullable": false
        },
        "field": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "run": {
      "type": "object",
      "required": [
        "runId",
        "cluster",
        "jobRunState"
      ],
      "properties": {
        "cluster": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "exitCode": {
          "type": "integer",
          "format": "int32",
          "x-nullable": true
        },
        "finished": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "jobRunState": {
          "type": "string",
          "enum": [
            "RUN_PENDING",
            "RUN_RUNNING",
            "RUN_SUCCEEDED",
            "RUN_FAILED",
            "RUN_TERMINATED",
            "RUN_PREEMPTED",
            "RUN_UNABLE_TO_SCHEDULE",
            "RUN_LEASE_RETURNED",
            "RUN_LEASE_EXPIRED",
            "RUN_MAX_RUNS_EXCEEDED",
            "RUN_LEASED"
          ],
          "x-nullable": false
        },
        "leased": {
          "type": "string",
          "format": "date-time",
          "minLength": 1,
          "x-nullable": true
        },
        "node": {
          "type": "string",
          "x-nullable": true
        },
        "pending": {
          "type": "string",
          "format": "date-time",
          "minLength": 1,
          "x-nullable": true
        },
        "runId": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "started": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        }
      }
    }
  },
  "parameters": {
    "backend": {
      "enum": [
        "jsonb"
      ],
      "type": "string",
      "description": "The backend to use for this request.",
      "name": "backend",
      "in": "query"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Lookout v2 API",
    "version": "2.0.0"
  },
  "paths": {
    "/api/v1/jobGroups": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "groupJobs",
        "parameters": [
          {
            "name": "groupJobsRequest",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "filters",
                "order",
                "groupedField",
                "aggregates"
              ],
              "properties": {
                "activeJobSets": {
                  "description": "Only include jobs in active job sets",
                  "type": "boolean"
                },
                "aggregates": {
                  "description": "Additional fields to compute aggregates on",
                  "type": "array",
                  "items": {
                    "type": "string",
                    "x-nullable": false
                  },
                  "x-nullable": true
                },
                "filters": {
                  "description": "Filters to apply to jobs before grouping.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/filter"
                  },
                  "x-nullable": true
                },
                "groupedField": {
                  "type": "object",
                  "required": [
                    "field"
                  ],
                  "properties": {
                    "field": {
                      "description": "Field or annotation key to group by",
                      "type": "string",
                      "x-nullable": false
                    },
                    "isAnnotation": {
                      "type": "boolean",
                      "x-nullable": false
                    }
                  }
                },
                "order": {
                  "description": "Ordering to apply to job groups.",
                  "x-nullable": true,
                  "$ref": "#/definitions/order"
                },
                "skip": {
                  "description": "First elements to ignore from the full set of results. Used for pagination.",
                  "type": "integer"
                },
                "take": {
                  "description": "Number of job groups to fetch.",
                  "type": "integer"
                }
              }
            }
          },
          {
            "enum": [
              "jsonb"
            ],
            "type": "string",
            "description": "The backend to use for this request.",
            "name": "backend",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns job groups from API",
            "schema": {
              "type": "object",
              "required": [
                "groups"
              ],
              "properties": {
                "groups": {
                  "description": "List of Job groups",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/group"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/v1/jobRunError": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "getJobRunError",
        "parameters": [
          {
            "name": "getJobRunErrorRequest",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "runId"
              ],
              "properties": {
                "runId": {
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Returns error for specific job run (if present)",
            "schema": {
              "type": "object",
              "properties": {
                "errorString": {
                  "description": "Error for individual job run",
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/v1/jobSpec": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "getJobSpec",
        "parameters": [
          {
            "name": "getJobSpecRequest",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "jobId"
              ],
              "properties": {
                "jobId": {
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Returns raw Job spec",
            "schema": {
              "type": "object",
              "properties": {
                "job": {
                  "description": "Job Spec object",
                  "type": "object",
                  "x-nullable": false
                }
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/v1/jobs": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "getJobs",
        "parameters": [
          {
            "name": "getJobsRequest",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "filters",
                "order"
              ],
              "properties": {
                "activeJobSets": {
                  "description": "Only include jobs in active job sets",
                  "type": "boolean"
                },
                "filters": {
                  "description": "Filters to apply to jobs.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/filter"
                  },
                  "x-nullable": true
                },
                "order": {
                  "description": "Ordering to apply to jobs.",
                  "x-nullable": true,
                  "$ref": "#/definitions/order"
                },
                "skip": {
                  "description": "First elements to ignore from the full set of results. Used for pagination.",
                  "type": "integer"
                },
                "take": {
                  "description": "Number of jobs to fetch.",
                  "type": "integer"
                }
              }
            }
          },
          {
            "enum": [
              "jsonb"
            ],
            "type": "string",
            "description": "The backend to use for this request.",
            "name": "backend",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns jobs from API",
            "schema": {
              "type": "object",
              "properties": {
                "jobs": {
                  "description": "List of jobs found",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/job"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "GroupJobsParamsBodyGroupedField": {
      "type": "object",
      "required": [
        "field"
      ],
      "properties": {
        "field": {
          "description": "Field or annotation key to group by",
          "type": "string",
          "x-nullable": false
        },
        "isAnnotation": {
          "type": "boolean",
          "x-nullable": false
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "filter": {
      "type": "object",
      "required": [
        "field",
        "value",
        "match"
      ],
      "properties": {
        "field": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "isAnnotation": {
          "type": "boolean",
          "x-nullable": false
        },
        "match": {
          "type": "string",
          "enum": [
            "exact",
            "anyOf",
            "startsWith",
            "contains",
            "greaterThan",
            "lessThan",
            "greaterThanOrEqualTo",
            "lessThanOrEqualTo",
            "exists"
          ],
          "x-nullable": false
        },
        "value": {
          "type": "object"
        }
      }
    },
    "group": {
      "type": "object",
      "required": [
        "name",
        "count",
        "aggregates"
      ],
      "properties": {
        "aggregates": {
          "type": "object",
          "additionalProperties": {
            "type": "object"
          },
          "x-nullable": false
        },
        "count": {
          "type": "integer",
          "x-nullable": false
        },
        "name": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "job": {
      "type": "object",
      "required": [
        "jobId",
        "queue",
        "owner",
        "jobSet",
        "cpu",
        "memory",
        "ephemeralStorage",
        "gpu",
        "priority",
        "submitted",
        "state",
        "lastTransitionTime",
        "duplicate",
        "annotations",
        "runs"
      ],
      "properties": {
        "annotations": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "x-nullable": false
        },
        "cancelReason": {
          "type": "string",
          "x-nullable": true
        },
        "cancelled": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "cpu": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "duplicate": {
          "type": "boolean",
          "x-nullable": false
        },
        "ephemeralStorage": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "gpu": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "jobId": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "jobSet": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "lastActiveRunId": {
          "type": "string",
          "x-nullable": true
        },
        "lastTransitionTime": {
          "type": "string",
          "format": "date-time",
          "minLength": 1,
          "x-nullable": false
        },
        "memory": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "namespace": {
          "type": "string",
          "x-nullable": true
        },
        "owner": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "priority": {
          "type": "integer",
          "format": "int64",
          "x-nullable": false
        },
        "priorityClass": {
          "type": "string",
          "x-nullable": true
        },
        "queue": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "runs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/run"
          },
          "x-nullable": false
        },
        "state": {
          "type": "string",
          "enum": [
            "QUEUED",
            "PENDING",
            "RUNNING",
            "SUCCEEDED",
            "FAILED",
            "CANCELLED",
            "PREEMPTED",
            "LEASED"
          ],
          "x-nullable": false
        },
        "submitted": {
          "type": "string",
          "format": "date-time",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "order": {
      "type": "object",
      "required": [
        "field",
        "direction"
      ],
      "properties": {
        "direction": {
          "type": "string",
          "enum": [
            "ASC",
            "DESC"
          ],
          "x-nullable": false
        },
        "field": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "run": {
      "type": "object",
      "required": [
        "runId",
        "cluster",
        "jobRunState"
      ],
      "properties": {
        "cluster": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "exitCode": {
          "type": "integer",
          "format": "int32",
          "x-nullable": true
        },
        "finished": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "jobRunState": {
          "type": "string",
          "enum": [
            "RUN_PENDING",
            "RUN_RUNNING",
            "RUN_SUCCEEDED",
            "RUN_FAILED",
            "RUN_TERMINATED",
            "RUN_PREEMPTED",
            "RUN_UNABLE_TO_SCHEDULE",
            "RUN_LEASE_RETURNED",
            "RUN_LEASE_EXPIRED",
            "RUN_MAX_RUNS_EXCEEDED",
            "RUN_LEASED"
          ],
          "x-nullable": false
        },
        "leased": {
          "type": "string",
          "format": "date-time",
          "minLength": 1,
          "x-nullable": true
        },
        "node": {
          "type": "string",
          "x-nullable": true
        },
        "pending": {
          "type": "string",
          "format": "date-time",
          "minLength": 1,
          "x-nullable": true
        },
        "runId": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "started": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        }
      }
    }
  },
  "parameters": {
    "backend": {
      "enum": [
        "jsonb"
      ],
      "type": "string",
      "description": "The backend to use for this request.",
      "name": "backend",
      "in": "query"
    }
  }
}`))
}
