package streamanalytics

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// BindingType enumerates the values for binding type.
type BindingType string

const (
	// BindingTypeFunctionRetrieveDefaultDefinitionParameters ...
	BindingTypeFunctionRetrieveDefaultDefinitionParameters BindingType = "FunctionRetrieveDefaultDefinitionParameters"
	// BindingTypeMicrosoftMachineLearningWebService ...
	BindingTypeMicrosoftMachineLearningWebService BindingType = "Microsoft.MachineLearning/WebService"
	// BindingTypeMicrosoftStreamAnalyticsJavascriptUdf ...
	BindingTypeMicrosoftStreamAnalyticsJavascriptUdf BindingType = "Microsoft.StreamAnalytics/JavascriptUdf"
)

// PossibleBindingTypeValues returns an array of possible values for the BindingType const type.
func PossibleBindingTypeValues() []BindingType {
	return []BindingType{BindingTypeFunctionRetrieveDefaultDefinitionParameters, BindingTypeMicrosoftMachineLearningWebService, BindingTypeMicrosoftStreamAnalyticsJavascriptUdf}
}

// CompatibilityLevel enumerates the values for compatibility level.
type CompatibilityLevel string

const (
	// OneFullStopZero ...
	OneFullStopZero CompatibilityLevel = "1.0"
)

// PossibleCompatibilityLevelValues returns an array of possible values for the CompatibilityLevel const type.
func PossibleCompatibilityLevelValues() []CompatibilityLevel {
	return []CompatibilityLevel{OneFullStopZero}
}

// Encoding enumerates the values for encoding.
type Encoding string

const (
	// UTF8 ...
	UTF8 Encoding = "UTF8"
)

// PossibleEncodingValues returns an array of possible values for the Encoding const type.
func PossibleEncodingValues() []Encoding {
	return []Encoding{UTF8}
}

// EventsOutOfOrderPolicy enumerates the values for events out of order policy.
type EventsOutOfOrderPolicy string

const (
	// Adjust ...
	Adjust EventsOutOfOrderPolicy = "Adjust"
	// Drop ...
	Drop EventsOutOfOrderPolicy = "Drop"
)

// PossibleEventsOutOfOrderPolicyValues returns an array of possible values for the EventsOutOfOrderPolicy const type.
func PossibleEventsOutOfOrderPolicyValues() []EventsOutOfOrderPolicy {
	return []EventsOutOfOrderPolicy{Adjust, Drop}
}

// JSONOutputSerializationFormat enumerates the values for json output serialization format.
type JSONOutputSerializationFormat string

const (
	// Array ...
	Array JSONOutputSerializationFormat = "Array"
	// LineSeparated ...
	LineSeparated JSONOutputSerializationFormat = "LineSeparated"
)

// PossibleJSONOutputSerializationFormatValues returns an array of possible values for the JSONOutputSerializationFormat const type.
func PossibleJSONOutputSerializationFormatValues() []JSONOutputSerializationFormat {
	return []JSONOutputSerializationFormat{Array, LineSeparated}
}

// OutputErrorPolicy enumerates the values for output error policy.
type OutputErrorPolicy string

const (
	// OutputErrorPolicyDrop ...
	OutputErrorPolicyDrop OutputErrorPolicy = "Drop"
	// OutputErrorPolicyStop ...
	OutputErrorPolicyStop OutputErrorPolicy = "Stop"
)

// PossibleOutputErrorPolicyValues returns an array of possible values for the OutputErrorPolicy const type.
func PossibleOutputErrorPolicyValues() []OutputErrorPolicy {
	return []OutputErrorPolicy{OutputErrorPolicyDrop, OutputErrorPolicyStop}
}

// OutputStartMode enumerates the values for output start mode.
type OutputStartMode string

const (
	// CustomTime ...
	CustomTime OutputStartMode = "CustomTime"
	// JobStartTime ...
	JobStartTime OutputStartMode = "JobStartTime"
	// LastOutputEventTime ...
	LastOutputEventTime OutputStartMode = "LastOutputEventTime"
)

