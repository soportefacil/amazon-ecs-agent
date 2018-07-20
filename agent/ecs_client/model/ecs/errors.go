// Copyright 2014-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package ecs

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have authorization to perform the requested action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAttributeLimitExceededException for service response error code
	// "AttributeLimitExceededException".
	//
	// You can apply up to 10 custom attributes per resource. You can view the attributes
	// of a resource with ListAttributes. You can remove existing attributes on
	// a resource with DeleteAttributes.
	ErrCodeAttributeLimitExceededException = "AttributeLimitExceededException"

	// ErrCodeBlockedException for service response error code
	// "BlockedException".
	//
	// Your AWS account has been blocked. Contact AWS Support (http://aws.amazon.com/contact-us/)
	// for more information.
	ErrCodeBlockedException = "BlockedException"

	// ErrCodeClientException for service response error code
	// "ClientException".
	//
	// These errors are usually caused by a client action, such as using an action
	// or resource on behalf of a user that doesn't have permissions to use the
	// action or resource, or specifying an identifier that is not valid.
	ErrCodeClientException = "ClientException"

	// ErrCodeClusterContainsContainerInstancesException for service response error code
	// "ClusterContainsContainerInstancesException".
	//
	// You cannot delete a cluster that has registered container instances. You
	// must first deregister the container instances before you can delete the cluster.
	// For more information, see DeregisterContainerInstance.
	ErrCodeClusterContainsContainerInstancesException = "ClusterContainsContainerInstancesException"

	// ErrCodeClusterContainsServicesException for service response error code
	// "ClusterContainsServicesException".
	//
	// You cannot delete a cluster that contains services. You must first update
	// the service to reduce its desired task count to 0 and then delete the service.
	// For more information, see UpdateService and DeleteService.
	ErrCodeClusterContainsServicesException = "ClusterContainsServicesException"

	// ErrCodeClusterContainsTasksException for service response error code
	// "ClusterContainsTasksException".
	//
	// You cannot delete a cluster that has active tasks.
	ErrCodeClusterContainsTasksException = "ClusterContainsTasksException"

	// ErrCodeClusterNotFoundException for service response error code
	// "ClusterNotFoundException".
	//
	// The specified cluster could not be found. You can view your available clusters
	// with ListClusters. Amazon ECS clusters are region-specific.
	ErrCodeClusterNotFoundException = "ClusterNotFoundException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// The specified parameter is invalid. Review the available parameters for the
	// API request.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeMissingVersionException for service response error code
	// "MissingVersionException".
	//
	// Amazon ECS is unable to determine the current version of the Amazon ECS container
	// agent on the container instance and does not have enough information to proceed
	// with an update. This could be because the agent running on the container
	// instance is an older or custom version that does not use our version information.
	ErrCodeMissingVersionException = "MissingVersionException"

	// ErrCodeNoUpdateAvailableException for service response error code
	// "NoUpdateAvailableException".
	//
	// There is no update available for this Amazon ECS container agent. This could
	// be because the agent is already running the latest version, or it is so old
	// that there is no update path to the current version.
	ErrCodeNoUpdateAvailableException = "NoUpdateAvailableException"

	// ErrCodePlatformTaskDefinitionIncompatibilityException for service response error code
	// "PlatformTaskDefinitionIncompatibilityException".
	//
	// The specified platform version does not satisfy the task definition's required
	// capabilities.
	ErrCodePlatformTaskDefinitionIncompatibilityException = "PlatformTaskDefinitionIncompatibilityException"

	// ErrCodePlatformUnknownException for service response error code
	// "PlatformUnknownException".
	//
	// The specified platform version does not exist.
	ErrCodePlatformUnknownException = "PlatformUnknownException"

	// ErrCodeServerException for service response error code
	// "ServerException".
	//
	// These errors are usually caused by a server issue.
	ErrCodeServerException = "ServerException"

	// ErrCodeServiceNotActiveException for service response error code
	// "ServiceNotActiveException".
	//
	// The specified service is not active. You can't update a service that is inactive.
	// If you have previously deleted a service, you can re-create it with CreateService.
	ErrCodeServiceNotActiveException = "ServiceNotActiveException"

	// ErrCodeServiceNotFoundException for service response error code
	// "ServiceNotFoundException".
	//
	// The specified service could not be found. You can view your available services
	// with ListServices. Amazon ECS services are cluster-specific and region-specific.
	ErrCodeServiceNotFoundException = "ServiceNotFoundException"

	// ErrCodeTargetNotFoundException for service response error code
	// "TargetNotFoundException".
	//
	// The specified target could not be found. You can view your available container
	// instances with ListContainerInstances. Amazon ECS container instances are
	// cluster-specific and region-specific.
	ErrCodeTargetNotFoundException = "TargetNotFoundException"

	// ErrCodeUnsupportedFeatureException for service response error code
	// "UnsupportedFeatureException".
	//
	// The specified task is not supported in this region.
	ErrCodeUnsupportedFeatureException = "UnsupportedFeatureException"

	// ErrCodeUpdateInProgressException for service response error code
	// "UpdateInProgressException".
	//
	// There is already a current Amazon ECS container agent update in progress
	// on the specified container instance. If the container agent becomes disconnected
	// while it is in a transitional stage, such as PENDING or STAGING, the update
	// process can get stuck in that state. However, when the agent reconnects,
	// it resumes where it stopped previously.
	ErrCodeUpdateInProgressException = "UpdateInProgressException"
)
