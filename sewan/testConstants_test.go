package sewan

import (
	"errors"
)

var (
	testVdcMap = map[string]interface{}{
		nameField:       "Unit test vdc resource",
		enterpriseField: "unit test enterprise",
		datacenterField: "dc1",
		vdcResourceField: []interface{}{
			map[string]interface{}{
				resourceField: ramField,
				totalField:    20,
			},
			map[string]interface{}{
				resourceField: cpuField,
				totalField:    1,
			},
			map[string]interface{}{
				resourceField: "storage_enterprise",
				totalField:    10,
			},
			map[string]interface{}{
				resourceField: "storage_performance",
				totalField:    10,
			},
			map[string]interface{}{
				resourceField: "storage_high_performance",
				totalField:    10,
			},
		},
	}
	noTemplateVMMap = map[string]interface{}{
		nameField:       "Unit test vm",
		enterpriseField: "unit test enterprise",
		stateField:      "UP",
		osField:         "Debian",
		ramField:        8,
		cpuField:        4,
		disksField: []interface{}{
			map[string]interface{}{
				nameField: "disk 1",
				sizeField: 24,
				slugField: slugField,
			},
		},
		nicsField: []interface{}{
			map[string]interface{}{
				vlanNameField:  "vlan 1 update",
				macAdressField: "24",
				connectedField: true,
			},
			map[string]interface{}{
				vlanNameField:  "vlan 2",
				macAdressField: "24",
				connectedField: true,
			},
		},
		vdcField:          vdcField,
		bootField:         "on disk",
		storageClassField: "storage_enterprise",
		slugField:         "42",
		tokenField:        "424242",
		backupField:       "backup-no_backup",
		diskImageField:    "",
		platformNameField: "42",
		backupSizeField:   42,
		commentField:      "42",
		dynamicField:      "42",
	}
	unitTestNonCriticalResourceList = []interface{}{
		map[string]interface{}{"nonCriticalResourceElem1": "unit test value"},
		map[string]interface{}{"nonCriticalResourceElem2": "unit test value"},
	}
	unitTestCriticalResourceList = []interface{}{
		map[string]interface{}{"criticalResourceElem1": "unit test value"},
		map[string]interface{}{"criticalResourceElem2": "unit test value"},
	}
	unitTestOtherResourceList = []interface{}{
		map[string]interface{}{"otherResourceElem1": "unit test value"},
		map[string]interface{}{"otherResourceElem2": "unit test value"},
	}
	errGetEnvMetaFailure         = errors.New("getEnvMetaFailure error")
	errCheckCloudDcStatusFailure = errors.New("checkCloudDcStatusFailure error")
)

const (
	unitTestAPIURL                  = "https://unitTestAPIURL.org"
	unitTestToken                   = "unit test token"
	unitTestEnterprise              = "unit test enterprise"
	vdcCreationFailure              = "VDC creation failed."
	vdcReadFailure                  = "VDC read failed."
	vdcUpdateFailure                = "VDC update failed."
	vdcDeletionFailure              = "VDC deletion failed."
	vmCreationFailure               = "VM creation failed."
	vmReadFailure                   = "VM read failed."
	vmUpdateFailure                 = "VM update failed."
	vmDeletionFailure               = "VM deletion failed."
	errTestResultDiffs              = "\n\rGot: \"%s\"\n\rWant: \"%s\""
	errorTcIDAndWrongVdcUpdateError = "\n\nTC %d : VDC update error was incorrect,"
	errorTcIDAndWrongVMUpdateError  = "\n\nTC %d : VM update error was incorrect,"
)
