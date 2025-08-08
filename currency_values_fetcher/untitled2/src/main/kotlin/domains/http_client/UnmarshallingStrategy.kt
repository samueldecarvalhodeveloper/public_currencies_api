package domains.http_client

interface UnmarshallingStrategy {
    fun <T> getUnmarshalledContent(marshalledContent: String): T
}