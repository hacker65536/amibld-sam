package main

// Buildcfg is build settings
type Buildcfg struct {
	Account     []string
	Region      []string
	Platform    string
	UserName    string
	Commit      string
	Role        []string
	AMI         string
	DynamoTable string
}

// NewBuildcfg is func
func NewBuildcfg(bldcfg Buildcfg) *Buildcfg {

	return &Buildcfg{
		Platform: "amazonlinux2",
	}

}
