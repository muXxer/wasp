// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use crate::*;

#[derive(Clone)]
pub struct ImmutableFinalizeAuctionParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableFinalizeAuctionParams {
    // NFT identifies the auction
    pub fn nft(&self) -> ScImmutableNftID {
		ScImmutableNftID::new(self.proxy.root(PARAM_NFT))
	}
}

#[derive(Clone)]
pub struct MutableFinalizeAuctionParams {
	pub(crate) proxy: Proxy,
}

impl MutableFinalizeAuctionParams {
    // NFT identifies the auction
    pub fn nft(&self) -> ScMutableNftID {
		ScMutableNftID::new(self.proxy.root(PARAM_NFT))
	}
}

#[derive(Clone)]
pub struct ImmutablePlaceBidParams {
	pub(crate) proxy: Proxy,
}

impl ImmutablePlaceBidParams {
    // NFT identifies the auction
    pub fn nft(&self) -> ScImmutableNftID {
		ScImmutableNftID::new(self.proxy.root(PARAM_NFT))
	}
}

#[derive(Clone)]
pub struct MutablePlaceBidParams {
	pub(crate) proxy: Proxy,
}

impl MutablePlaceBidParams {
    // NFT identifies the auction
    pub fn nft(&self) -> ScMutableNftID {
		ScMutableNftID::new(self.proxy.root(PARAM_NFT))
	}
}

#[derive(Clone)]
pub struct ImmutableSetOwnerMarginParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableSetOwnerMarginParams {
    // new SC owner margin in promilles
    pub fn owner_margin(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(PARAM_OWNER_MARGIN))
	}
}

#[derive(Clone)]
pub struct MutableSetOwnerMarginParams {
	pub(crate) proxy: Proxy,
}

impl MutableSetOwnerMarginParams {
    // new SC owner margin in promilles
    pub fn owner_margin(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(PARAM_OWNER_MARGIN))
	}
}

#[derive(Clone)]
pub struct ImmutableStartAuctionParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableStartAuctionParams {
    // description of the NFTs being auctioned
    pub fn description(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_DESCRIPTION))
	}

    // duration of auction in minutes
    pub fn duration(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_DURATION))
	}

    // minimum required amount for any bid
    pub fn minimum_bid(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(PARAM_MINIMUM_BID))
	}

    // NFT of the NFTs being auctioned
    pub fn nft(&self) -> ScImmutableNftID {
		ScImmutableNftID::new(self.proxy.root(PARAM_NFT))
	}
}

#[derive(Clone)]
pub struct MutableStartAuctionParams {
	pub(crate) proxy: Proxy,
}

impl MutableStartAuctionParams {
    // description of the NFTs being auctioned
    pub fn description(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_DESCRIPTION))
	}

    // duration of auction in minutes
    pub fn duration(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_DURATION))
	}

    // minimum required amount for any bid
    pub fn minimum_bid(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(PARAM_MINIMUM_BID))
	}

    // NFT of the NFTs being auctioned
    pub fn nft(&self) -> ScMutableNftID {
		ScMutableNftID::new(self.proxy.root(PARAM_NFT))
	}
}

#[derive(Clone)]
pub struct ImmutableGetAuctionInfoParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetAuctionInfoParams {
    // NFT identifies the auction
    pub fn nft(&self) -> ScImmutableNftID {
		ScImmutableNftID::new(self.proxy.root(PARAM_NFT))
	}
}

#[derive(Clone)]
pub struct MutableGetAuctionInfoParams {
	pub(crate) proxy: Proxy,
}

impl MutableGetAuctionInfoParams {
    // NFT identifies the auction
    pub fn nft(&self) -> ScMutableNftID {
		ScMutableNftID::new(self.proxy.root(PARAM_NFT))
	}
}
