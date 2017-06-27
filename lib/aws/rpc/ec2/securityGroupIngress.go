// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/resource/plugin"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

/* RPC stubs for SecurityGroupIngress resource provider */

// SecurityGroupIngressToken is the type token corresponding to the SecurityGroupIngress package type.
const SecurityGroupIngressToken = tokens.Type("aws:ec2/securityGroupIngress:SecurityGroupIngress")

// SecurityGroupIngressProviderOps is a pluggable interface for SecurityGroupIngress-related management functionality.
type SecurityGroupIngressProviderOps interface {
    Check(ctx context.Context, obj *SecurityGroupIngress, property string) error
    Create(ctx context.Context, obj *SecurityGroupIngress) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*SecurityGroupIngress, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *SecurityGroupIngress, new *SecurityGroupIngress, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *SecurityGroupIngress, new *SecurityGroupIngress, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// SecurityGroupIngressProvider is a dynamic gRPC-based plugin for managing SecurityGroupIngress resources.
type SecurityGroupIngressProvider struct {
    ops SecurityGroupIngressProviderOps
}

// NewSecurityGroupIngressProvider allocates a resource provider that delegates to a ops instance.
func NewSecurityGroupIngressProvider(ops SecurityGroupIngressProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &SecurityGroupIngressProvider{ops: ops}
}

func (p *SecurityGroupIngressProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupIngressToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    if failure := p.ops.Check(ctx, obj, ""); failure != nil {
        failures = append(failures, failure)
    }
    unks := req.GetUnknowns()
    if !unks["name"] {
        if failure := p.ops.Check(ctx, obj, "name"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "name", failure))
        }
    }
    if !unks["ipProtocol"] {
        if failure := p.ops.Check(ctx, obj, "ipProtocol"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "ipProtocol", failure))
        }
    }
    if !unks["cidrIp"] {
        if failure := p.ops.Check(ctx, obj, "cidrIp"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "cidrIp", failure))
        }
    }
    if !unks["cidrIpv6"] {
        if failure := p.ops.Check(ctx, obj, "cidrIpv6"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "cidrIpv6", failure))
        }
    }
    if !unks["fromPort"] {
        if failure := p.ops.Check(ctx, obj, "fromPort"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "fromPort", failure))
        }
    }
    if !unks["group"] {
        if failure := p.ops.Check(ctx, obj, "group"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "group", failure))
        }
    }
    if !unks["groupName"] {
        if failure := p.ops.Check(ctx, obj, "groupName"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "groupName", failure))
        }
    }
    if !unks["sourceSecurityGroup"] {
        if failure := p.ops.Check(ctx, obj, "sourceSecurityGroup"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "sourceSecurityGroup", failure))
        }
    }
    if !unks["sourceSecurityGroupName"] {
        if failure := p.ops.Check(ctx, obj, "sourceSecurityGroupName"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "sourceSecurityGroupName", failure))
        }
    }
    if !unks["sourceSecurityGroupOwnerId"] {
        if failure := p.ops.Check(ctx, obj, "sourceSecurityGroupOwnerId"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "sourceSecurityGroupOwnerId", failure))
        }
    }
    if !unks["toPort"] {
        if failure := p.ops.Check(ctx, obj, "toPort"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupIngress", "toPort", failure))
        }
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *SecurityGroupIngressProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupIngressToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[SecurityGroupIngress_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *SecurityGroupIngressProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupIngressToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{Id: string(id)}, nil
}

func (p *SecurityGroupIngressProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupIngressToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: plugin.MarshalProperties(
            nil, resource.NewPropertyMap(obj), plugin.MarshalOptions{}),
    }, nil
}

func (p *SecurityGroupIngressProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupIngressToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("ipProtocol") {
            replaces = append(replaces, "ipProtocol")
        }
        if diff.Changed("cidrIp") {
            replaces = append(replaces, "cidrIp")
        }
        if diff.Changed("cidrIpv6") {
            replaces = append(replaces, "cidrIpv6")
        }
        if diff.Changed("fromPort") {
            replaces = append(replaces, "fromPort")
        }
        if diff.Changed("group") {
            replaces = append(replaces, "group")
        }
        if diff.Changed("groupName") {
            replaces = append(replaces, "groupName")
        }
        if diff.Changed("sourceSecurityGroup") {
            replaces = append(replaces, "sourceSecurityGroup")
        }
        if diff.Changed("sourceSecurityGroupName") {
            replaces = append(replaces, "sourceSecurityGroupName")
        }
        if diff.Changed("sourceSecurityGroupOwnerId") {
            replaces = append(replaces, "sourceSecurityGroupOwnerId")
        }
        if diff.Changed("toPort") {
            replaces = append(replaces, "toPort")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *SecurityGroupIngressProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupIngressToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SecurityGroupIngressProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupIngressToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SecurityGroupIngressProvider) Unmarshal(
    v *pbstruct.Struct) (*SecurityGroupIngress, resource.PropertyMap, error) {
    var obj SecurityGroupIngress
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable SecurityGroupIngress structure(s) */

// SecurityGroupIngress is a marshalable representation of its corresponding IDL type.
type SecurityGroupIngress struct {
    Name *string `lumi:"name,optional"`
    IPProtocol string `lumi:"ipProtocol"`
    CIDRIP *string `lumi:"cidrIp,optional"`
    CIDRIPv6 *string `lumi:"cidrIpv6,optional"`
    FromPort *float64 `lumi:"fromPort,optional"`
    Group *resource.ID `lumi:"group,optional"`
    GroupName *string `lumi:"groupName,optional"`
    SourceSecurityGroup *resource.ID `lumi:"sourceSecurityGroup,optional"`
    SourceSecurityGroupName *string `lumi:"sourceSecurityGroupName,optional"`
    SourceSecurityGroupOwnerId *string `lumi:"sourceSecurityGroupOwnerId,optional"`
    ToPort *float64 `lumi:"toPort,optional"`
}

// SecurityGroupIngress's properties have constants to make dealing with diffs and property bags easier.
const (
    SecurityGroupIngress_Name = "name"
    SecurityGroupIngress_IPProtocol = "ipProtocol"
    SecurityGroupIngress_CIDRIP = "cidrIp"
    SecurityGroupIngress_CIDRIPv6 = "cidrIpv6"
    SecurityGroupIngress_FromPort = "fromPort"
    SecurityGroupIngress_Group = "group"
    SecurityGroupIngress_GroupName = "groupName"
    SecurityGroupIngress_SourceSecurityGroup = "sourceSecurityGroup"
    SecurityGroupIngress_SourceSecurityGroupName = "sourceSecurityGroupName"
    SecurityGroupIngress_SourceSecurityGroupOwnerId = "sourceSecurityGroupOwnerId"
    SecurityGroupIngress_ToPort = "toPort"
)


