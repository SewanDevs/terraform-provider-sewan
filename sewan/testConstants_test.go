package sewan

import (
	"errors"
)

var (
	testVdcMap = map[string]interface{}{
		nameField:       "Unit test vdc resource",
		enterpriseField: "unit test enterprise",
		dataCenterField: "dc1",
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
	errGetEnvMetaFailure         = errors.New("getEnvMetaFailure error")
	errCheckCloudDcStatusFailure = errors.New("checkCloudDcStatusFailure error")
	rightDatacenter              = "dc2"
	wrongDatacenter              = "wrongDatacenter"
	dataCenterMetaDataList       = []interface{}{
		map[string]interface{}{
			"id":       1,
			"created":  "2017-06-29T12:10:33+02:00",
			"modified": "2017-12-07T14:19:54+01:00",
			"name":     "DC2",
			"slug":     rightDatacenter,
			"cos":      "Mono",
			"weight":   1,
			"manager":  3,
		},
		map[string]interface{}{
			"id":       2,
			"created":  "2017-06-29T12:10:33+02:00",
			"modified": "2017-12-07T14:19:48+01:00",
			"name":     "DC1",
			"slug":     "dc1",
			"cos":      "Mono",
			"weight":   2,
			"manager":  4,
		},
		map[string]interface{}{
			"id":       3,
			"created":  "2017-06-29T12:10:33+02:00",
			"modified": "2018-04-24T11:37:01+02:00",
			"name":     "HA",
			"slug":     "ha",
			"cos":      "HA",
			"weight":   1,
			"manager":  7,
		},
	}
	enterpriseResourceMetaDataList = []interface{}{
		map[string]interface{}{
			"id":            4,
			"enterprise":    "unit-test-enterprise",
			"created":       "2017-06-29T12:10:35+02:00",
			"modified":      "2018-05-28T12:28:42+02:00",
			"cos":           "Mono",
			"name":          "ram",
			"used":          324,
			"total":         350,
			"slug":          "unit-test-enterprise-mono-ram",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            5,
			"enterprise":    "unit-test-enterprise",
			"created":       "2017-06-29T12:10:35+02:00",
			"modified":      "2018-05-28T12:28:32+02:00",
			"cos":           "Mono",
			"name":          "cpu",
			"used":          275,
			"total":         300,
			"slug":          "unit-test-enterprise-mono-cpu",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            6,
			"enterprise":    "unit-test-enterprise",
			"created":       "2017-06-29T12:10:35+02:00",
			"modified":      "2018-02-14T17:32:15+01:00",
			"cos":           "Mono",
			"name":          "storage_enterprise",
			"used":          7708,
			"total":         8000,
			"slug":          "unit-test-enterprise-mono-storage_enterprise",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            7,
			"enterprise":    "unit-test-enterprise",
			"created":       "2017-06-29T12:10:35+02:00",
			"modified":      "2018-07-31T15:55:06+02:00",
			"cos":           "Mono",
			"name":          "storage_performance",
			"used":          630,
			"total":         700,
			"slug":          "unit-test-enterprise-mono-storage_performance",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            8,
			"enterprise":    "unit-test-enterprise",
			"created":       "2017-06-29T12:10:35+02:00",
			"modified":      "2018-02-06T11:02:17+01:00",
			"cos":           "Mono",
			"name":          "storage_high_performance",
			"used":          10,
			"total":         20,
			"slug":          "unit-test-enterprise-mono-storage_high_performance",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            305,
			"enterprise":    "unit-test-enterprise",
			"created":       "2017-10-10T12:19:51+02:00",
			"modified":      "2017-10-10T12:19:51+02:00",
			"cos":           "HA",
			"name":          "cpu",
			"used":          5,
			"total":         10,
			"slug":          "resource-cpu-rd-ha",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            306,
			"enterprise":    "unit-test-enterprise",
			"created":       "2017-10-10T12:20:11+02:00",
			"modified":      "2017-10-10T12:20:11+02:00",
			"cos":           "HA",
			"name":          "ram",
			"used":          5,
			"total":         10,
			"slug":          "ram-ha-rd",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            314,
			"enterprise":    "unit-test-enterprise",
			"created":       "2018-04-03T16:09:32+02:00",
			"modified":      "2018-04-24T15:50:56+02:00",
			"cos":           "HA",
			"name":          "storage_enterprise",
			"used":          60,
			"total":         100,
			"slug":          "storage_enterprise-ha",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            315,
			"enterprise":    "unit-test-enterprise",
			"created":       "2018-04-24T12:35:55+02:00",
			"modified":      "2018-04-24T15:51:04+02:00",
			"cos":           "HA",
			"name":          "storage_performance",
			"used":          55,
			"total":         100,
			"slug":          "unit-test-enterprise-ha-storage_performance",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            316,
			"enterprise":    "unit-test-enterprise",
			"created":       "2018-04-24T12:36:15+02:00",
			"modified":      "2018-04-24T15:51:13+02:00",
			"cos":           "HA",
			"name":          "storage_high_performance",
			"used":          0,
			"total":         100,
			"slug":          "unit-test-enterprise-ha-storage_high_performance",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            55,
			"enterprise":    "unit-test-enterprise",
			"created":       "2017-06-29T12:10:35+02:00",
			"modified":      "2017-08-10T05:01:03+02:00",
			"cos":           "Global",
			"name":          "backup",
			"used":          220,
			"total":         220,
			"slug":          "unit-test-enterprise-clouddc-backup",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            196,
			"enterprise":    "unit-test-enterprise",
			"created":       "2017-06-29T12:10:35+02:00",
			"modified":      "2018-02-21T12:45:28+01:00",
			"cos":           "Global",
			"name":          "license_win_server",
			"used":          7,
			"total":         20,
			"slug":          "unit-test-enterprise-global-license_win_server",
			"dynamic_field": nil,
			"service":       1,
		},
		map[string]interface{}{
			"id":            313,
			"enterprise":    "unit-test-enterprise",
			"created":       "2018-02-15T18:39:17+01:00",
			"modified":      "2018-02-16T15:50:16+01:00",
			"cos":           "Global",
			"name":          "license_redhat",
			"used":          2,
			"total":         3,
			"slug":          "sewan-rd-cloud-daatcenter-vdc-rd-licence-redhat",
			"dynamic_field": nil,
			"service":       1,
		},
	}
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
