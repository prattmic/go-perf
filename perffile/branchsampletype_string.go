// Code generated by "bitstringer -type=BranchSampleType"; DO NOT EDIT

package perffile

import "strconv"

func (i BranchSampleType) String() string {
	if i == 0 {
		return "0"
	}
	s := ""
	if i&BranchSampleAbortTX != 0 {
		s += "AbortTX|"
	}
	if i&BranchSampleAny != 0 {
		s += "Any|"
	}
	if i&BranchSampleAnyCall != 0 {
		s += "AnyCall|"
	}
	if i&BranchSampleAnyReturn != 0 {
		s += "AnyReturn|"
	}
	if i&BranchSampleCall != 0 {
		s += "Call|"
	}
	if i&BranchSampleCallStack != 0 {
		s += "CallStack|"
	}
	if i&BranchSampleCond != 0 {
		s += "Cond|"
	}
	if i&BranchSampleHV != 0 {
		s += "HV|"
	}
	if i&BranchSampleInTX != 0 {
		s += "InTX|"
	}
	if i&BranchSampleIndCall != 0 {
		s += "IndCall|"
	}
	if i&BranchSampleIndJump != 0 {
		s += "IndJump|"
	}
	if i&BranchSampleKernel != 0 {
		s += "Kernel|"
	}
	if i&BranchSampleNoCycles != 0 {
		s += "NoCycles|"
	}
	if i&BranchSampleNoFlags != 0 {
		s += "NoFlags|"
	}
	if i&BranchSampleNoTX != 0 {
		s += "NoTX|"
	}
	if i&BranchSampleTypeSave != 0 {
		s += "TypeSave|"
	}
	if i&BranchSampleUser != 0 {
		s += "User|"
	}
	i &^= 131071
	if i == 0 {
		return s[:len(s)-1]
	}
	return s + "0x" + strconv.FormatUint(uint64(i), 16)
}
