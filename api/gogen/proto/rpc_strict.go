// Copyright (C) 2023  Ryan SU (https://github.com/suyuan32)

package proto

import "github.com/suyuan32/goctls/rpc/parser"

// rpcPresenceFlags marks which standard CRUD-style gRPC methods exist on the service for this model.
// Naming matches generated clients: create{Model}, update{Model}, delete{Model}, get{Model}List, get{Model}ById.
type rpcPresenceFlags struct {
	Create  bool
	Update  bool
	Delete  bool
	List    bool
	GetById bool
}

func rpcNamesForModel(modelName string) (create, update, delete, list, getById string) {
	return "create" + modelName,
		"update" + modelName,
		"delete" + modelName,
		"get" + modelName + "List",
		"get" + modelName + "ById"
}

func findService(services parser.Services, name string) *parser.Service {
	for i := range services {
		if services[i].Name == name {
			return &services[i]
		}
	}
	return nil
}

func computeRPCPresence(p *parser.Proto, rpcServiceName, modelName string) rpcPresenceFlags {
	var out rpcPresenceFlags
	cn, un, dn, ln, gn := rpcNamesForModel(modelName)
	svc := findService(p.Service, rpcServiceName)
	if svc == nil {
		return out
	}
	for _, rpc := range svc.RPC {
		if rpc == nil || rpc.RPC == nil {
			continue
		}
		switch rpc.Name {
		case cn:
			out.Create = true
		case un:
			out.Update = true
		case dn:
			out.Delete = true
		case ln:
			out.List = true
		case gn:
			out.GetById = true
		}
	}
	return out
}

// effectiveRPCFlags returns all-true when RpcStrict is off; otherwise matches proto service rpc set for ModelName.
func effectiveRPCFlags(ctx *GenLogicByProtoContext, p *parser.Proto) rpcPresenceFlags {
	if !ctx.RpcStrict {
		return rpcPresenceFlags{
			Create: true, Update: true, Delete: true, List: true, GetById: true,
		}
	}
	return computeRPCPresence(p, ctx.RPCServiceName, ctx.ModelName)
}
