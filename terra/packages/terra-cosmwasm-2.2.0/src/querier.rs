use cosmwasm_std::{QuerierWrapper, StdResult};

use crate::query::{TaxCapResponse, TaxRateResponse, TerraQuery};

/// This is a helper wrapper to easily use our custom queries
pub struct TerraQuerier<'a> {
    querier: &'a QuerierWrapper<'a>,
}

impl<'a> TerraQuerier<'a> {
    pub fn new(querier: &'a QuerierWrapper<'a>) -> Self {
        TerraQuerier { querier }
    }

    pub fn query_tax_cap<T: Into<String>>(&self, denom: T) -> StdResult<TaxCapResponse> {
        let request = TerraQuery::TaxCap {
            denom: denom.into(),
        }
        .into();

        self.querier.custom_query(&request)
    }

    pub fn query_tax_rate(&self) -> StdResult<TaxRateResponse> {
        let request = TerraQuery::TaxRate {}.into();

        self.querier.custom_query(&request)
    }
}
