package sewan

var (
	TEST_VDC_MAP = map[string]interface{}{
		NAME_FIELD:       "Unit test vdc resource",
		ENTERPRISE_FIELD: "unit test enterprise",
		DATACENTER_FIELD: "dc1",
		VDC_RESOURCE_FIELD: []interface{}{
			map[string]interface{}{
				RESOURCE_FIELD: RAM_FIELD,
				TOTAL_FIELD:    20,
			},
			map[string]interface{}{
				RESOURCE_FIELD: CPU_FIELD,
				TOTAL_FIELD:    1,
			},
			map[string]interface{}{
				RESOURCE_FIELD: "storage_enterprise",
				TOTAL_FIELD:    10,
			},
			map[string]interface{}{
				RESOURCE_FIELD: "storage_performance",
				TOTAL_FIELD:    10,
			},
			map[string]interface{}{
				RESOURCE_FIELD: "storage_high_performance",
				TOTAL_FIELD:    10,
			},
		},
	}
	NO_TEMPLATE_VM_MAP = map[string]interface{}{
		NAME_FIELD:       "Unit test vm",
		ENTERPRISE_FIELD: "unit test enterprise",
		STATE_FIELD:      "UP",
		OS_FIELD:         "Debian",
		RAM_FIELD:        8,
		CPU_FIELD:        4,
		DISKS_FIELD: []interface{}{
			map[string]interface{}{
				NAME_FIELD: "disk 1",
				SIZE_FIELD: 24,
				SLUG_FIELD: SLUG_FIELD,
			},
		},
		NICS_FIELD: []interface{}{
			map[string]interface{}{
				VLAN_NAME_FIELD:  "vlan 1 update",
				MAC_ADRESS_FIELD: "24",
				CONNECTED_FIELD:  true,
			},
			map[string]interface{}{
				VLAN_NAME_FIELD:  "vlan 2",
				MAC_ADRESS_FIELD: "24",
				CONNECTED_FIELD:  true,
			},
		},
		VDC_FIELD:           VDC_FIELD,
		BOOT_FIELD:          "on disk",
		STORAGE_CLASS_FIELD: "storage_enterprise",
		SLUG_FIELD:          "42",
		TOKEN_FIELD:         "424242",
		BACKUP_FIELD:        "backup-no_backup",
		DISK_IMAGE_FIELD:    "",
		PLATFORM_NAME_FIELD: "42",
		BACKUP_SIZE_FIELD:   42,
		COMMENT_FIELD:       "42",
		DYNAMIC_FIELD:       "42",
	}
)

const (
	VDC_CREATION_FAILURE = "VDC creation failed."
	VDC_READ_FAILURE     = "VDC read failed."
	VDC_UPDATE_FAILURE   = "VDC update failed."
	VDC_DELETION_FAILURE = "VDC deletion failed."
	VM_CREATION_FAILURE  = "VM creation failed."
	VM_READ_FAILURE      = "VM read failed."
	VM_UPDATE_FAILURE    = "VM update failed."
	VM_DELETION_FAILURE  = "VM deletion failed."
	UNIT_TEST_API_URL    = "https://unitTestApiUrl.org"
)
