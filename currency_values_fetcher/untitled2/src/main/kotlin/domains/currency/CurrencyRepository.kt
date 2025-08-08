package domains.currency

class CurrencyRepository(
    private val currencyGateway: CurrencyGateway, private val currencyDataAccessObject: CurrencyDataAccessObject
) {
    fun fetchCurrencies() {
        val currentCurrencies = currencyGateway.getCurrentCurrencies()

        currencyDataAccessObject.storeCurrentCurrencies(currentCurrencies)
    }
}
