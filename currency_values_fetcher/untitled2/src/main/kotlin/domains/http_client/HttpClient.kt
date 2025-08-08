package domains.http_client

interface HttpClient {
    fun getGetRequestResponse(url: String, unmarshallingStrategy: UnmarshallingStrategy)

    fun executePostRequest(url: String, body: String)

    fun executePatchRequest(url: String, body: String)
}