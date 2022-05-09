package mbse

/*
//	User defined utility functions
*/

func (s *SmartContract) Contains(st []string, e string) bool {
	for _, v := range st {
		if v == e {
			return true
		}
	}
	return false
}

// Get the Org for a Project for a PORTType
func (s *SmartContract) GetOrgId(orgRoles []OrgRolesStruct, porttype string) []string {
	result := []string{}

	if porttype != "" {
		for i := 0; i < len(orgRoles); i++ {
			if orgRoles[i].PORTType == porttype {
				result = append(result, orgRoles[i].OrgId)
			}
		}
	} else if porttype == "" {
		for i := 0; i < len(orgRoles); i++ {
			result = append(result, orgRoles[i].OrgId)
		}
	}

	return result
}
