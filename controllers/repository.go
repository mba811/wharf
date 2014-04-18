// Docker Push & Pull
// 执行 docker push 命令流程：
//     1. docker 向 registry 服务器注册 repository： PUT /v1/repositories/<username>/<repository> -> PUTRepository()
//     2. 参数是 JSON 格式的 <repository> 所有 image 的 id 列表，按照 image 的构建顺序排列。
//     3. 根据 <repository> 的 <tags> 进行循环：
//        3.1 获取 <image> 的 JSON 文件：GET /v1/images/<image_id>/json -> image.go#GETJSON()
//        3.2 如果没有此文件或内容返回 404 。
//        3.3 docker push 认为服务器没有 image 对应的文件，向服务器上传 image 相关文件。
//            3.3.1 写入 <image> 的 JSON 文件：PUT /v1/images/<image_id>/json -> image.go#PUTJSON()
//            3.3.2 写入 <image> 的 layer 文件：PUT /v1/images/<image_id>/layer -> image.go#PUTLayer()
//            3.3.3 写入 <image> 的 checksum 信息：PUT /v1/images/<image_id>/checksum -> image.go#PUTChecksum()
//        3.4 上传完此 tag 的所有 image 后，向服务器写入 tag 信息：PUT /v1/repositories/(namespace)/(repository)/tags/(tag) -> PUTTag()
//     4. 所有 tags 的 image 上传完成后，向服务器发送所有 images 的校验信息，PUT /v1/repositories/(namespace)/(repo_name)/images -> PUTRepositoryImages()
package controllers

import (
	"github.com/astaxie/beego"
)

type RepositoryController struct {
	beego.Controller
}

func (this *RepositoryController) Prepare() {
  //TODO Generate token & read endpoints from beego.AppConfig.String("Endpoints")
  this.Ctx.Output.Context.ResponseWriter.Header().Set("X-Docker-Endpoints", "")
  this.Ctx.Output.Context.ResponseWriter.Header().Set("WWW-Authenticate", "")
  this.Ctx.Output.Context.ResponseWriter.Header().Set("X-Docker-Token", "")
}

// http://docs.docker.io/en/latest/reference/api/registry_api/#tags
// GET /v1/repositories/(namespace)/(repository)/tags
// get all of the tags for the given repo.
// Example Request:
//    GET /v1/repositories/foo/bar/tags HTTP/1.1
//    Host: registry-1.docker.io
//    Accept: application/json
//    Content-Type: application/json
//    X-Docker-Registry-Version: 0.6.0
//    Cookie: (Cookie provided by the Registry)
// Parameters: 
//    namespace – namespace for the repo
//    repository – name for the repo
// Example Response:
//    HTTP/1.1 200
//    Vary: Accept
//    Content-Type: application/json
//    X-Docker-Registry-Version: 0.6.0
//    {
//      "latest": "9e89cc6f0bc3c38722009fe6857087b486531f9a779a0c17e3ed29dae8f12c4f",
//      "0.1.1":  "b486531f9a779a0c17e3ed29dae8f12c4f9e89cc6f0bc3c38722009fe6857087"
//    }
// Status Codes: 
//    200 – OK
//    401 – Requires authorization
//    404 – Repository not found
func (this *RepositoryController) GETTags() {

}

// http://docs.docker.io/en/latest/reference/api/registry_api/#tags
// GET /v1/repositories/(namespace)/(repository)/tags/(tag)
// get a tag for the given repo.
// Example Request:
//    GET /v1/repositories/foo/bar/tags/latest HTTP/1.1
//    Host: registry-1.docker.io
//    Accept: application/json
//    Content-Type: application/json
//    X-Docker-Registry-Version: 0.6.0
//    Cookie: (Cookie provided by the Registry)
// Parameters: 
//    namespace – namespace for the repo
//    repository – name for the repo
//    tag – name of tag you want to get
// Example Response:
//    HTTP/1.1 200
//    Vary: Accept
//    Content-Type: application/json
//    X-Docker-Registry-Version: 0.6.0
//    "9e89cc6f0bc3c38722009fe6857087b486531f9a779a0c17e3ed29dae8f12c4f"
// Status Codes: 
//    200 – OK
//    401 – Requires authorization
//    404 – Tag not found
func (this *RepositoryController) GETTag() {

}

