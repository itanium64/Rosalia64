package execution

import "rosalia64/core/declarations"

func InitializeFunctionDeclarations() {
	declarations.NopInteger = ExecuteNopInteger
	declarations.NopBranch = ExecuteNopBranch
	declarations.NopMemory = ExecuteNopMemory
	declarations.IntegerStoreRegister = ExecuteIntegerStoreRegister
	declarations.IntegerLoadNoBaseUpdateForm = ExecuteIntegerLoadNoBaseUpdateForm
	declarations.BranchIndirectReturn = ExecuteBranchIndirectReturn
	declarations.AddlImm22 = ExecuteAddlImm22
	declarations.AddsImm14 = ExecuteAddsImm14
	declarations.IntegerCompareRegisterForm = ExecuteIntegerCompareRegisterForm
	declarations.ExecuteIPRelativeBranchCond = ExecuteIPRelativeBranchCond
}
