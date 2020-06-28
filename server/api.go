// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package server

import (
	"context"
	"net/http"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// AccountAPIRouter defines the required methods for binding the api requests to a responses for the
// AccountAPI
// The AccountAPIRouter implementation should parse necessary information from the http request,
// pass the data to a AccountAPIServicer to perform the required actions, then write the service
// results to the http response.
type AccountAPIRouter interface {
	AccountBalance(http.ResponseWriter, *http.Request)
}

// BlockAPIRouter defines the required methods for binding the api requests to a responses for the
// BlockAPI
// The BlockAPIRouter implementation should parse necessary information from the http request,
// pass the data to a BlockAPIServicer to perform the required actions, then write the service
// results to the http response.
type BlockAPIRouter interface {
	Block(http.ResponseWriter, *http.Request)
	BlockTransaction(http.ResponseWriter, *http.Request)
}

// ConstructionAPIRouter defines the required methods for binding the api requests to a responses
// for the ConstructionAPI
// The ConstructionAPIRouter implementation should parse necessary information from the http
// request,
// pass the data to a ConstructionAPIServicer to perform the required actions, then write the
// service results to the http response.
type ConstructionAPIRouter interface {
	ConstructionCombine(http.ResponseWriter, *http.Request)
	ConstructionDerive(http.ResponseWriter, *http.Request)
	ConstructionHash(http.ResponseWriter, *http.Request)
	ConstructionMetadata(http.ResponseWriter, *http.Request)
	ConstructionParse(http.ResponseWriter, *http.Request)
	ConstructionPayloads(http.ResponseWriter, *http.Request)
	ConstructionPreprocess(http.ResponseWriter, *http.Request)
	ConstructionSubmit(http.ResponseWriter, *http.Request)
}

// MempoolAPIRouter defines the required methods for binding the api requests to a responses for the
// MempoolAPI
// The MempoolAPIRouter implementation should parse necessary information from the http request,
// pass the data to a MempoolAPIServicer to perform the required actions, then write the service
// results to the http response.
type MempoolAPIRouter interface {
	Mempool(http.ResponseWriter, *http.Request)
	MempoolTransaction(http.ResponseWriter, *http.Request)
}

// NetworkAPIRouter defines the required methods for binding the api requests to a responses for the
// NetworkAPI
// The NetworkAPIRouter implementation should parse necessary information from the http request,
// pass the data to a NetworkAPIServicer to perform the required actions, then write the service
// results to the http response.
type NetworkAPIRouter interface {
	NetworkList(http.ResponseWriter, *http.Request)
	NetworkOptions(http.ResponseWriter, *http.Request)
	NetworkStatus(http.ResponseWriter, *http.Request)
}

// AccountAPIServicer defines the api actions for the AccountAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AccountAPIServicer interface {
	AccountBalance(
		context.Context,
		*types.AccountBalanceRequest,
	) (*types.AccountBalanceResponse, *types.Error)
}

// BlockAPIServicer defines the api actions for the BlockAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type BlockAPIServicer interface {
	Block(context.Context, *types.BlockRequest) (*types.BlockResponse, *types.Error)
	BlockTransaction(
		context.Context,
		*types.BlockTransactionRequest,
	) (*types.BlockTransactionResponse, *types.Error)
}

// ConstructionAPIServicer defines the api actions for the ConstructionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ConstructionAPIServicer interface {
	ConstructionCombine(
		context.Context,
		*types.ConstructionCombineRequest,
	) (*types.ConstructionCombineResponse, *types.Error)
	ConstructionDerive(
		context.Context,
		*types.ConstructionDeriveRequest,
	) (*types.ConstructionDeriveResponse, *types.Error)
	ConstructionHash(
		context.Context,
		*types.ConstructionHashRequest,
	) (*types.ConstructionHashResponse, *types.Error)
	ConstructionMetadata(
		context.Context,
		*types.ConstructionMetadataRequest,
	) (*types.ConstructionMetadataResponse, *types.Error)
	ConstructionParse(
		context.Context,
		*types.ConstructionParseRequest,
	) (*types.ConstructionParseResponse, *types.Error)
	ConstructionPayloads(
		context.Context,
		*types.ConstructionPayloadsRequest,
	) (*types.ConstructionPayloadsResponse, *types.Error)
	ConstructionPreprocess(
		context.Context,
		*types.ConstructionPreprocessRequest,
	) (*types.ConstructionPreprocessResponse, *types.Error)
	ConstructionSubmit(
		context.Context,
		*types.ConstructionSubmitRequest,
	) (*types.ConstructionSubmitResponse, *types.Error)
}

// MempoolAPIServicer defines the api actions for the MempoolAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type MempoolAPIServicer interface {
	Mempool(context.Context, *types.NetworkRequest) (*types.MempoolResponse, *types.Error)
	MempoolTransaction(
		context.Context,
		*types.MempoolTransactionRequest,
	) (*types.MempoolTransactionResponse, *types.Error)
}

// NetworkAPIServicer defines the api actions for the NetworkAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type NetworkAPIServicer interface {
	NetworkList(context.Context, *types.MetadataRequest) (*types.NetworkListResponse, *types.Error)
	NetworkOptions(
		context.Context,
		*types.NetworkRequest,
	) (*types.NetworkOptionsResponse, *types.Error)
	NetworkStatus(
		context.Context,
		*types.NetworkRequest,
	) (*types.NetworkStatusResponse, *types.Error)
}