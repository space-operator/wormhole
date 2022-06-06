#![allow(unused_variables)]

use near_contract_standards::fungible_token::metadata::{
    FungibleTokenMetadata, FungibleTokenMetadataProvider,
};

use near_contract_standards::fungible_token::FungibleToken;
use near_sdk::collections::LazyOption;
use near_sdk::json_types::U128;

use near_sdk::borsh::{self, BorshDeserialize, BorshSerialize};

use near_sdk::{
    env, near_bindgen, AccountId, Balance, PanicOnDefault, PromiseOrValue, StorageUsage,
};

const CHAIN_ID_NEAR: u16 = 15;

#[derive(BorshDeserialize, BorshSerialize, PanicOnDefault)]
pub struct FTContractMeta {
    metadata: FungibleTokenMetadata,
    vaa: Vec<u8>,
    sequence: u64,
}

#[near_bindgen]
#[derive(BorshDeserialize, BorshSerialize, PanicOnDefault)]
pub struct FTContract {
    token: FungibleToken,
    meta: LazyOption<FTContractMeta>,
    controller: AccountId,
    hash: Vec<u8>,
}

pub fn get_string_from_32(v: &[u8]) -> String {
    let s = String::from_utf8_lossy(v);
    s.chars().filter(|c| c != &'\0').collect()
}

#[near_bindgen]
impl FTContract {
    fn on_account_closed(&mut self, account_id: AccountId, balance: Balance) {
        env::log_str(&format!("Closed @{} with {}", account_id, balance));
    }

    fn on_tokens_burned(&mut self, account_id: AccountId, amount: Balance) {
        env::log_str(&format!("Account @{} burned {}", account_id, amount));
    }

    #[init]
    pub fn new(metadata: FungibleTokenMetadata, asset_meta: Vec<u8>, seq_number: u64) -> Self {
        assert!(!env::state_exists(), "Already initialized");

        metadata.assert_valid();

        let meta = FTContractMeta {
            metadata,
            vaa: asset_meta,
            sequence: seq_number,
        };

        let acct = env::current_account_id();
        let astr = acct.to_string();

        Self {
            token: FungibleToken::new(b"ft".to_vec()),
            meta: LazyOption::new(b"md".to_vec(), Some(&meta)),
            controller: env::predecessor_account_id(),
            hash: env::sha256(astr.as_bytes()),
        }
    }

    pub fn update_ft(
        &mut self,
        metadata: FungibleTokenMetadata,
        asset_meta: Vec<u8>,
        seq_number: u64,
    ) {
        if env::predecessor_account_id() != self.controller {
            env::panic_str("CrossContractInvalidCaller");
        }

        if seq_number <= self.meta.get().unwrap().sequence {
            env::panic_str("AssetMetaDataRollback");
        }

        let meta = FTContractMeta {
            metadata,
            vaa: asset_meta,
            sequence: seq_number,
        };

        self.meta.replace(&meta);
    }

    pub fn vaa_withdraw(
        &mut self,
        from: AccountId,
        amount: u128,
        receiver: String,
        chain: u16,
        fee: u128,
        payload: String,
    ) -> String {
        if env::predecessor_account_id() != self.controller {
            env::panic_str("CrossContractInvalidCaller");
        }

        let vaa = self.meta.get().unwrap().vaa;

        let mut p = [
            // PayloadID uint8 = 1
            (if payload.is_empty() { 1 } else { 3 } as u8)
                .to_be_bytes()
                .to_vec(),
            // Amount uint256
            vec![0; 24],
            (amount as u64).to_be_bytes().to_vec(),
            //TokenAddress bytes32
            vaa[0..32].to_vec(),
            // TokenChain uint16
            vaa[32..34].to_vec(),
            // To bytes32
            vec![0; (64 - receiver.len()) / 2],
            hex::decode(receiver).unwrap(),
            // ToChain uint16
            (chain as u16).to_be_bytes().to_vec(),
        ]
        .concat();

        if payload.is_empty() {
            p = [p, vec![0; 24], (fee as u64).to_be_bytes().to_vec()].concat();
            if p.len() != 133 {
                env::panic_str(&format!("paylod1 formatting errro  len = {}", p.len()));
            }
        } else {
            p = [p, hex::decode(&payload).unwrap()].concat();
            if p.len() != (133 + (payload.len() / 2)) {
                env::panic_str(&format!("paylod3 formatting errro  len = {}", p.len()));
            }
        }

        self.token.internal_withdraw(&from, amount);

        near_contract_standards::fungible_token::events::FtBurn {
            owner_id: &from,
            amount: &U128::from(amount),
            memo: Some("Wormhole burn"),
        }
        .emit();

        hex::encode(p)
    }

    pub fn vaa_transfer(
        &mut self,
        amount: u128,
        token_address: Vec<u8>,
        token_chain: u16,
        account_id: AccountId,
        recipient_chain: u16,
        fee: u128,
    ) {
        if env::predecessor_account_id() != self.controller {
            env::panic_str("CrossContractInvalidCaller");
        }

        if recipient_chain != CHAIN_ID_NEAR {
            env::panic_str("InvalidRecipientChain");
        }

        if amount == 0 {
            env::panic_str("ZeroAmountWastesGas");
        }

        self.token.internal_deposit(&account_id, amount - fee);

        near_contract_standards::fungible_token::events::FtMint {
            owner_id: &account_id,
            amount: &U128::from(amount - fee),
            memo: Some("wormhole minted tokens"),
        }
        .emit();

        if fee != 0 {
            self.token.internal_deposit(&env::signer_account_id(), fee);

            near_contract_standards::fungible_token::events::FtMint {
                owner_id: &env::signer_account_id(),
                amount: &U128::from(fee),
                memo: Some("wormhole minted tokens"),
            }
            .emit();
        }

        env::log_str("vaa_transfer called in ft");
    }

    pub fn account_storage_usage(&self) -> StorageUsage {
        self.token.account_storage_usage
    }

    /// Return true if the caller is either controller or self
    pub fn controller_or_self(&self) -> bool {
        let caller = env::predecessor_account_id();
        caller == self.controller || caller == env::current_account_id()
    }
}

near_contract_standards::impl_fungible_token_core!(FTContract, token, on_tokens_burned);
near_contract_standards::impl_fungible_token_storage!(FTContract, token, on_account_closed);

#[near_bindgen]
impl FungibleTokenMetadataProvider for FTContract {
    fn ft_metadata(&self) -> FungibleTokenMetadata {
        self.meta.get().unwrap().metadata
    }
}