// http://docs.docker.io/en/latest/reference/api/registry_api/#tags
// DELETE /v1/repositories/(namespace)/(repository)/tags/(tag)
// delete the tag for the repo
// Example Request:
//    DELETE /v1/repositories/foo/bar/tags/latest HTTP/1.1
//    Host: registry-1.docker.io
//    Accept: application/json
//    Content-Type: application/json
//    Cookie: (Cookie provided by the Registry)
// Parameters: 
//    namespace – namespace for the repo
//    repository – name for the repo
//    tag – name of tag you want to delete
// Example Response:
//    HTTP/1.1 200
//    Vary: Accept
//    Content-Type: application/json
//    X-Docker-Registry-Version: 0.6.0
//    ""
// Status Codes: 
//    200 – OK
//    401 – Requires authorization
//    404 – Tag not found
func (this *RepositoryController) DELETETag() {

}

// http://docs.docker.io/en/latest/reference/api/registry_api/#tags
// PUT /v1/repositories/(namespace)/(repository)/tags/(tag)
// put a tag for the given repo.
// Example Request:
//    PUT /v1/repositories/foo/bar/tags/latest HTTP/1.1
//    Host: registry-1.docker.io
//    Accept: application/json
//    Content-Type: application/json
//    Cookie: (Cookie provided by the Registry)
//    "9e89cc6f0bc3c38722009fe6857087b486531f9a779a0c17e3ed29dae8f12c4f"
// Parameters: 
//    namespace – namespace for the repo
//    repository – name for the repo
//    tag – name of tag you want to add
// Example Response:
//    HTTP/1.1 200
//    Vary: Accept
//    Content-Type: application/json
//    X-Docker-Registry-Version: 0.6.0
//    ""
// Status Codes: 
//    200 – OK
//    400 – Invalid data
//    401 – Requires authorization
//    404 – Image not found
func (this *RepositoryController) PUTTag() {

}

// http://docs.docker.io/en/latest/reference/api/registry_api/#repositories
// DELETE /v1/repositories/(namespace)/(repository)/
// delete a repository
// Example Request:
//    DELETE /v1/repositories/foo/bar/ HTTP/1.1
//    Host: registry-1.docker.io
//    Accept: application/json
//    Content-Type: application/json
//    Cookie: (Cookie provided by the Registry)
//    ""
// Parameters: 
//    namespace – namespace for the repo
//    repository – name for the repo
// Example Response:
//    HTTP/1.1 200
//    Vary: Accept
//    Content-Type: application/json
//    X-Docker-Registry-Version: 0.6.0
//    ""
// Status Codes: 
//    200 – OK
//    401 – Requires authorization
//    404 – Repository not found
func (this *RepositoryController) DELETERepositoryImages() {

}

// http://docs.docker.io/en/latest/reference/api/index_api/#repository
// Create a user repository with the given namespace and repo_name.
// Example Request: 
//    PUT /v1/repositories/foo/bar/ HTTP/1.1
//    Host: index.docker.io
//    Accept: application/json
//    Content-Type: application/json
//    Authorization: Basic akmklmasadalkm==
//    X-Docker-Token: true
//    [{"id": "9e89cc6f0bc3c38722009fe6857087b486531f9a779a0c17e3ed29dae8f12c4f"}]
// Parameters:
//    namespace – the namespace for the repo
//    repo_name – the name for the repo
// Example Response:
//    HTTP/1.1 200
//    Vary: Accept
//    Content-Type: application/json
//    WWW-Authenticate: Token signature=123abc,repository="foo/bar",access=write
//    X-Docker-Token: signature=123abc,repository="foo/bar",access=write
//    X-Docker-Endpoints: registry-1.docker.io [, registry-2.docker.io]
//    ""
// Status Codes:
//    200 – Created
//    400 – Errors (invalid json, missing or invalid fields, etc)
//    401 – Unauthorized
//    403 – Account is not Active
func (this *RepositoryController) PUTRepository() {

}

