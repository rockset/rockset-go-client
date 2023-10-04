# TLSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCert** | Pointer to **string** | PEM-formatted certificate chain of the Certificate Authority used to verify remote server. If empty, Rockset, will use publicly trusted CAs | [optional] 
**ClientCert** | **string** | PEM-formatted certificate chain to use for client authentication | 
**ClientCertExpiry** | Pointer to **int64** | Expiration date of the client certificate (represented as number of ms since epoch) | [optional] [readonly] 
**ClientCertSubject** | Pointer to **string** | Subject of the client certificate, containing common name and other attributes | [optional] [readonly] 
**ClientKey** | **string** | PEM-formatted private key to be used for client authentication | 

## Methods

### NewTLSConfig

`func NewTLSConfig(clientCert string, clientKey string, ) *TLSConfig`

NewTLSConfig instantiates a new TLSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSConfigWithDefaults

`func NewTLSConfigWithDefaults() *TLSConfig`

NewTLSConfigWithDefaults instantiates a new TLSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCert

`func (o *TLSConfig) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *TLSConfig) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *TLSConfig) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *TLSConfig) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetClientCert

`func (o *TLSConfig) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *TLSConfig) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *TLSConfig) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.


### GetClientCertExpiry

`func (o *TLSConfig) GetClientCertExpiry() int64`

GetClientCertExpiry returns the ClientCertExpiry field if non-nil, zero value otherwise.

### GetClientCertExpiryOk

`func (o *TLSConfig) GetClientCertExpiryOk() (*int64, bool)`

GetClientCertExpiryOk returns a tuple with the ClientCertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertExpiry

`func (o *TLSConfig) SetClientCertExpiry(v int64)`

SetClientCertExpiry sets ClientCertExpiry field to given value.

### HasClientCertExpiry

`func (o *TLSConfig) HasClientCertExpiry() bool`

HasClientCertExpiry returns a boolean if a field has been set.

### GetClientCertSubject

`func (o *TLSConfig) GetClientCertSubject() string`

GetClientCertSubject returns the ClientCertSubject field if non-nil, zero value otherwise.

### GetClientCertSubjectOk

`func (o *TLSConfig) GetClientCertSubjectOk() (*string, bool)`

GetClientCertSubjectOk returns a tuple with the ClientCertSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertSubject

`func (o *TLSConfig) SetClientCertSubject(v string)`

SetClientCertSubject sets ClientCertSubject field to given value.

### HasClientCertSubject

`func (o *TLSConfig) HasClientCertSubject() bool`

HasClientCertSubject returns a boolean if a field has been set.

### GetClientKey

`func (o *TLSConfig) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *TLSConfig) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *TLSConfig) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


