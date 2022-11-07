package execution

import "rosalia64/core/declarations"

func InitializeFunctionDeclarations() {
	declarations.NopInteger = ExecuteNopInteger
}