// http://docs.docker.io/en/latest/reference/api/index_api/#repository
// DELETE /v1/repositories/(namespace)/(repo_name)/
// Delete a user repository with the given namespace and repo_name.
// Example Request:
//    DELETE /v1/repositories/foo/bar/ HTTP/1.1
//    Host: index.docker.io
//    Accept: application/json
//    Content-Type: application/json
//    Authorization: Basic akmklmasadalkm==
//    X-Docker-Token: true
//    ""
// Parameters:
//    namespace – the namespace for the repo
//    repo_name – the name for the repo
// Example Response:
//    HTTP/1.1 202
//    Vary: Accept
//    Content-Type: application/json
//    WWW-Authenticate: Token signature=123abc,repository="foo/bar",access=delete
//    X-Docker-Token: signature=123abc,repository="foo/bar",access=delete
//    X-Docker-Endpoints: registry-1.docker.io [, registry-2.docker.io]
//    ""
// Status Codes:
//    200 – Deleted
//    202 – Accepted
//    400 – Errors (invalid json, missing or invalid fields, etc)
//    401 – Unauthorized
//    403 – Account is not Active
func (this *RepositoryController) DELETERepository() {

}

// http://docs.docker.io/reference/api/index_api/#repository-images
// PUT /v1/repositories/(namespace)/(repo_name)/images
// Update the images for a user repo.
// Example Request:
//    PUT /v1/repositories/foo/bar/images HTTP/1.1
//    Host: index.docker.io
//    Accept: application/json
//    Content-Type: application/json
//    Authorization: Basic akmklmasadalkm==
//    [
//      {
//        "id": "9e89cc6f0bc3c38722009fe6857087b486531f9a779a0c17e3ed29dae8f12c4f",
//        "checksum": "b486531f9a779a0c17e3ed29dae8f12c4f9e89cc6f0bc3c38722009fe6857087"
//      }
//    ]
// Parameters:
//    namespace – the namespace for the repo
//    repo_name – the name for the repo
// Example Response:
//    HTTP/1.1 204
//    Vary: Accept
//    Content-Type: application/json
//    ""
// Status Codes:
//    204 – Created
//    400 – Errors (invalid json, missing or invalid fields, etc)
//    401 – Unauthorized
//    403 – Account is not Active or permission denied
func (this *RepositoryController) PUTRepositoryImages() {

}

// http://docs.docker.io/reference/api/index_api/#repository-images
// GET /v1/repositories/(namespace)/(repo_name)/images
// get the images for a user repo.
// Example Request:
//    GET /v1/repositories/foo/bar/images HTTP/1.1
//    Host: index.docker.io
//    Accept: application/json
// Parameters:
//    namespace – the namespace for the repo
//    repo_name – the name for the repo
// Example Response:
//    HTTP/1.1 200
//    Vary: Accept
//    Content-Type: application/json
//     [{"id": "9e89cc6f0bc3c38722009fe6857087b486531f9a779a0c17e3ed29dae8f12c4f",
//     "checksum": "b486531f9a779a0c17e3ed29dae8f12c4f9e89cc6f0bc3c38722009fe6857087"},
//     {"id": "ertwetewtwe38722009fe6857087b486531f9a779a0c1dfddgfgsdgdsgds",
//     "checksum": "34t23f23fc17e3ed29dae8f12c4f9e89cc6f0bsdfgfsdgdsgdsgerwgew"}]
// Status Codes:
//    200 – OK
//    404 – Not found
func (this *RepositoryController) GETRepositoryImages() {

}

// http://docs.docker.io/reference/api/index_api/#repository-authorization
// PUT /v1/repositories/(namespace)/(repo_name)/auth
// authorize a token for a user repo
// Example Request:
//    PUT /v1/repositories/foo/bar/auth HTTP/1.1
//    Host: index.docker.io
//    Accept: application/json
//    Authorization: Token signature=123abc,repository="foo/bar",access=write
// Parameters:
//    namespace – the namespace for the repo
//    repo_name – the name for the repo
// Example Response:
//    HTTP/1.1 200
//    Vary: Accept
//    Content-Type: application/json
//    "OK"
// Status Codes:
//    200 – OK
//    403 – Permission denied
//    404 – Not found
func (this *RepositoryController) PUTRepositoryAuth() {

}

// Undocumented API
// PUT /v1/repositories/:username/:repository/properties
func (this *RepositoryController) PUTProperties() {

}

// Undocumented API
// GET /v1/repositories/:username/:repository/properties
func (this *RepositoryController) GETProperties() {

}

// Undocumented API
// GET /v1/repositories/:username/:repository/json
func (this *RepositoryController) GETRepositoryJSON() {

}

// Undocumented API
// GET /v1/repositories/:username/:repository/tags/:tag/json
func (this *RepositoryController) GETTagJSON() {

}

// Undocumented API
// DELETE /v1/repositories/:username/:repository/tags/:tag/json
func (this *RepositoryController) DELETERepositoryTags() {

}