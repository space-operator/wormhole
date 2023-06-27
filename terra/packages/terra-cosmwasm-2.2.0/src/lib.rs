mod querier;
mod query;

pub use querier::TerraQuerier;
pub use query::{
    ContractInfoResponse, ExchangeRateItem, ExchangeRatesResponse, SwapResponse, TaxCapResponse,
    TaxRateResponse, TerraQuery,
};

// This export is added to all contracts that import this package, signifying that they require
// "terra" support on the chain they run on.
#[no_mangle]
extern "C" fn requires_terra() {}
