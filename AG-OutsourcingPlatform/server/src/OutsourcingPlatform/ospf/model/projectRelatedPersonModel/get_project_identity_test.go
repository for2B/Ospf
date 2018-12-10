package projectRelatedPersonModel

import (
	"testing"
	"OutsourcingPlatform/ospf/infrastructure"
	"OutsourcingPlatform/ospf/infrastructure/option"
	"os"
)
var model *ProjectRelatedPersonModel
func TestMain(m *testing.M) {
	proDir := "E:/w外包项目/OSPF/OutsourcingPlatform/AG-OutsourcingPlatform/server/bin"
	opts := option.NewOpts(proDir, "")
	model = &ProjectRelatedPersonModel{
		infrastructure.NewInfrastructure(opts),
	}
	ExitCode := m.Run()
	os.Exit(ExitCode)
}

func Benchmark_GetProjectIdentity(b *testing.B) {
	for i := 0 ; i<b.N; i++{
		model.GetProjectIdentity("60","mockopenid3")
	}
}


