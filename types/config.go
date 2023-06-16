package types

// StandardConfiguration presents the standard-json
// configuration generated as output of zksolc compiler
type StandardConfiguration struct {
	Format       string `json:"_format"`
	ContractName string `json:"contractName"`
	SourceName   string `json:"sourceName"`
	Abi          []struct {
		Inputs []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			Indexed      bool   `json:"indexed,omitempty"`
		} `json:"inputs"`
		StateMutability string `json:"stateMutability,omitempty"`
		Type            string `json:"type"`
		Anonymous       bool   `json:"anonymous,omitempty"`
		Name            string `json:"name,omitempty"`
		Outputs         []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"outputs,omitempty"`
	} `json:"abi"`
	Bytecode         string `json:"bytecode"`
	DeployedBytecode string `json:"deployedBytecode"`
	LinkReferences   struct {
	} `json:"linkReferences"`
	DeployedLinkReferences struct {
	} `json:"deployedLinkReferences"`
	FactoryDeps struct {
	} `json:"factoryDeps"`
}
