// Code generated by smithy-go-codegen DO NOT EDIT.

package snowballdevice

import (
	"context"
	"fmt"
	"github.com/aws/eks-anywhere/internal/aws-sdk-go-v2/service/snowballdevice/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAttachPciDevice struct {
}

func (*validateOpAttachPciDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAttachPciDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AttachPciDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAttachPciDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAutoStartConfiguration struct {
}

func (*validateOpCreateAutoStartConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAutoStartConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAutoStartConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAutoStartConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDirectNetworkInterface struct {
}

func (*validateOpCreateDirectNetworkInterface) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDirectNetworkInterface) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDirectNetworkInterfaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDirectNetworkInterfaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTags struct {
}

func (*validateOpCreateTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateVirtualNetworkInterface struct {
}

func (*validateOpCreateVirtualNetworkInterface) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateVirtualNetworkInterface) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateVirtualNetworkInterfaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateVirtualNetworkInterfaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAutoStartConfiguration struct {
}

func (*validateOpDeleteAutoStartConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAutoStartConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAutoStartConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAutoStartConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCertificate struct {
}

func (*validateOpDeleteCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDirectNetworkInterface struct {
}

func (*validateOpDeleteDirectNetworkInterface) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDirectNetworkInterface) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDirectNetworkInterfaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDirectNetworkInterfaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTags struct {
}

func (*validateOpDeleteTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteVirtualNetworkInterface struct {
}

func (*validateOpDeleteVirtualNetworkInterface) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteVirtualNetworkInterface) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteVirtualNetworkInterfaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteVirtualNetworkInterfaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeService struct {
}

func (*validateOpDescribeService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDetachPciDevice struct {
}

func (*validateOpDetachPciDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDetachPciDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DetachPciDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDetachPciDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCertificate struct {
}

func (*validateOpGetCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNotificationConfiguration struct {
}

func (*validateOpGetNotificationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNotificationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNotificationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNotificationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSecretAccessKey struct {
}

func (*validateOpGetSecretAccessKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSecretAccessKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSecretAccessKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSecretAccessKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAlertHistory struct {
}

func (*validateOpListAlertHistory) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAlertHistory) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAlertHistoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAlertHistoryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutNotificationConfiguration struct {
}

func (*validateOpPutNotificationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutNotificationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutNotificationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutNotificationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetAlertSubscription struct {
}

func (*validateOpSetAlertSubscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetAlertSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetAlertSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetAlertSubscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateAutoStartConfiguration struct {
}

func (*validateOpUpdateAutoStartConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateAutoStartConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateAutoStartConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateAutoStartConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDirectNetworkInterface struct {
}

func (*validateOpUpdateDirectNetworkInterface) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDirectNetworkInterface) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDirectNetworkInterfaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDirectNetworkInterfaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdatePhysicalNetworkConfiguration struct {
}

func (*validateOpUpdatePhysicalNetworkConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdatePhysicalNetworkConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdatePhysicalNetworkConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdatePhysicalNetworkConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateVirtualNetworkInterface struct {
}

func (*validateOpUpdateVirtualNetworkInterface) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateVirtualNetworkInterface) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateVirtualNetworkInterfaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateVirtualNetworkInterfaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAttachPciDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAttachPciDevice{}, middleware.After)
}

func addOpCreateAutoStartConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAutoStartConfiguration{}, middleware.After)
}

func addOpCreateDirectNetworkInterfaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDirectNetworkInterface{}, middleware.After)
}

func addOpCreateTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateTags{}, middleware.After)
}

func addOpCreateVirtualNetworkInterfaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateVirtualNetworkInterface{}, middleware.After)
}

func addOpDeleteAutoStartConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAutoStartConfiguration{}, middleware.After)
}

func addOpDeleteCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCertificate{}, middleware.After)
}

func addOpDeleteDirectNetworkInterfaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDirectNetworkInterface{}, middleware.After)
}

func addOpDeleteTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTags{}, middleware.After)
}

func addOpDeleteVirtualNetworkInterfaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteVirtualNetworkInterface{}, middleware.After)
}

func addOpDescribeServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeService{}, middleware.After)
}

func addOpDetachPciDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDetachPciDevice{}, middleware.After)
}

func addOpGetCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCertificate{}, middleware.After)
}

func addOpGetNotificationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNotificationConfiguration{}, middleware.After)
}

func addOpGetSecretAccessKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSecretAccessKey{}, middleware.After)
}

func addOpListAlertHistoryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAlertHistory{}, middleware.After)
}

func addOpPutNotificationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutNotificationConfiguration{}, middleware.After)
}

func addOpSetAlertSubscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetAlertSubscription{}, middleware.After)
}

func addOpUpdateAutoStartConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateAutoStartConfiguration{}, middleware.After)
}

func addOpUpdateDirectNetworkInterfaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDirectNetworkInterface{}, middleware.After)
}

func addOpUpdatePhysicalNetworkConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdatePhysicalNetworkConfiguration{}, middleware.After)
}

func addOpUpdateVirtualNetworkInterfaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateVirtualNetworkInterface{}, middleware.After)
}

func validateStaticIpAddressConfiguration(v *types.StaticIpAddressConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StaticIpAddressConfiguration"}
	if v.IpAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IpAddress"))
	}
	if v.Netmask == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Netmask"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAttachPciDeviceInput(v *AttachPciDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AttachPciDeviceInput"}
	if v.InstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceId"))
	}
	if v.PciDeviceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PciDeviceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAutoStartConfigurationInput(v *CreateAutoStartConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAutoStartConfigurationInput"}
	if len(v.PhysicalConnectorType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("PhysicalConnectorType"))
	}
	if len(v.IpAddressAssignment) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("IpAddressAssignment"))
	}
	if v.StaticIpAddressConfiguration != nil {
		if err := validateStaticIpAddressConfiguration(v.StaticIpAddressConfiguration); err != nil {
			invalidParams.AddNested("StaticIpAddressConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.LaunchTemplateId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LaunchTemplateId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDirectNetworkInterfaceInput(v *CreateDirectNetworkInterfaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDirectNetworkInterfaceInput"}
	if v.PhysicalNetworkInterfaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PhysicalNetworkInterfaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTagsInput(v *CreateTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTagsInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateVirtualNetworkInterfaceInput(v *CreateVirtualNetworkInterfaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateVirtualNetworkInterfaceInput"}
	if v.PhysicalNetworkInterfaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PhysicalNetworkInterfaceId"))
	}
	if len(v.IpAddressAssignment) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("IpAddressAssignment"))
	}
	if v.StaticIpAddressConfiguration != nil {
		if err := validateStaticIpAddressConfiguration(v.StaticIpAddressConfiguration); err != nil {
			invalidParams.AddNested("StaticIpAddressConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAutoStartConfigurationInput(v *DeleteAutoStartConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAutoStartConfigurationInput"}
	if v.AutoStartConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AutoStartConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCertificateInput(v *DeleteCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDirectNetworkInterfaceInput(v *DeleteDirectNetworkInterfaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDirectNetworkInterfaceInput"}
	if v.DirectNetworkInterfaceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectNetworkInterfaceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTagsInput(v *DeleteTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTagsInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteVirtualNetworkInterfaceInput(v *DeleteVirtualNetworkInterfaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteVirtualNetworkInterfaceInput"}
	if v.VirtualNetworkInterfaceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VirtualNetworkInterfaceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeServiceInput(v *DescribeServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeServiceInput"}
	if v.ServiceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDetachPciDeviceInput(v *DetachPciDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DetachPciDeviceInput"}
	if v.PciDeviceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PciDeviceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCertificateInput(v *GetCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNotificationConfigurationInput(v *GetNotificationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNotificationConfigurationInput"}
	if v.ServiceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSecretAccessKeyInput(v *GetSecretAccessKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSecretAccessKeyInput"}
	if v.AccessKeyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessKeyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAlertHistoryInput(v *ListAlertHistoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAlertHistoryInput"}
	if v.Namespace == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Namespace"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutNotificationConfigurationInput(v *PutNotificationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutNotificationConfigurationInput"}
	if v.ServiceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceId"))
	}
	if v.Enabled == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Enabled"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetAlertSubscriptionInput(v *SetAlertSubscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetAlertSubscriptionInput"}
	if v.Namespace == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Namespace"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.State) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("State"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateAutoStartConfigurationInput(v *UpdateAutoStartConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateAutoStartConfigurationInput"}
	if v.AutoStartConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AutoStartConfigurationArn"))
	}
	if v.StaticIpAddressConfiguration != nil {
		if err := validateStaticIpAddressConfiguration(v.StaticIpAddressConfiguration); err != nil {
			invalidParams.AddNested("StaticIpAddressConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateDirectNetworkInterfaceInput(v *UpdateDirectNetworkInterfaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDirectNetworkInterfaceInput"}
	if v.DirectNetworkInterfaceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DirectNetworkInterfaceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdatePhysicalNetworkConfigurationInput(v *UpdatePhysicalNetworkConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePhysicalNetworkConfigurationInput"}
	if v.PhysicalNetworkInterfaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PhysicalNetworkInterfaceId"))
	}
	if len(v.IpAddressAssignment) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("IpAddressAssignment"))
	}
	if v.StaticIpAddressConfiguration != nil {
		if err := validateStaticIpAddressConfiguration(v.StaticIpAddressConfiguration); err != nil {
			invalidParams.AddNested("StaticIpAddressConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateVirtualNetworkInterfaceInput(v *UpdateVirtualNetworkInterfaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateVirtualNetworkInterfaceInput"}
	if v.VirtualNetworkInterfaceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VirtualNetworkInterfaceArn"))
	}
	if len(v.IpAddressAssignment) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("IpAddressAssignment"))
	}
	if v.StaticIpAddressConfiguration != nil {
		if err := validateStaticIpAddressConfiguration(v.StaticIpAddressConfiguration); err != nil {
			invalidParams.AddNested("StaticIpAddressConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
