// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

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

/* Marshalable LoginProfile structure(s) */

// LoginProfile is a marshalable representation of its corresponding IDL type.
type LoginProfile struct {
    Password string `lumi:"password"`
    PasswordResetRequired *bool `lumi:"passwordResetRequired,optional"`
}

// LoginProfile's properties have constants to make dealing with diffs and property bags easier.
const (
    LoginProfile_Password = "password"
    LoginProfile_PasswordResetRequired = "passwordResetRequired"
)

/* RPC stubs for User resource provider */

// UserToken is the type token corresponding to the User package type.
const UserToken = tokens.Type("aws:iam/user:User")

// UserProviderOps is a pluggable interface for User-related management functionality.
type UserProviderOps interface {
    Check(ctx context.Context, obj *User, property string) error
    Create(ctx context.Context, obj *User) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*User, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *User, new *User, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *User, new *User, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// UserProvider is a dynamic gRPC-based plugin for managing User resources.
type UserProvider struct {
    ops UserProviderOps
}

// NewUserProvider allocates a resource provider that delegates to a ops instance.
func NewUserProvider(ops UserProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &UserProvider{ops: ops}
}

func (p *UserProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(UserToken))
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
                resource.NewPropertyError("User", "name", failure))
        }
    }
    if !unks["userName"] {
        if failure := p.ops.Check(ctx, obj, "userName"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("User", "userName", failure))
        }
    }
    if !unks["groups"] {
        if failure := p.ops.Check(ctx, obj, "groups"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("User", "groups", failure))
        }
    }
    if !unks["loginProfile"] {
        if failure := p.ops.Check(ctx, obj, "loginProfile"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("User", "loginProfile", failure))
        }
    }
    if !unks["managedPolicies"] {
        if failure := p.ops.Check(ctx, obj, "managedPolicies"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("User", "managedPolicies", failure))
        }
    }
    if !unks["path"] {
        if failure := p.ops.Check(ctx, obj, "path"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("User", "path", failure))
        }
    }
    if !unks["policies"] {
        if failure := p.ops.Check(ctx, obj, "policies"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("User", "policies", failure))
        }
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *UserProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(UserToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[User_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *UserProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(UserToken))
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

func (p *UserProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(UserToken))
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

func (p *UserProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(UserToken))
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
        if diff.Changed("userName") {
            replaces = append(replaces, "userName")
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

func (p *UserProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(UserToken))
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

func (p *UserProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(UserToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *UserProvider) Unmarshal(
    v *pbstruct.Struct) (*User, resource.PropertyMap, error) {
    var obj User
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable User structure(s) */

// User is a marshalable representation of its corresponding IDL type.
type User struct {
    Name *string `lumi:"name,optional"`
    UserName *string `lumi:"userName,optional"`
    Groups *[]resource.ID `lumi:"groups,optional"`
    LoginProfile *LoginProfile `lumi:"loginProfile,optional"`
    ManagedPolicies *[]resource.ID `lumi:"managedPolicies,optional"`
    Path *string `lumi:"path,optional"`
    Policies *[]InlinePolicy `lumi:"policies,optional"`
}

// User's properties have constants to make dealing with diffs and property bags easier.
const (
    User_Name = "name"
    User_UserName = "userName"
    User_Groups = "groups"
    User_LoginProfile = "loginProfile"
    User_ManagedPolicies = "managedPolicies"
    User_Path = "path"
    User_Policies = "policies"
)


