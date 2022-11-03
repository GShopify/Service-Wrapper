## Service Wrapper

Base wrapper for federated services

### New service initializing

```shell
./scripts/init_service.sh --path=../wakanda
```

Define schema for the service in `graphql/schema.graphql`

#### Generate models
```shell
cd graphql
go run github.com/99designs/gqlgen
```

Now we are ready to implement `graphql/generated/resolver.go` and start server `cmd/main.go`.

### Storefront API

#### Cart
- [ ] `query` cart
- [ ] `mutation` cartAttributesUpdate
- [ ] `mutation` cartBuyerIdentityUpdate
- [ ] `mutation` cartCreate
- [ ] `mutation` cartDiscountCodesUpdate
- [ ] `mutation` cartLinesAdd
- [ ] `mutation` cartLinesRemove
- [ ] `mutation` cartLinesUpdate
- [ ] `mutation` cartNoteUpdate
- [ ] `mutation` cartSelectedDeliveryOptionsUpdate

#### Checkouts
- [ ] `query` locations
- [ ] `mutation` checkoutAttributesUpdateV2
- [ ] `mutation` checkoutCompleteFree
- [ ] `mutation` checkoutCompleteWithCreditCardV2
- [ ] `mutation` checkoutCompleteWithTokenizedPaymentV3
- [ ] `mutation` checkoutCreate
- [ ] `mutation` checkoutCustomerAssociateV2
- [ ] `mutation` checkoutCustomerDisassociateV2
- [ ] `mutation` checkoutDiscountCodeApplyV2
- [ ] `mutation` checkoutDiscountCodeRemove
- [ ] `mutation` checkoutEmailUpdateV2
- [ ] `mutation` checkoutGiftCardRemoveV2
- [ ] `mutation` checkoutGiftCardsAppend
- [ ] `mutation` checkoutLineItemsAdd
- [ ] `mutation` checkoutLineItemsRemove
- [ ] `mutation` checkoutLineItemsReplace
- [ ] `mutation` checkoutLineItemsUpdate
- [ ] `mutation` checkoutShippingAddressUpdateV2
- [ ] `mutation` checkoutShippingLineUpdate

#### Common
- [ ] `query` contentEntries
- [ ] `query` contentEntry
- [ ] `query` localization
- [ ] `query` node
- [ ] `query` nodes
- [ ] `query` publicApiVersions

#### Customers
- [ ] `query` customer `onProgress`
- [x] `mutation` customerAccessTokenCreate
- [ ] `mutation` customerAccessTokenCreateWithMultipass
- [x] `mutation` customerAccessTokenDelete
- [x] `mutation` customerAccessTokenRenew
- [ ] `mutation` customerActivate
- [ ] `mutation` customerActivateByUrl
- [ ] `mutation` customerAddressCreate
- [ ] `mutation` customerAddressDelete
- [ ] `mutation` customerAddressUpdate
- [ ] `mutation` customerCreate
- [ ] `mutation` customerDefaultAddressUpdate
- [ ] `mutation` customerRecover
- [ ] `mutation` customerReset
- [ ] `mutation` customerResetByUrl
- [ ] `mutation` customerUpdate

#### Online store
- [ ] `query` articles
- [ ] `query` blog
- [ ] `query` blogs
- [ ] `query` menu
- [ ] `query` page
- [ ] `query` pages
- [ ] `query` shop
- [ ] `query` urlRedirects

#### Orders

#### Products
- [ ] `query` collection
- [ ] `query` collections
- [ ] `query` product
- [ ] `query` productRecommendations
- [ ] `query` products
- [ ] `query` productTags
- [ ] `query` productTypes
