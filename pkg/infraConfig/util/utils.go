/*
 * Copyright (c) 2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"github.com/devtron-labs/devtron/pkg/infraConfig/units"
)

// GetUnitSuffix loosely typed method to get the unit suffix using the unitKey type
func GetUnitSuffix(unitKey ConfigKeyStr, unitStr string) units.UnitSuffix {
	switch unitKey {
	case CPU_LIMIT, CPU_REQUEST:
		return units.CPUUnitStr(unitStr).GetCPUUnit()
	case MEMORY_LIMIT, MEMORY_REQUEST:
		return units.MemoryUnitStr(unitStr).GetMemoryUnit()
	}
	return units.TimeUnitStr(unitStr).GetTimeUnit()
}

// GetUnitSuffixStr loosely typed method to get the unit suffix using the unitKey type
func GetUnitSuffixStr(unitKey ConfigKey, unit units.UnitSuffix) string {
	switch unitKey {
	case CPULimit, CPURequest:
		return string(unit.GetCPUUnitStr())
	case MemoryLimit, MemoryRequest:
		return string(unit.GetMemoryUnitStr())
	}
	return string(unit.GetTimeUnitStr())
}

// GetDefaultConfigKeysMap returns a map of default config keys
func GetDefaultConfigKeysMap() map[ConfigKeyStr]bool {
	return map[ConfigKeyStr]bool{
		CPU_LIMIT:      true,
		CPU_REQUEST:    true,
		MEMORY_LIMIT:   true,
		MEMORY_REQUEST: true,
		TIME_OUT:       true,
	}
}

func GetConfigKeyStr(configKey ConfigKey) ConfigKeyStr {
	switch configKey {
	case CPULimit:
		return CPU_LIMIT
	case CPURequest:
		return CPU_REQUEST
	case MemoryLimit:
		return MEMORY_LIMIT
	case MemoryRequest:
		return MEMORY_REQUEST
	case TimeOut:
		return TIME_OUT
	}
	return ""
}

func GetConfigKey(configKeyStr ConfigKeyStr) ConfigKey {
	switch configKeyStr {
	case CPU_LIMIT:
		return CPULimit
	case CPU_REQUEST:
		return CPURequest
	case MEMORY_LIMIT:
		return MemoryLimit
	case MEMORY_REQUEST:
		return MemoryRequest
	case TIME_OUT:
		return TimeOut
	}
	return 0
}