// PossibleOutputStartModeValues returns an array of possible values for the OutputStartMode const type.
func PossibleOutputStartModeValues() []OutputStartMode {
	return []OutputStartMode{CustomTime, JobStartTime, LastOutputEventTime}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Standard ...
	Standard SkuName = "Standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Standard}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeAvro ...
	TypeAvro Type = "Avro"
	// TypeCsv ...
	TypeCsv Type = "Csv"
	// TypeJSON ...
	TypeJSON Type = "Json"
	// TypeSerialization ...
	TypeSerialization Type = "Serialization"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeAvro, TypeCsv, TypeJSON, TypeSerialization}
}

// TypeBasicFunctionBinding enumerates the values for type basic function binding.
type TypeBasicFunctionBinding string

const (
	// TypeFunctionBinding ...
	TypeFunctionBinding TypeBasicFunctionBinding = "FunctionBinding"
	// TypeMicrosoftMachineLearningWebService ...
	TypeMicrosoftMachineLearningWebService TypeBasicFunctionBinding = "Microsoft.MachineLearning/WebService"
	// TypeMicrosoftStreamAnalyticsJavascriptUdf ...
	TypeMicrosoftStreamAnalyticsJavascriptUdf TypeBasicFunctionBinding = "Microsoft.StreamAnalytics/JavascriptUdf"
)

// PossibleTypeBasicFunctionBindingValues returns an array of possible values for the TypeBasicFunctionBinding const type.
func PossibleTypeBasicFunctionBindingValues() []TypeBasicFunctionBinding {
	return []TypeBasicFunctionBinding{TypeFunctionBinding, TypeMicrosoftMachineLearningWebService, TypeMicrosoftStreamAnalyticsJavascriptUdf}
}

// TypeBasicFunctionProperties enumerates the values for type basic function properties.
type TypeBasicFunctionProperties string

const (
	// TypeFunctionProperties ...
	TypeFunctionProperties TypeBasicFunctionProperties = "FunctionProperties"
	// TypeScalar ...
	TypeScalar TypeBasicFunctionProperties = "Scalar"
)

// PossibleTypeBasicFunctionPropertiesValues returns an array of possible values for the TypeBasicFunctionProperties const type.
func PossibleTypeBasicFunctionPropertiesValues() []TypeBasicFunctionProperties {
	return []TypeBasicFunctionProperties{TypeFunctionProperties, TypeScalar}
}

// TypeBasicInputProperties enumerates the values for type basic input properties.
type TypeBasicInputProperties string

const (
	// TypeInputProperties ...
	TypeInputProperties TypeBasicInputProperties = "InputProperties"
	// TypeReference ...
	TypeReference TypeBasicInputProperties = "Reference"
	// TypeStream ...
	TypeStream TypeBasicInputProperties = "Stream"
)

// PossibleTypeBasicInputPropertiesValues returns an array of possible values for the TypeBasicInputProperties const type.
func PossibleTypeBasicInputPropertiesValues() []TypeBasicInputProperties {
	return []TypeBasicInputProperties{TypeInputProperties, TypeReference, TypeStream}
}

// TypeBasicOutputDataSource enumerates the values for type basic output data source.
type TypeBasicOutputDataSource string

const (
	// TypeMicrosoftDataLakeAccounts ...
	TypeMicrosoftDataLakeAccounts TypeBasicOutputDataSource = "Microsoft.DataLake/Accounts"
	// TypeMicrosoftServiceBusEventHub ...
	TypeMicrosoftServiceBusEventHub TypeBasicOutputDataSource = "Microsoft.ServiceBus/EventHub"
	// TypeMicrosoftServiceBusQueue ...
	TypeMicrosoftServiceBusQueue TypeBasicOutputDataSource = "Microsoft.ServiceBus/Queue"
	// TypeMicrosoftServiceBusTopic ...
	TypeMicrosoftServiceBusTopic TypeBasicOutputDataSource = "Microsoft.ServiceBus/Topic"
	// TypeMicrosoftSQLServerDatabase ...
	TypeMicrosoftSQLServerDatabase TypeBasicOutputDataSource = "Microsoft.Sql/Server/Database"
	// TypeMicrosoftStorageBlob ...
	TypeMicrosoftStorageBlob TypeBasicOutputDataSource = "Microsoft.Storage/Blob"
	// TypeMicrosoftStorageDocumentDB ...
	TypeMicrosoftStorageDocumentDB TypeBasicOutputDataSource = "Microsoft.Storage/DocumentDB"
	// TypeMicrosoftStorageTable ...
	TypeMicrosoftStorageTable TypeBasicOutputDataSource = "Microsoft.Storage/Table"
	// TypeOutputDataSource ...
	TypeOutputDataSource TypeBasicOutputDataSource = "OutputDataSource"
	// TypePowerBI ...
	TypePowerBI TypeBasicOutputDataSource = "PowerBI"
)

