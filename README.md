> | `500`         | `application/json`                | `{"description":"Repository error","error":"Task not found"}`                           |# Task Manager API
A Go API made for studies that manages tasks with user authentication

## Routes
#### Users

<details>
 <summary><code>POST</code> <code><b>/register</b></code> <code>(registers new user)</code></summary>

##### JSON Body Params

> | name      |  type     | data type               |
> |-----------|-----------|-------------------------|
> | email     |  required | string (valid email)    |
> | password  |  required | string                  |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"ok":"use /login to authenticate"}`                                |
> | `400`         | `application/json`                | `{"description":"Validation error","error":"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`                            |
> | `400`         | `application/json`                | `{"description":"Validation error","error":"Key: 'UserDTO.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`                            |
> | `400`         | `application/json`                | `{"description":"Validation error","error":"Key: 'UserDTO.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`                            |
</details>

<details>
 <summary><code>POST</code> <code><b>/login</b></code> <code>(logs in user)</code></summary>

##### JSON Body Params

> | name      |  type     | data type               |
> |-----------|-----------|-------------------------|
> | email     |  required | string    |
> | password  |  required | string                  |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `{"token":"..."}`                                |
> | `400`         | `application/json`                | `{"description":"Validation error","error":"Key: 'UserDTO.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`                            |
> | `400`         | `application/json`                | `{"description":"Validation error","error":"Key: 'UserDTO.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`                           |
> | `400`         | `application/json`                | `{"description":"Validation error","error":"crypto/bcrypt: hashedPassword is not the hash of the given password"}`                           |
> | `500`         | `application/json`                | `{"description":"Repository error","error":"User not found"}`                           |
</details>

<details>
 <summary><code>PATCH</code> <code><b>/user</b></code> <code>(updates user's properties)</code></summary>

##### JSON Body Params

> | name      |  type     | data type               |
> |-----------|-----------|-------------------------|
> | email     |  optional | string    |
> | password  |  optional | string                  |

##### Headers

> | name | data type |
> |------|-----------|
> | Authorization | string (JWT token) |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `{"status":"success"}`                                |
> | `400`         | `application/json`                | `{"description":"Validation error","error":"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token is malformed: ..."}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"Authorization header not found"}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token has invalid claims: token is expired"}`                            |                        |                          |
</details>

<details>
 <summary><code>DELETE</code> <code><b>/user</b></code> <code>(deletes user)</code></summary>

##### Headers

> | name | data type |
> |------|-----------|
> | Authorization | string (JWT token) |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `{"status":"success"}`                                |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token is malformed: ..."}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"Authorization header not found"}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token has invalid claims: token is expired"}`                            |                        |                          |
</details>

#### Tasks

<details>
 <summary><code>GET</code> <code><b>/tasks</b></code> <code>(gets all authenticated user's tasks)</code></summary>

##### Headers

> | name | data type |
> |------|-----------|
> | Authorization | string (JWT token) |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `[{...}]`                                |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token is malformed: ..."}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"Authorization header not found"}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token has invalid claims: token is expired"}`                            |                        |                          |
</details>

<details>
 <summary><code>POST</code> <code><b>/tasks</b></code> <code>(creates a task in authenticated user's tasks list)</code></summary>

##### JSON Body Params

> | name      |  type     | data type               |
> |-----------|-----------|-------------------------|
> | title     |  required | string    |
> | description  |  optional | string                  |
> | toDate  |  required | string (dd/mm/yy hh:mm)                  |
> | tags  |  optional | []string                 |

##### Headers

> | name | data type |
> |------|-----------|
> | Authorization | string (JWT token) |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"createdTaskId":"..."}`              
> | `400`         | `application/json`                | `{"description":"Validation error","error":"Key: 'CreateTaskDTO.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`                            |                        |                          | 
> | `400`         | `application/json`                | `{"description":"Validation error","error":"Key: 'CreateTaskDTO.ToDate' Error:Field validation for 'ToDate' failed on the 'required' tag"}`                            |                        |                          | 
> | `400`         | `application/json`                | `{"description":"Validation error","error":"parsing time..."}`                            |                        |                          |   
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token is malformed: ..."}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"Authorization header not found"}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token has invalid claims: token is expired"}`                            |                        |                          |
</details>

<details>
 <summary><code>PATCH</code> <code><b>/tasks/:id</b></code> <code>(updates a task in authenticated user's tasks list)</code></summary>

##### JSON Body Params

> | name      |  type     | data type               |
> |-----------|-----------|-------------------------|
> | title     |  optional | string    |
> | description  |  optional | string                  |
> | toDate  |  optional | string (dd/mm/yy hh:mm)                  |
> | completed | optional | bool |
> | tags  |  optional | []string                 |

##### Headers

> | name | data type |
> |------|-----------|
> | Authorization | string (JWT token) |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `{"status":"success"}`              
> | `400`         | `application/json`                | `{"description":"Validation error","error":"parsing time..."}`                            |                        |                          |   
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token is malformed: ..."}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"Authorization header not found"}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token has invalid claims: token is expired"}`                            |                        |                          |
> | `500`         | `application/json`                | `{"description":"Repository error","error":"Task not found"}`                           |
</details>

<details>
 <summary><code>DELETE</code> <code><b>/tasks/:id</b></code> <code>(deletes a task in authenticated user's tasks list)</code></summary>

##### Headers

> | name | data type |
> |------|-----------|
> | Authorization | string (JWT token) |

##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`        | `{"status":"success"}`              
> | `400`         | `application/json`                | `{"description":"Validation error","error":"parsing time..."}`                            |                        |                          |   
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token is malformed: ..."}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"Authorization header not found"}`                            |                        |                          |
> | `401`         | `application/json`                | `{"description":"You must be authenticated","error":"token has invalid claims: token is expired"}`                            |                        |                          |
> | `500`         | `application/json`                | `{"description":"Repository error","error":"Task not found"}`                           |
</details>