// PossibleTypeBasicOutputDataSourceValues returns an array of possible values for the TypeBasicOutputDataSource const type.
func PossibleTypeBasicOutputDataSourceValues() []TypeBasicOutputDataSource {
	return []TypeBasicOutputDataSource{TypeMicrosoftDataLakeAccounts, TypeMicrosoftServiceBusEventHub, TypeMicrosoftServiceBusQueue, TypeMicrosoftServiceBusTopic, TypeMicrosoftSQLServerDatabase, TypeMicrosoftStorageBlob, TypeMicrosoftStorageDocumentDB, TypeMicrosoftStorageTable, TypeOutputDataSource, TypePowerBI}
}

// TypeBasicReferenceInputDataSource enumerates the values for type basic reference input data source.
type TypeBasicReferenceInputDataSource string

const (
	// TypeBasicReferenceInputDataSourceTypeMicrosoftStorageBlob ...
	TypeBasicReferenceInputDataSourceTypeMicrosoftStorageBlob TypeBasicReferenceInputDataSource = "Microsoft.Storage/Blob"
	// TypeBasicReferenceInputDataSourceTypeReferenceInputDataSource ...
	TypeBasicReferenceInputDataSourceTypeReferenceInputDataSource TypeBasicReferenceInputDataSource = "ReferenceInputDataSource"
)

// PossibleTypeBasicReferenceInputDataSourceValues returns an array of possible values for the TypeBasicReferenceInputDataSource const type.
func PossibleTypeBasicReferenceInputDataSourceValues() []TypeBasicReferenceInputDataSource {
	return []TypeBasicReferenceInputDataSource{TypeBasicReferenceInputDataSourceTypeMicrosoftStorageBlob, TypeBasicReferenceInputDataSourceTypeReferenceInputDataSource}
}

// TypeBasicStreamInputDataSource enumerates the values for type basic stream input data source.
type TypeBasicStreamInputDataSource string

const (
	// TypeBasicStreamInputDataSourceTypeMicrosoftDevicesIotHubs ...
	TypeBasicStreamInputDataSourceTypeMicrosoftDevicesIotHubs TypeBasicStreamInputDataSource = "Microsoft.Devices/IotHubs"
	// TypeBasicStreamInputDataSourceTypeMicrosoftServiceBusEventHub ...
	TypeBasicStreamInputDataSourceTypeMicrosoftServiceBusEventHub TypeBasicStreamInputDataSource = "Microsoft.ServiceBus/EventHub"
	// TypeBasicStreamInputDataSourceTypeMicrosoftStorageBlob ...
	TypeBasicStreamInputDataSourceTypeMicrosoftStorageBlob TypeBasicStreamInputDataSource = "Microsoft.Storage/Blob"
	// TypeBasicStreamInputDataSourceTypeStreamInputDataSource ...
	TypeBasicStreamInputDataSourceTypeStreamInputDataSource TypeBasicStreamInputDataSource = "StreamInputDataSource"
)

// PossibleTypeBasicStreamInputDataSourceValues returns an array of possible values for the TypeBasicStreamInputDataSource const type.
func PossibleTypeBasicStreamInputDataSourceValues() []TypeBasicStreamInputDataSource {
	return []TypeBasicStreamInputDataSource{TypeBasicStreamInputDataSourceTypeMicrosoftDevicesIotHubs, TypeBasicStreamInputDataSourceTypeMicrosoftServiceBusEventHub, TypeBasicStreamInputDataSourceTypeMicrosoftStorageBlob, TypeBasicStreamInputDataSourceTypeStreamInputDataSource}
}

// UdfType enumerates the values for udf type.
type UdfType string

const (
	// Scalar ...
	Scalar UdfType = "Scalar"
)

// PossibleUdfTypeValues returns an array of possible values for the UdfType const type.
func PossibleUdfTypeValues() []UdfType {
	return []UdfType{Scalar}
}
